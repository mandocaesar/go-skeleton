package authentication

import message "github.com/go-skeleton/Base/message"

//Response : authentication base response
type Response struct {
	message.BaseMessage
	token string
}
