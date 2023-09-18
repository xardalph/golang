package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
)

type conf struct {
	Ipaddress string `env:"WEBIPADDRESS" env-default:"127.0.0.1"`
	WebPort   string `env:"WEBPORT" env-default:"5432"`

	DBHost     string `env:"MYSQL_HOST" env-default:"127.0.0.1"`
	DBPort     string `env:"MYSQL_PORT" env-default:"3658"`
	DBName     string `env:"DBNAME" env-default:"urlShortener"`
	DBUser     string `env:"MYSQL_USER" env-default:"root"`
	DBPassword string `env:"MYSQL_PASSWORD" env-default:"PASSFUCK"`
	DBHandle   *sql.DB
}

func (cfg conf) shortHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	if err != nil {
		fmt.Println("error when saving a document : ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func setupConfig() conf {
	var cfg conf

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	temp := mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPassword,
		Net:                  "tcp",
		Addr:                 cfg.DBHost + ":" + cfg.DBPort,
		DBName:               cfg.DBName,
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var db *sql.DB
	db, err = sql.Open("mysql", temp.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return cfg
}

func main() {
	var cfg = setupConfig()
	http.HandleFunc("/", cfg.shortHandler)

	log.Fatal(http.ListenAndServe(cfg.Ipaddress+":"+cfg.WebPort, nil))

}
