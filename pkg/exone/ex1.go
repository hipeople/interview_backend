package exone

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

// We receive notification via a webhook, for candidates changes. The webhook contains only a little part of the resource we need in our platform.
// For each request we need to retrieve the relevant information needed and store them in a database.
// The operation of retrieval is syncronous.
// Lately we noticed an increased number of errors with http status 429.
// You are assigned to the team that needs to undestand and fix the problem

type Hook struct {
	ID       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
	Target   string `json:"target"`
}

type HookPayload[T any] struct {
	Hook          Hook   `json:"hook"`
	LinkedAccount string `json:"linked_account"`
	Data          T      `json:"data"`
}

type BaseApplication[C string, J string, St string] struct {
	ID           string    `json:"id"`
	RemoteID     string    `json:"remote_id"`
	Candidate    C         `json:"candidate"`
	Job          J         `json:"job"`
	AppliedAt    time.Time `json:"applied_at"`
	RejectedAt   time.Time `json:"rejected_at"`
	Source       string    `json:"source"`
	CreditedTo   any       `json:"credited_to"`
	CurrentStage St        `json:"current_stage"`
	RejectReason any       `json:"reject_reason"`
	RemoteData   any       `json:"remote_data"`
}

type Application = BaseApplication[string, string, string]

type Candidate struct {
	ID       string
	FullName string
	Email    string
}

type Job struct {
	ID          string
	Name        string
	Description string
}

func Webhook(c echo.Context) error {
	var hp HookPayload[Application]
	if err := json.NewDecoder(c.Request().Body).Decode(&hp); err != nil {
		return err
	}

	candidate, err := getCandidate(hp.Data.Candidate)
	if err != nil {
		return err
	}

	job, err := getJob(hp.Data.Job)
	if err != nil {
		return err
	}

	if err := storeCandidate(candidate); err != nil {
		return err
	}

	if err := storeJob(job); err != nil {
		return err
	}

	return nil
}

func getCandidate(id string) (*Candidate, error) {
	// retrieve the candidate information from external service
	return &Candidate{}, nil
}

func getJob(id string) (*Job, error) {
	// retrieve the job information from external service
	return &Job{}, nil
}

func storeCandidate(c *Candidate) error {
	// store stuff in database
	return nil
}

func storeJob(j *Job) error {
	// store stuff in database
	return nil
}
