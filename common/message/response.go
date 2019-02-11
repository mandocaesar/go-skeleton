package common

//Response : authentication base response
type Response struct {
	BaseMessage
	Token string
}
