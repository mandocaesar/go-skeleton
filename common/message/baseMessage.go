package common

//BaseMessage base message sturcture
type BaseMessage struct {
	Data    interface{}
	Code    int
	Message string
}
