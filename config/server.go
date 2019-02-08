package config

//Server : sturct hold model for server confguration
type Server struct {
	Mode            string
	Addr            string
	Environment     string
	LogDuration     int
	ShutdownTimeout int
	BaseURL         string
	ClientURL       string
	Name            string
}
