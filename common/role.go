package common

type RoleType int

const (
	Administrator RoleType = iota
	Moderator
	Member
)
