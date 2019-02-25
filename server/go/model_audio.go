package swagger

type Audio struct {
	Id        int64  `json:"id,omitempty"`
	StudentId int64  `json:"student,omitempty"`
	Path      string `json:"path,omitempty"`
}
