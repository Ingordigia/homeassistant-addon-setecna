package helpers

type Season int64

const (
	Summer Season = iota
	Winter
)

func (s Season) String() string {
	switch s {
	case Summer:
		return "summer"
	case Winter:
		return "winter"
	}
	return "unknown"
}
