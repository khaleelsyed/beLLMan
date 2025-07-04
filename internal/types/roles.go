package types

type Role int

const (
	RoleSystem Role = iota
	RoleAssistant
	RoleUser
)

var roleName = map[Role]string{
	RoleSystem:    "system",
	RoleAssistant: "assistant",
	RoleUser:      "user",
}

func (role Role) String() string {
	return roleName[role]
}
