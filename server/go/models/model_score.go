package models

type Score struct {
	Id        int64 `json:"-"`
	Sum       int
	Base      int
	Reading   int

	Writing   int
	Listening int

	BaseAmount       int
	ReadingAmount    int
	WritingAmount    int
	ListeningAmount  int
	SumAmount        int
	RecommendedLevel string
}

type Result struct {
	Id    int
	Name  string
	Score *Score
}