package authentication

import (
	fmt "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mandocaesar/go-skeleton/domain/authentication/messages"
	"gopkg.in/go-playground/validator.v8"
)

//Controller : controller form authentication
type Controller struct {
	_service *Service
}

//New
func NewController(service *Service) (*Controller, error) {
	return &Controller{_service: service}, nil
}

//Login : Login controller
func (c *Controller) Login(ctx *gin.Context) {

	var request messages.LoginRequest
	var errors []string

	if err := ctx.ShouldBindWith(&request, binding.JSON); err != nil {
		ve, ok := err.(validator.ValidationErrors)
		if ok {
			for _, v := range ve {
				msg := fmt.Sprintf("%s is %s", v.Field, v.Tag)
				if v.Tag == "len" {
					msg = fmt.Sprintf("%s %s should be %s", v.Field, v.Tag, v.Param)
				}
				errors = append(errors, msg)
			}
		} else {
			errors = append(errors, err.Error())
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "errors": errors})
		return
	}

	result := c._service.Authenticate(request.Username)

	ctx.JSON(http.StatusOK, result)
}
