package models

type Test struct {
	Id               int             `json:"id,omitempty"`
	BaseQuestions    []Question      `json:"baseQuestions" sql:",notnull"`
	ReadingQuestions *Reading        `json:"reading"`
	Writing          string          `json:"writing"`
	Answers          AnswerContainer `json:"answers"`
}

type Question struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	OptionA  string `json:"optionA"`
	OptionB  string `json:"optionB"`
	OptionC  string `json:"optionC"`
	OptionD  string `json:"optionD"`
}

type Reading struct {
	Question     string     `json:"question"`
	BaseQuestion []Question `json:"questions"`
}

type AnswerContainer struct {
	Base    []Answer `json:"base"`
	Reading []Answer `json:"reading"`
}

type Answer struct {
	Id     int    `json:"id"`
	Answer string `json:"answer"`
}

type ProxyQuestions struct {
	BaseQuestions    []Question `json:"baseQuestions" sql:",notnull"`
	ReadingQuestions *Reading   `json:"reading"`
	Writing          string     `json:"writing"`
}

//func (m Question) Value() (driver.Value, error) {
//	b, err := json.Marshal(m)
//	if err != nil {
//		return nil, err
//	}
//	return string(b), nil
//}
