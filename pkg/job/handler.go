package job

import (
	"net/http"
	"fmt"
	"github.com/DanielJos/go-scraper-task-manager/pkg/scrapers"
)

type Handler struct {
	Manager Manager
}

func NewHandler() Handler {
	return Handler{}
}

func (h *Handler) ListJobs(w http.ResponseWriter, r *http.Request) {
	jobs := h.Manager.ListJobs()
	fmt.Fprintf(w, "%+v", jobs)
}

func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	jobURL := r.URL.Query().Get("URL")

	if jobURL != "" {
		http.Error(w, "insufficient params", 400)
	}

	scraper := scrapers.NewStandard()

	h.Manager.CreateJob(jobURL, &scraper)

	fmt.Fprintf(w, "Job added successfully.")
}