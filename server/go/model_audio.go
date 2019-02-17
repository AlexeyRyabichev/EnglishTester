package swagger

type Audio struct {
	Id int `json:"id,omitempty"`
	Student Student `json:"student,omitempty"`
	Path string `json:"path,omitempty"`
}
