package types

type EndpointStatus string

const (
	Active  EndpointStatus = "active"
	Paused  EndpointStatus = "paused"
	Deleted EndpointStatus = "deleted"
)

func (s EndpointStatus) String() string {
	return string(s)
}

func (s EndpointStatus) IsValid() bool {
	switch s {
	case Active, Paused, Deleted:
		return true
	}
	return false
}
