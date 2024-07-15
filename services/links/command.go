package links

import (
	"context"
	"database/sql"
	"log"
)

type MemoParams struct {
	ID             int
	ChannelMessage string
}

type TagParams struct {
	ID   int
	Tags []string
}

func AppendTag(ctx context.Context, db *sql.DB, params *TagParams) error {
	for _, tag := range params.Tags {
		query := "INSERT INTO tag(tag_name) VALUES($1) ON CONFLICT DO NOTHING"
		_, err := db.ExecContext(ctx, query, tag)
		if err != nil {
			log.Printf("insert into 'tag' error: %s", err)
			return err
		}

		query = "INSERT INTO note_tag(note, tag) VALUES($1, $2) ON CONFLICT DO NOTHING"
		_, err = db.ExecContext(ctx, query, params.ID, tag)
		if err != nil {
			log.Printf("insert into 'note_tag' error: %s", err)
			return err
		}
	}
	return nil
}

func CreateMemo(ctx context.Context, db *sql.DB, params *MemoParams) error {
	query := "INSERT INTO note(id, link) VALUES($1, $2)"
	insert, err := db.ExecContext(ctx, query, params.ID, params.ChannelMessage)
	if err != nil {
		log.Println("insert into 'note' error: %s", err)
		return err
	}
	rows, err := insert.RowsAffected()
	if err != nil {
		log.Println("unable to read insert results: %s", err)
		return err
	}
	log.Printf("rows inserted: %d", rows)
	return nil
}
