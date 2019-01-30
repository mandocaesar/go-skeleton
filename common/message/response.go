package common

//Response : authentication base response
type Response struct {
	BaseMessage
	token string
}
