package links

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	pb "rivers-memo-cli/generated/memo"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
)

type CreateLinkMemoParams struct {
	ID             int
	ChannelMessage string
	Tags           []string
}

type Memo struct {
	ID   int32
	Link string
	Tags []string
}

func ConfigRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/memo", func(c echo.Context) error {
		arg1 := c.QueryParam("size")
		if arg1 == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"err": "Missing required argumnets: size"})
		}
		s, err := strconv.ParseInt(arg1, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"err": "Size must be a number"})
		}
		size := int32(s)

		last := c.QueryParam("last")
		if last == "" {
			last = "999999"
		}

		memos, err := GetMemos(
			context.Background(),
			db,
			&pb.PageRequest{Last: &last, Size: size})
		return c.JSON(http.StatusOK, memos)
	})
}

func GetMemos(ctx context.Context, db *sql.DB, req *pb.PageRequest) ([]pb.MemoReply, error) {
	q := `SELECT id, link FROM note 
	WHERE id < $1
	ORDER BY id desc 
	LIMIT $2`

	rows, err := db.Query(q, req.Last, req.Size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memos []pb.MemoReply
	for rows.Next() {
		var m Memo
		if err := rows.Scan(&m.ID, &m.Link); err != nil {
			return nil, err
		}
		memos = append(memos, pb.MemoReply{Id: m.ID, Link: m.Link, Tags: nil})
	}
	return memos, nil
}

func ProcessChanncelMessage(db *sql.DB, param *CreateLinkMemoParams) error {
	msg := param.ChannelMessage

	switch getMessageType(msg) {
	case Command:
		msg = msg[5:len(msg)]
		args := strings.Split(msg, " ")
		switch cmd := args[0]; cmd {
		case "tag":
			offset, err := strconv.Atoi(args[1])
			if err != nil {
				return errors.New("unable to convert index to int")
			}
			if offset > 0 {
				return errors.New("invalid index")
			}
			offset = offset * -1

			id, err := ReadMessageByRelativeIndex(context.Background(), db, &MemoReadParam{Offset: offset})
			if err != nil {
				log.Printf("unable to find message: %s", err.Error())
			}

			err = AppendTag(context.Background(), db, &TagParams{ID: id, Tags: args[2:len(args)]})
			if err != nil {
				log.Printf("tag insert error: %s", err.Error())
			}
		}
	case Link:
		CreateMemo(context.Background(), db, &MemoParams{ID: param.ID, ChannelMessage: msg})

	case Message:

	}

	return nil
}
