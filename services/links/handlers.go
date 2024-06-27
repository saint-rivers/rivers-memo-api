package links

import (
	"context"
	"database/sql"
	"log"
)

type CreateLinkMemoParams struct {
	ID   int
	Link string
	Tags []string
}

type Memo struct {
	ID   int
	Link string
	Tags []string
}

func ListMemo(db *sql.DB) ([]Memo, error) {
	q := "SELECT id, link FROM note ORDER BY id desc"

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memos []Memo
	for rows.Next() {
		var m Memo
		if err := rows.Scan(&m.ID, &m.Link); err != nil {
			return nil, err
		}
		memos = append(memos, m)
	}
	return memos, nil
}

func CreateLinkMemo(db *sql.DB, param *CreateLinkMemoParams) error {
	q := "INSERT INTO note(id, link) VALUES($1, $2)"
	insert, err := db.ExecContext(context.Background(), q, param.ID, param.Link)
	if err != nil {
		log.Fatalf("insert error: %s", err)
		return err
	}
	rows, err := insert.RowsAffected()
	if err != nil {
		log.Fatalf("unable to read insert results: %s", err)
		return err
	}
	log.Printf("rows affected: %d", rows)
	return nil
}
