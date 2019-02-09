package config

//Loggly : struct for loggly configuration
type Loggly struct {
	Token string
	Host  string
	Level string
	Tags  []string
}
