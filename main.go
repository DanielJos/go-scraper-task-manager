package main

import (
	"net/http"
	"time"
	"html"
	"fmt"

	"github.com/DanielJos/go-scraper-task-manager/pkg/job"
	// "github.com/DanielJos/go-scraper-task-manager/pkg/results"
)

func main () {
	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	jobHandler := job.NewHandler()
	_ = jobHandler

	jobsManager := job.NewManager()
	jobHandler.Manager = jobsManager

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%q", html.EscapeString(r.URL.Path))
	})

	// Jobs
	http.HandleFunc("/jobs", jobHandler.ListJobs)
	http.HandleFunc("/job", jobHandler.CreateJob)

	fmt.Println("Running...")

	if e := server.ListenAndServe(); e != nil {
		panic(e)
	}
}