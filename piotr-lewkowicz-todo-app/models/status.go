package models

type Status int

const (
	Todo Status = iota
	InProgress
	Done
)

func (status Status) String() string {
	return [3]string{"Todo", "InProgress", "Done"}[status]
}
