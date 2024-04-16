package config

type Configuration struct {
	LogLevel string
}

func NewConfiguration() *Configuration {
	return &Configuration{
		LogLevel: "info",
	}
}
