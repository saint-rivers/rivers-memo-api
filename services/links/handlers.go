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
	// ID             int      `json:"id"`
	ChannelMessage string   `form:"message" json:"message"`
	Tags           []string `form:"tags" json:"tags"`
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

type TagsUpdateRequest struct {
	Removed []string `query:"removed"`
	Added   []string `query:"added"`
}

func (s *MemoServer) GetMemos(ctx context.Context, in *PageRequest) (*pb.MemoReply, error) {
	var memos []*pb.Memo
	var err error

	if in.Last == 0 {
		in.Last = 999999
	}
	if len(in.Tags) == 0 {
		memos, err = GetMemos(
			context.Background(),
			client.Persistence,
			&PageRequest{Last: in.Last, Size: in.Size, Tags: []string{}},
		)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		memos, err = GetMemosWithTagFilter(
			context.Background(),
			client.Persistence,
			&PageRequest{Last: in.Last, Size: in.Size, Tags: in.Tags},
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

		lastreq := c.QueryParam("last")

		last, err := strconv.ParseInt(lastreq, 10, 64)
		if last == 0 {
			last = 999999
		}

		tags := c.QueryParams()["tag"]
		var memos []*pb.Memo
		if len(tags) == 0 {
			memos, err = GetMemos(
				context.Background(),
				db,
				&PageRequest{Last: last, Size: size, Tags: []string{}},
			)
		} else {
			memos, err = GetMemosWithTagFilter(
				context.Background(),
				db,
				&PageRequest{Last: last, Size: size, Tags: tags},
			)
		}

		if err != nil {
			log.Println(err.Error())
		}
		return c.JSON(http.StatusOK, memos)
	})

	e.PUT("/api/v1/memo/:id/tags", func(c echo.Context) error {
		mId := c.Param("id")
		id, err := strconv.Atoi(mId)

		var req TagsUpdateRequest
		err = c.Bind(&req)
		if err != nil {
			return c.JSON(http.StatusBadRequest, api.BadRequestError("required at least two empty arrays"))
		}

		if len(req.Removed) > 0 {
			err = RemoveTag(context.Background(), db, &TagParams{ID: id, Tags: req.Removed})
			if err != nil {
				return c.JSON(http.StatusBadRequest, api.BadRequestError("remove tags failed"))
			}
		}
		log.Println(req.Added)
		if len(req.Added) > 0 {
			err = AppendTag(context.Background(), db, &TagParams{ID: id, Tags: req.Added})
			if err != nil {
				return c.JSON(http.StatusBadRequest, api.BadRequestError("append tags failed"))
			}
		}

		return c.JSON(http.StatusOK, api.SuccessMessage("updated successfully"))
	})

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

	e.POST("/api/v1/memo", func(c echo.Context) error {
		req := new(CreateLinkMemoParams)
		if err := c.Bind(req); err != nil {
			msg := fmt.Sprintf("invalid request body")
			return c.JSON(http.StatusBadRequest, api.BadRequestError(msg))
		}
		err := insertMemo(db, req)
		if err != nil {
			msg := fmt.Sprintf("unabel to insert message")
			return c.JSON(http.StatusBadRequest, api.BadRequestError(msg))
		}
		msg := fmt.Sprintf("successfully inserted message")
		return c.JSON(http.StatusOK, api.SuccessMessage(msg))
	})

	e.DELETE("/api/v1/memo/:id", func(c echo.Context) error {
		mId := strings.Split(c.Request().URL.Path, "/")[4]
		id, err := strconv.ParseInt(mId, 10, 64)
		if err != nil {
			msg := fmt.Sprintf("invalid id")
			return c.JSON(http.StatusBadRequest, api.NotFoundError(msg))
		}
		err = DeleteMemo(context.Background(), db, id)
		if err != nil {
			msg := fmt.Sprintf("id not found. unable to delete")
			return c.JSON(http.StatusBadRequest, api.NotFoundError(msg))
		}
		msg := fmt.Sprintf("successfully deleted message")
		return c.JSON(http.StatusOK, api.SuccessMessage(msg))
	})

	e.GET("/api/v1/memo/:key", func(c echo.Context) error {
		req := new(MemoSearchRequest)

		if err := c.Bind(req); err != nil {
			msg := fmt.Sprintf("invalid request")
			return c.JSON(http.StatusBadRequest, api.BadRequestError(msg))
		}

		if req.Last == 0 {
			req.Last = 9999
		}

		payload, err := SearchMemo(context.Background(), db, req)
		if err != nil {
			msg := fmt.Sprintf(err.Error())
			return c.JSON(http.StatusBadRequest, api.NotFoundError(msg))
		}
		return c.JSON(http.StatusOK, payload)
	})
}

type MemoSearchPath struct {
	Key string `param:"key" json:"key"`
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
		insertMemo(db, param)

	case Message:

	}

	return nil
}

func insertMemo(db *sql.DB, param *CreateLinkMemoParams) error {
	link := param.ChannelMessage
	ogData := make(chan *opengraph.OpenGraph)
	go func() {
		v, _ := og.ExtractOGData(link)
		ogData <- v
	}()

	ogd := <-ogData
	var memo *MemoParams
	if len(ogd.Images) > 0 {
		log.Println(ogd.Images[0].URL)
		memo = &MemoParams{
			// ID:             param.ID,
			ChannelMessage: link,
			Title:          ogd.Title,
			Description:    ogd.Description,
			Image:          ogd.Images[0].URL,
		}
	} else {
		log.Println("no image")
		memo = &MemoParams{
			// ID:             param.ID,
			ChannelMessage: link,
			Title:          ogd.Title,
			Description:    ogd.Description,
			Image:          "",
		}
	}
	err := CreateMemo(context.Background(), db, memo)
	return err
}
