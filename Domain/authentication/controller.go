package authentication

//Controller : controller form authentication
type Controller struct {
	_service Service
}

//New
func New(service Service) (*Controller, error) {
	return &Controller{_service: service}, nil
}

//Login : Login controller
func (c *Controller) Login(input string) string {
	return c._service.Authenticate(input)
}
