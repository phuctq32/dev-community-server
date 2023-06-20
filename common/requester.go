package common

const ReqUser string = "user"

type Requester interface {
	GetUserId() string
	GetRoleType() RoleType
}
