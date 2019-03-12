package models

type Student struct {
	Id          int64           `json:"id,omitempty" sql:",pk,unique" xlsx:"-"`
	Name        string          `json:"name,omitempty" xlsx:"0"`
	Email       string          `json:"email,omitempty" xlsx:"1"`
	Password    string          `json:"password,omitempty" xlsx:"2"`
	TestId      int64           `json:"testId" xlsx:"-"`
	Test        *Test           `json:"-" xlsx:"-"`
	ScoreId     int64           `json:"-" xlsx:"-"`
	Score       *Score          `json:"-" xlsx:"-"`
	Answers     AnswerContainer `json:"answers" xlsx:"-"`
	AccessToken string          `json:"accessToken,omitempty" xlsx:"-"`
}
