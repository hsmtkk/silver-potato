package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type GourmetSearcher struct {
	apiKey string
}

func (g *GourmetSearcher) SearchGourmet(ctx echo.Context) error {
	smallArea := ctx.QueryParam("small_area")
	g.callAPI(smallArea)
	// レスポンスXMLを解析し、何らか整形する
	return ctx.String(http.StatusOK, smallArea)
}

func (g *GourmetSearcher) callAPI(smallArea string) error {
	// TODO
	fmt.Printf("call API using API key: %s; smallArea: %s\n", g.apiKey, smallArea)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s apiKey", os.Args[0])
	}
	apiKey := os.Args[1]
	searcher := &GourmetSearcher{apiKey: apiKey}
	e := echo.New()
	e.GET("/", searcher.SearchGourmet)
	e.Logger.Fatal(e.Start(":1323"))
}
