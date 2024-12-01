package job

import (
	"fmt"

	"github.com/DanielJos/go-scraper-task-manager/pkg/results"
	"github.com/DanielJos/go-scraper-task-manager/pkg/scrapers"
)

type Job struct {
	ID int
	URL string
	Tries int
	Status string
	Scraper scrapers.Scraper
}

type Manager struct {
	Jobs []Job
	jobsChannel chan Job
	resultsChannel chan results.Result
}

func NewManager () Manager {
	return Manager{
		Jobs: []Job{},
		jobsChannel: make(chan Job, 10),
		resultsChannel: make(chan results.Result, 10),
	}
}

func (h *Manager) ListJobs() []Job {
	return h.Jobs
}


func (h *Manager) CreateJob(url string, scraper scrapers.Scraper) error {
	if url == "" {
		return fmt.Errorf("no job url provided")
	}

	// Generate ID
	id := len(h.Jobs)
	
	j := Job{
		ID: id,
		URL: url,
		Tries: 0,
	}
	h.Jobs = append(h.Jobs, j)

	h.jobsChannel <- j

	return nil
}

// Await jobs in the Jobs slice, when available spawn a worker where numberOfWorkers >= concurrentJobs
func (h *Manager) StartWorkerPool(workers int, resultsStore *results.ResultsStore) {
	for i:=0; i<workers; i++ {
		go h.worker()
	}

	// Handle results
	for result := range h.resultsChannel {
		resultsStore.Add(result)
	}
}

func (h *Manager) worker() {
	for job := range h.jobsChannel {
		result, e := h.process(job)
		if e != nil {
			job.Tries++
			if job.Tries < 3 {
				job.Status = "RETRYING"
				h.jobsChannel <- job
			}
			job.Status = "FAILED"
		}

		h.resultsChannel <- result
	}
}

func (h *Manager) process(job Job) (results.Result, error) {

	fmt.Printf("Processing %d", job.ID) // TODO: Add logging proper

	result, e := job.Scraper.Scrape()
	if e != nil {
		return results.Result{}, e
	}

	return result, nil
}