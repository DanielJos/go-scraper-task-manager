package scrapers

import (
	// "net/http"
	"github.com/DanielJos/go-scraper-task-manager/pkg/results"
)

type Standard struct {
	
}

func NewStandard() Standard {
	return Standard{}
}

func (s *Standard) Scrape() (results.Result, error) {

	return results.Result{}, nil
}