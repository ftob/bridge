package bridge

type Stat struct {
	ID string
	Data string
}

type Repository interface {
	Store(stat Stat) error
}