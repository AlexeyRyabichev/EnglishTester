package Roles

type Role int

const (
	Student = iota
	Teacher
	Admin
)

func (role Role) String() string {
	return [...]string{"student", "teacher", "admin"}[role]
}
