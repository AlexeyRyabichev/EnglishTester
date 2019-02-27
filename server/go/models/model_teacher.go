package models

type Teacher struct {
	Id          int64  `json:"id,omitempty"`
	Role        string `json:"role,omitempty"`
	Email       string `json:"login,omitempty"`
	Password    string `json:"password,omitempty"`
	Name        string `json:"name,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	FatherName  string `json:"fatherName,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
}
