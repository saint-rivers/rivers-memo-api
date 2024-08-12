package links

import (
	"context"
	"database/sql"
	"log"

	"github.com/lib/pq"
)

type MemoParams struct {
	ID             int
	ChannelMessage string
	Title          string
	Description    string
	Image          string
}

type TagParams struct {
	ID   int
	Tags []string
}

func RemoveTag(ctx context.Context, db *sql.DB, params *TagParams) error {
	q := `DELETE FROM note_tag WHERE note = $1 AND tag = any($2)`
	_, err := db.ExecContext(ctx, q, params.ID, pq.Array(params.Tags))
	if err != nil {
		return err
	}
	log.Printf("removed tag (%s) from (%d)", params.Tags, params.ID)
	return nil
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
	query := `
	INSERT INTO note(link, title, description, image_url) 
	VALUES($1, $2, $3, $4)`

	insert, err := db.ExecContext(
		ctx,
		query,
		params.ChannelMessage,
		params.Title,
		params.Description,
		params.Image,
	)
	if err != nil {
		log.Printf("insert into 'note' error: %s", err)
		return err
	}
	rows, err := insert.RowsAffected()
	if err != nil {
		log.Printf("unable to read insert results: %s", err)
		return err
	}
	log.Printf("rows inserted: %d", rows)
	return nil
}

func UpdateMemo(ctx context.Context, db *sql.DB, image *string, title *string, desc *string, link string) error {
	q := `
	update note set title = $1, description = $2, image_url = $3 
	where link = $4`

	_, err := db.Exec(q, title, desc, image, link)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func DeleteMemo(ctx context.Context, db *sql.DB, id int64) error {
	q := `
	delete from note 
	where id = $1`

	_, err := db.Exec(q, id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
