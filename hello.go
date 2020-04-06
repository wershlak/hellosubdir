package hello

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/wershlak/hellosubdir/subdir"
)

// We need a dependency so let's add a simple config just to illustrate

type Config struct {
	Environment string `env:"HELLO_ENV" envDefault:"dev"`
}

// NewConfig returns a populated config.
func NewConfig() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}

// HelloWorld prints something
func HelloSubDir(w http.ResponseWriter, r *http.Request) {

	config, err := NewConfig()
	if err != nil {
		log.Fatalf("that didn't work", err)
	}

	message, err := subdir.GetMessage(config.Environment)
	if err != nil {
		log.Fatalf("that didn't work", err)
	}

	fmt.Fprintln(w, message)
}
