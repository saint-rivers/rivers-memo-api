package links

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	pb "rivers-memo-cli/generated/memo"
	"rivers-memo-cli/utils/og"
	"strconv"

	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/labstack/echo/v4"
)

type CreateLinkMemoParams struct {
	ID             int
	ChannelMessage string
	Tags           []string
}

type Memo struct {
	ID          int32
	Link        string
	Tags        []string
	Title       *string
	Description *string
	ImageUrl    *string
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
			&pb.PageRequest{Last: &last, Size: size},
		)
		if err != nil {
			log.Println(err.Error())
		}

		return c.JSON(http.StatusOK, memos)
	})

	// e.GET("/scripts/memo", func(c echo.Context) error {
	// 	last := "999999"
	// 	memos, err := GetMemos(
	// 		context.Background(),
	// 		db,
	// 		&pb.PageRequest{Last: &last, Size: 20},
	// 	)
	// 	if err != nil {
	// 		print(err)
	// 	}

	// 	for i := range memos {
	// 		m := &memos[i]
	// 		ogData, err := og.ExtractOGData(m.GetLink())
	// 		if err != nil {
	// 			print(err)
	// 			continue
	// 		}
	// 		_ = UpdateMemo(
	// 			context.Background(),
	// 			db,
	// 			&ogData.Images[0].URL, &ogData.Title, &ogData.Description, m.Link,
	// 		)
	// 	}
	// 	return c.JSON(http.StatusOK, memos)
	// })
}

func ProcessChanncelMessage(db *sql.DB, param *CreateLinkMemoParams) error {
	msg := param.ChannelMessage

	switch getMessageType(msg) {
	case Command:
		msg = msg[5:]
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

			err = AppendTag(context.Background(), db, &TagParams{ID: id, Tags: args[2:]})
			if err != nil {
				log.Printf("tag insert error: %s", err.Error())
			}
		}
	case Link:
		link := param.ChannelMessage
		ogData := make(chan *opengraph.OpenGraph)
		go func() {
			v, _ := og.ExtractOGData(link)
			ogData <- v
		}()

		ogd := <-ogData
		log.Println(ogd.Images[0].URL)
		CreateMemo(context.Background(), db,
			&MemoParams{
				ID:             param.ID,
				ChannelMessage: link,
				Title:          ogd.Title,
				Description:    ogd.Description,
				Image:          ogd.Images[0].URL,
			})

	case Message:

	}

	return nil
}
