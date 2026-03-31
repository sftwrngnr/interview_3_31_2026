package config

type Config struct {
	Port            int    `yaml:"port"`
	BindAddress     string `yaml:"bind_address"`
	EnableTelemetry bool   `yaml:"enable_telemetry"`
}
