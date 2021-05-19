package main

import (
	"log"
	"os"
	"strconv"

	"github.com/GabMar/diretta-it-scraper/internal/pkg/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Printf("No .env file loaded.")
	}

	mpr := getMaxParallelRequestsNumber()

	mh := handler.MatchHandler{
		MaxParallelRequests: mpr,
	}

	e := echo.New()
	e.GET("/v1/matches", mh.Handle)
	e.Logger.Fatal(e.Start(":17171"))
}

func getMaxParallelRequestsNumber() int {
	if os.Getenv("MAX_PARALLEL_HTTP_REQUESTS") == "" {
		log.Printf("No value for 'MAX_PARALLEL_HTTP_REQUESTS', using default.")

		return 1
	}

	mpr, err := strconv.Atoi(os.Getenv("MAX_PARALLEL_HTTP_REQUESTS"))

	if err != nil {
		log.Fatal("Invalid value for 'MAX_PARALLEL_HTTP_REQUESTS'.")
	}

	return mpr
}
