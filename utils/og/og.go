package og

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
)

func fetchHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func ExtractOGData(url string) (*opengraph.OpenGraph, error) {
	html, err := fetchHTML(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(strings.NewReader(string(html)))

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return og, nil
}
