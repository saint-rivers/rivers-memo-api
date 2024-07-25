package tags

import (
	"database/sql"
	"net/http"
	"rivers-memo-cli/api"

	"github.com/labstack/echo/v4"
)

func ConfigRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/api/v1/tags", func(c echo.Context) error {
		tags, err := getTags(db)
		if err != nil {
			return c.JSON(http.StatusBadRequest, api.GenericError())
		}
		return c.JSON(http.StatusOK, tags)
	})
}

type Tag string

func getTags(db *sql.DB) ([]Tag, error) {
	q := `
	SELECT tag_name FROM tag
	`

	rows, err := db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []Tag

	for rows.Next() {
		var t Tag
		if err := rows.Scan(&t); err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}
	return tags, nil
}
