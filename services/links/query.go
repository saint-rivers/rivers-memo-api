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

func GetMemosWithTagFilter(ctx context.Context, db *sql.DB, req *pb.PageRequest) ([]*pb.Memo, error) {
	q := `
			SELECT n.id, link, title, description, image_url, string_agg(
				nt.tag, ','
			) as tags
			FROM note_tag nt
			JOIN note n ON n.id = nt.note
			WHERE id < $1
			AND tag = ANY($2)
			GROUP BY n.id
			ORDER BY id desc
			LIMIT $3`

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
		tags := strings.Split(tag, ",")
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

func GetMemos(ctx context.Context, db *sql.DB, req *pb.PageRequest) ([]*pb.Memo, error) {
	q := `
			SELECT n.id, link, title, description, image_url, string_agg(
				nt.tag, ','
			) AS tags
			FROM note_tag nt
			JOIN note n ON n.id = nt.note
			WHERE id < $1
			GROUP BY n.id
			ORDER BY id desc
			LIMIT $2`

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
		tags := strings.Split(tag, ",")
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
