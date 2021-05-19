package handler

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/GabMar/diretta-it-scraper/internal/pkg/response"
	"github.com/gocolly/colly"
	"github.com/labstack/echo/v4"
)

type MatchHandler struct {
	MaxParallelRequests int
}

func (m *MatchHandler) Handle(c echo.Context) error {
	mc := make(chan response.Match)
	wg := new(sync.WaitGroup)
	matches := []response.Match{}
	matchTypes := []string{"sched", "live", "fin"}

	wg.Add(len(matchTypes))

	for _, matchType := range matchTypes {
		go m.getMatches(matchType, mc, wg)
	}

	go func() {
		wg.Wait()
		close(mc)
	}()

	for match := range mc {
		matches = append(matches, match)
	}

	return c.JSON(http.StatusOK, matches)
}

func (m *MatchHandler) getMatches(matchtype string, mc chan response.Match, wg *sync.WaitGroup) {
	url := "https://m.diretta.it/"
	cl := colly.NewCollector(
		colly.Async(true),
	)

	configErr := cl.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: m.MaxParallelRequests})

	if configErr != nil {
		log.Fatal("Invalid configuration for concurrent requests.")
	}

	cl.OnHTML(fmt.Sprintf("#score-data .%v", matchtype), func(e *colly.HTMLElement) {
		nextPage := e.Request.AbsoluteURL(e.Attr("href"))
		err := cl.Visit(nextPage)

		if err != nil {
			fmt.Println("Request URL:", nextPage, "\nError:", err)
		}
	})

	cl.OnHTML("#main.soccer", func(e *colly.HTMLElement) {
		match := getMatchFromNode(e)

		if match != nil {
			mc <- *match
		}
	})

	// Set error handler
	cl.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := cl.Visit(url)

	if err != nil {
		fmt.Println("Request URL:", url, "\nError:", err)
	}

	cl.Wait()
	wg.Done()
}

func getMatchFromNode(e *colly.HTMLElement) *response.Match {
	match := response.Match{}
	name := e.ChildText("h3")

	if name == "" {
		return nil
	}

	match.Name = name
	status := "SCHEDULED"

	for key, node := range e.DOM.ChildrenFiltered(".detail").Nodes {
		if key == 0 {
			res := ""

			switch node.FirstChild.Data {
			case "Posticipata":
				status = "POSTPONED"
			case "Rinviata":
				status = "DEFERRED"
			case "b":
				status = "ENDED"
				res = node.FirstChild.FirstChild.Data
			case "span":
				status = "LIVE"
				res = node.FirstChild.FirstChild.FirstChild.Data
			}

			if res != "" {
				match.Result = &res
			}
		}

		if key == (len(e.DOM.ChildrenFiltered(".detail").Nodes) - 1) {
			match.Date = node.FirstChild.Data
		}
	}

	match.Status = status

	return &match
}
