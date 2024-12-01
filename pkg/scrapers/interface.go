package scrapers

import (
	"github.com/DanielJos/go-scraper-task-manager/pkg/results"
)

type Scraper interface {
	Scrape() (results.Result, error)
}