package dto

type UserExp struct {
	ID       string
	Username string
	ExpValue int
}

type LeaderBoard struct {
	Leaders []UserExp
}

type LimitsBoard struct {
	Limit uint64
}

type UpdateExp struct {
	ID       string
	ExpValue int
}

type AddExp struct {
	ID          string
	AddExpValue int
}
