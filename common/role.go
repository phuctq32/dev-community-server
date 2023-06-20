package common

type RoleType int

const (
	ADMINISTRATOR RoleType = iota
	MODERATOR
	MEMBER
)
