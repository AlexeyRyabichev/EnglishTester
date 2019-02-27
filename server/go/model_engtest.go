package swagger

import (
	"database/sql/driver"
	"encoding/json"
)

type Test struct {
	Id        int        `json:"id,omitempty"`
	Questions []Question `json:"questions" sql:",notnull"`
	Answers   []string   `json:"answers" sql:",array"`
}

type Question struct {
	Section int    `json:"section"`
	Text    string `json:"text"`
}

func (m Question) Value() (driver.Value, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
