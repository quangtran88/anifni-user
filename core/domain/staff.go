package domain

type StaffRole struct {
	Name       string
	Permission []string
}

type Staff struct {
	Active     bool
	Role       StaffRole
	Permission []string
}
