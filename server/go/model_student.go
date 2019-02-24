package swagger

type Student struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	LastName string `json:"lastName,omitempty"`

	FatherName string `json:"fatherName,omitempty"`

	Email string `json:"email,omitempty"`

	Password string `json:"password,omitempty"`
}
