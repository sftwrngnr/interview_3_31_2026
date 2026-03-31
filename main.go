package interview_3_31_2026

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Port        int    `yaml:"port"`
	BindAddress string `yaml:"bind_address"`
}

// ReadConfig returns the configuration information or an error if the config file
// doesn't exist

func ReadConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}
	return config, nil
}

// CreateRouter returns a gorilla mux router that establishes the endpoints for the API.
// Handlers are located in the handlers directory
func CreateRouter(ctx context.Context) *mux.Router {
	ctx.Value("Logger").(*zap.Logger).Sugar().Info("Create router")
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/mainendpoint", MainEndpoint)
	return router
}

// CreateContext returns a context with a named reference to the sugared logger and config data
// that are passed into the function.
func CreateContext(sl *zap.SugaredLogger, cfg *Config) context.Context {
	rval := context.Background()
	rval = context.WithValue(rval, "logger", sl)
	rval = context.WithValue(rval, "config", cfg)
	return rval
}

func main() {
	cfg, cerr := ReadConfig("config.yaml")
	if cerr != nil {
		panic(cerr)
	}
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	sugar := logger.Sugar()
	ctx := CreateContext(sugar, cfg)
	defer logger.Sync()
	r := CreateRouter(ctx)
	sugar.Info("Starting interview microservice")
	http.Handle("/", r)
}
