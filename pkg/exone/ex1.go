package exone

import (
	"encoding/json"
	"io"
	"time"

	"github.com/labstack/echo/v4"
)

/*
We receive notification via a webhook, for new candidates status change. The webhook contains only a little part of the resource we need in our platfomr.
For each request we need to retrieve the relevant infomration needed and store them in a database.
The operation of retrieval is syncronous.
Lately we noticed an increased number of errors with http status 429.
You are assigned to the team that needs to undestand and fix the problem
*/

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

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	var hp HookPayload[Application]
	if err := json.Unmarshal(body, &hp); err != nil {
		return err
	}

	// ......

}

func getCandidate(id string) Candidate {
	return Candidate{}
}

func getJob(id string) Job {
	return Job{}
}
