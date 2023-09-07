package extwo

/*
	In our candidate view we compute the assessment status using 4 boolean values. Those value are stores in our database, then we have a frontend
	algo that computes a status using the values.
	Product wants two new features to be delivered: filtering and sorting on those statuses, we also need to keep supporting pagination.
	How can we approach the problem? Which data structure do we need?
*/

type Status string

const (
	StatusSubmitted  Status = "submitted"
	StatusExpired    Status = "expired"
	StatusInProgress Status = "in-progress"
	StatusDisabled   Status = "disabled"
)

type Row struct {
	SubmittedAt int
	ExpiredAt   int
	IsActive    bool
}
