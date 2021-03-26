package controllers

import (
	"net/http"

	"github.com/google/jsonapi"
)

type jobTypeEnum string
type experienceEnum string

const (
	entryLevel  experienceEnum = "entryLevel"
	midLevel    experienceEnum = "midLevel"
	seniorLevel experienceEnum = "seniorLevel"

	fullTime   jobTypeEnum = "full-time"
	partTime   jobTypeEnum = "part-time"
	freelance  jobTypeEnum = "freelance"
	contract   jobTypeEnum = "contract"
	temporary  jobTypeEnum = "temporary"
	internship jobTypeEnum = "internship"
)

// Job is available positions
// a company is looking to hire for
type Job struct {
	ID             int            `jsonapi:"primary,jobs"`
	Title          string         `validate:"required" jsonapi:"attr,title" db:"title"`
	Bio            string         `validate:"required" jsonapi:"attr,bio" db:"bio"`
	Facebook       string         `jsonapi:"attr,facebook" db:"facebook"`
	Twitter        string         `jsonapi:"attr,twitter" db:"twitter"`
	LinkedIn       string         `jsonapi:"attr,linked_in" db:"linked_in"`
	ApplyLink      string         `validate:"required" jsonapi:"attr,apply_link" db:"apply_link"`
	JobType        jobTypeEnum    `validate:"required" jsonapi:"attr,job_type" db:"job_type" enums:"full-time, part-time, freelance, contract, temporary, internship"`
	Xp             experienceEnum `validate:"required" jsonapi:"attr,xp" db:"xp" enums:"entryLevel, midLevel, seniorLevel"`
	Remote         bool           `validate:"required" jsonapi:"attr,remote" db:"remote"`
	JobDescription string         `validate:"required" jsonapi:"attr,job_description" db:"job_description"`
	City           string         `validate:"required" jsonapi:"attr,city" db:"city"`
	State          string         `validate:"required" jsonapi:"attr,state" db:"state"`
	CompanyName    string         `validate:"required" jsonapi:"attr,company_name" db:"company_name"`
	TechStack      string         `validate:"required" jsonapi:"attr,tech_stack" db:"tech_stack"`
}

// Jobs array
type Jobs []*Job

// First endpoint... naming will be tweaked/simplified later
func ShowAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs := Jobs{
		{
			ID:             1,
			Title:          "Test Title",
			Bio:            "lorem ipsum...",
			Facebook:       "facebook.com/borodev",
			Twitter:        "twitter.com/borodev",
			LinkedIn:       "linkedin.com/borodev",
			ApplyLink:      "apply.com/job",
			JobType:        "full-time",
			Xp:             "midLevel",
			Remote:         true,
			JobDescription: "lorem ipsum...",
			City:           "Murfreesboro",
			State:          "TN",
			CompanyName:    "BudgetBird",
			TechStack:      "Ember & Ruby On Rails",
		},
	}
	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(http.StatusOK)

	// This serializes the endpoint to match JSON API Specs. it's beautifully simple
	if err := jsonapi.MarshalPayload(w, jobs); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// CORS headers are automatically set under the
// API endpoints where this will most likely be moved.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
