package Roles

const (
	Student = iota
	Teacher
	Admin
)

var rolesText = map[int]string{
	Student: "student",
	Teacher: "teacher",
	Admin:   "admin",
}

func RolesText(role int) string {
	return rolesText[role]
}
