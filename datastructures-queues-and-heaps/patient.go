package queue

type SeverityStatus int

const (
	Mild SeverityStatus = iota
	Moderate
	Critical
)

type Patient struct {
	name     string
	severity SeverityStatus

	// required by heap.Interface
	index int
}
