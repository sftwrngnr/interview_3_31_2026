package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"interview_3_31_2026/config"
	"interview_3_31_2026/handlers"
	//"interview_3_31_2026/telemetry"
)

// ReadConfig returns the configuration information or an error if the config file
// doesn't exist

func ReadConfig(configPath string) (*config.Config, error) {
	cfg := &config.Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// CreateRouter returns a gorilla mux router that establishes the endpoints for the API.
// Handlers are located in the handlers directory
func CreateRouter(ctx context.Context) *mux.Router {
	ctx.Value("logger").(*zap.SugaredLogger).Info("Create router")
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.HomeHandler)
	router.HandleFunc("/mainendpoint", handlers.MainEndpointHandler)
	return router
}

// CreateContext returns a context with a named reference to the sugared logger and config data
// that are passed into the function.
func CreateContext(sl *zap.SugaredLogger, cfg *config.Config) context.Context {
	rval := context.Background()
	rval = context.WithValue(rval, "logger", sl)
	fmt.Printf("%v\n", cfg)
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
	tctx := CreateContext(sugar, cfg)
	//telem := telemetry.NewTelemetryProvider(tctx, cfg)
	//pSpan := telem.StartTracer("main")
	//defer pSpan.End()
	defer logger.Sync()
	r := CreateRouter(tctx)
	sugar.Info("Starting interview microservice on port: " + strconv.Itoa(cfg.Port) + " host: " + cfg.BindAddress)

	//http.Handle("/", r)
	http.ListenAndServe(cfg.BindAddress+":"+strconv.Itoa(cfg.Port), r)
}
