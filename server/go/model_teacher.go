package swagger

type Teacher struct {
	Id         int64  `json:"id,omitempty"`
	Role       string `json:"role,omitempty"`
	Login      string `json:"login,omitempty"`
	Password   string `json:"password,omitempty"`
	Name       string `json:"name,omitempty"`
	SurName    string `json:"surname,omitempty"`
	FatherName string `json:"fathername,omitempty"`
}
