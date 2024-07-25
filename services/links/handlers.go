package links

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"rivers-memo-cli/api"
	"rivers-memo-cli/client"
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

type MemoServer struct {
	pb.UnimplementedMemoServiceServer
}

func (s *MemoServer) GetMemos(ctx context.Context, in *pb.PageRequest) (*pb.MemoReply, error) {
	var memos []*pb.Memo
	var err error

	last := "999999"
	if in.Last == nil {
		in.Last = &last
	}
	if len(in.Tags) == 0 {
		memos, err = GetMemos(
			context.Background(),
			client.Persistence,
			&pb.PageRequest{Last: in.Last, Size: in.Size, Tags: []string{}},
		)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		memos, err = GetMemosWithTagFilter(
			context.Background(),
			client.Persistence,
			&pb.PageRequest{Last: in.Last, Size: in.Size, Tags: in.Tags},
		)
		if err != nil {
			log.Println(err.Error())
		}
	}
	return &pb.MemoReply{Memos: memos}, nil
}

func ConfigRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/api/v1/memo", func(c echo.Context) error {
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

		tags := c.QueryParams()["tag"]
		var memos []*pb.Memo
		if len(tags) == 0 {
			memos, err = GetMemos(
				context.Background(),
				db,
				&pb.PageRequest{Last: &last, Size: size, Tags: []string{}},
			)
		} else {
			memos, err = GetMemosWithTagFilter(
				context.Background(),
				db,
				&pb.PageRequest{Last: &last, Size: size, Tags: tags},
			)
		}

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
	e.POST("/api/v1/memo/:id/tags", func(c echo.Context) error {
		mId := strings.Split(c.Request().URL.Path, "/")[4]

		u := new(TagParams)
		if err := c.Bind(u); err != nil {
			return c.JSON(http.StatusBadRequest, api.BadRequestError("bad request body"))
		}
		var err error
		u.ID, err = strconv.Atoi(mId)
		err = AppendTag(context.Background(), db, u)
		if err != nil {
			return c.JSON(http.StatusBadRequest, api.BadRequestError("unable to append tags"))
		}
		msg := fmt.Sprintf("updated memo: %s", mId)
		return c.JSON(http.StatusOK, api.SuccessMessage(msg))
	})
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
