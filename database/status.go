package database

type status int

const (
	New status = iota
	Processing
	Processed
)

func (s status) String() string {
	switch s {
	case New:
		return "new"
	case Processing:
		return "processing"
	case Processed:
		return "processed"
	default:
		return ""
	}
}
