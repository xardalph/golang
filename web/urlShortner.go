package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"os"
)

type ConfigDatabase struct {
	Port     string `env:"PORT" env-default:"5432"`
	Host     string `env:"HOST" env-default:"localhost"`
	Name     string `env:"USERNAME" env-default:"postgres"`
	User     string `env:"USER" env-default:"user"`
	Password string `env:"PASSWORD" end-default:"motdepasse"`
}

func shortHandler(w http.ResponseWriter, r *http.Request) {

}

func setupConfig() {
	var cfg ConfigDatabase

	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(cfg)
}
func main() {
	setupConfig()
	http.HandleFunc("/", shortHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
