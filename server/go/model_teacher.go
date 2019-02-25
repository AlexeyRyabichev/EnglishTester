package swagger

type Teacher struct {
	Id          int64  `json:"id,omitempty"`
	Role        string `json:"role,omitempty"`
	Email       string `json:"login,omitempty"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name,omitempty"`
	LastName    string `json:"surname,omitempty"`
	FatherName  string `json:"father	name,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
