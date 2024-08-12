package links

import (
	"context"
	"database/sql"
	"fmt"
	pb "rivers-memo-cli/generated/memo"
	"strconv"
	"strings"

	"github.com/lib/pq"
	// "github.com/lib/pq"
)

type TagSearchParam struct {
	Tags []string
}

type MemoReadParam struct {
	Offset int
}

type MemoResponse struct {
	ID   int
	Link string
	Tags []string
}

type PageRequest struct {
	Last int64    `json:"last,omitempty"`
	Size int32    `json:"size,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

type MemoSearchRequest struct {
	Key  string `param:"key" validate:"required"`
	Last int32  `query:"last"`
	Size int32  `query:"size" validate:"required"`
}

func ReadMessageByRelativeIndex(ctx context.Context, db *sql.DB, params *MemoReadParam) (int, error) {
	query := "select id from note order by id desc limit 1 offset $1"
	var id int
	if err := db.QueryRowContext(ctx, query, params.Offset).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return id, fmt.Errorf("not found")
		}
		return id, err
	}
	return id, nil
}

func SearchByTags(ctx context.Context, db *sql.DB, params *TagSearchParam) ([]MemoResponse, error) {
	query := `select id, link from note_tag nt join note n on n.id = nt.note `
	query = buildQuery(query, len(params.Tags))
	query += "order by id desc"

	var memos []MemoResponse
	tags := make([]interface{}, len(params.Tags))
	rows, err := db.QueryContext(ctx, query, tags...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var memo MemoResponse
		if err := rows.Scan(&memo.ID, &memo.Link); err != nil {
			memos = append(memos, memo)
			return memos, err
		}
	}
	if err = rows.Err(); err != nil {
		return memos, err
	}
	return memos, nil
}

func buildQuery(q string, size int) string {
	if size == 1 {
		q += "where tag = ? "
	} else if size > 1 {
		q += "where tag = ? "
		size = size - 1

		i := 0
		for i < size {
			q += "and tag = ? "
			i++
		}
	}
	return q
}

func placeholders(n int) string {
	offset := 3
	ps := ""
	for i := 0 + offset; i < n+offset; i++ {
		ps += "$" + strconv.Itoa(i)
		if i < n+offset-1 {
			ps += ","
		}
	}
	return ps
}
func unpackArray[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))
	for i, e := range s {
		r[i] = e
	}
	return r
}

func GetMemosWithTagFilter(ctx context.Context, db *sql.DB, req *PageRequest) ([]*pb.Memo, error) {
	q := `
WITH s AS (
  SELECT distinct id, link, title, description, image_url
  FROM note n
  INNER JOIN note_tag nt 
  ON nt.note = n.id
  WHERE id < $1
  AND tag = any($2)
  ORDER BY id DESC
  LIMIT $3
)
SELECT  s.id, s.link, s.title, s.description, image_url, coalesce(string_agg(nt.tag, ','), '') AS tags 
FROM s JOIN note_tag nt
ON s.id = nt.note
GROUP BY s.id, s.link, s.title, s.description, s.image_url
ORDER BY id DESC
	`

	rows, err := db.Query(q, req.Last, pq.Array(req.Tags), req.Size) //, req.Last, pq.Array(req.Tags), req.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memos []*pb.Memo

	for rows.Next() {
		var m Memo
		var tag string
		if err := rows.Scan(&m.ID, &m.Link, &m.Title, &m.Description, &m.ImageUrl, &tag); err != nil {
			return nil, err
		}
		var tags []string
		if tag == "" {
			tags = []string{}
		} else {
			tags = strings.Split(tag, ",")
		}
		memos = append(
			memos,
			&pb.Memo{
				Id:          m.ID,
				Link:        m.Link,
				Tags:        tags,
				Title:       m.Title,
				Description: m.Description,
				ImageUrl:    m.ImageUrl,
			},
		)
	}
	return memos, nil
}

func GetMemos(ctx context.Context, db *sql.DB, req *PageRequest) ([]*pb.Memo, error) {
	q := `
select s.id, s.link, s.title, s.description, image_url, coalesce(string_agg(nt.tag, ','), '') AS tags
from note s
left join note_tag nt
on s.id = nt.note
where id < $1
group by s.id, s.link, s.title, s.description
order by id desc
limit $2`

	rows, err := db.Query(q, req.Last, req.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var memos []*pb.Memo

	for rows.Next() {
		var m Memo
		var tag string
		if err := rows.Scan(&m.ID, &m.Link, &m.Title, &m.Description, &m.ImageUrl, &tag); err != nil {
			return nil, err
		}
		var tags []string
		if tag == "" {
			tags = []string{}
		} else {
			tags = strings.Split(tag, ",")
		}
		memos = append(
			memos,
			&pb.Memo{
				Id:          m.ID,
				Link:        m.Link,
				Tags:        tags,
				Title:       m.Title,
				Description: m.Description,
				ImageUrl:    m.ImageUrl,
			},
		)
	}
	return memos, nil
}

func SearchMemo(ctx context.Context, db *sql.DB, req *MemoSearchRequest) ([]*pb.Memo, error) {
	q := `
	select s.id, s.link, s.title, s.description, image_url
	from note s
	where title || link ilike CONCAT('%', cast($1 as varchar), '%')
	and id < $2
	order by id desc
	limit $3
	`
	rows, err := db.QueryContext(ctx, q, req.Key, req.Last, req.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var memos []*pb.Memo

	for rows.Next() {
		var m Memo
		if err := rows.Scan(&m.ID, &m.Link, &m.Title, &m.Description, &m.ImageUrl); err != nil {
			return nil, err
		}
		memos = append(
			memos,
			&pb.Memo{
				Id:          m.ID,
				Link:        m.Link,
				Tags:        []string{},
				Title:       m.Title,
				Description: m.Description,
				ImageUrl:    m.ImageUrl,
			},
		)
	}
	return memos, nil

}
