package common

var ReqUser string = "user"

type Requester interface {
	GetUserId() string
}
