package links

import (
	"context"
	"database/sql"
	"fmt"
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
			return id, fmt.Errorf("Not found")
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
