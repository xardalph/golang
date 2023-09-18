package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/ilyakaznacheev/cleanenv"
	"html/template"
	"log"
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type conf struct {
	Ipaddress string `env:"WEBIPADDRESS" env-default:"127.0.0.1"`
	WebPort   string `env:"WEBPORT" env-default:"5432"`

	DBHost     string `env:"MYSQL_HOST" env-default:"127.0.0.1"`
	DBPort     string `env:"MYSQL_PORT" env-default:"6379"`
	DBName     string `env:"DBNAME" env-default:"urlShortener"`
	DBUser     string `env:"MYSQL_USER" env-default:"root"`
	DBPassword string `env:"PASSWORD" env-default:""`
	Redis      *redis.Client
}

func (cfg conf) shortHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	if err != nil {
		fmt.Println("error when saving a document : ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func (cfg conf) ProvideFormHandler(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("addUrl.html")
	err := t.Execute(w, nil)
	if err != nil {
		return
	}

}
func (cfg conf) SaveUrlHandler(w http.ResponseWriter, r *http.Request) {
	urlToSave := r.FormValue("url")
	randomString := RandStringBytes(10)
	cfg.Redis.Set(randomString, urlToSave, 0)
	fmt.Printf("url to save : %s \nurl created : %s ", urlToSave, randomString)
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

/*
	 func renderHtml(w http.ResponseWriter, tmpl string) {
		err := templates.ExecuteTemplate(w, tmpl+".html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
*/
func setupConfig() conf {
	var cfg conf

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Redis = redis.NewClient(&redis.Options{
		Addr:     cfg.DBHost + ":" + cfg.DBPort,
		Password: cfg.DBPassword,
		DB:       0,
	})

	pong, err := cfg.Redis.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)
	return cfg
}

func main() {
	var cfg = setupConfig()
	err := cfg.Redis.Set("name", "Evan", 0).Err()
	if err != nil {
		println("err trying to set a value : ")
		println(err)
	}
	val, err := cfg.Redis.Get("gGOrQmuHRL").Result()
	if err != nil {
		println(err)
	}
	println("value is : " + val)

	http.HandleFunc("/", cfg.ProvideFormHandler)
	http.HandleFunc("/save", cfg.SaveUrlHandler)
	http.HandleFunc("/s/*", cfg.shortHandler)

	log.Fatal(http.ListenAndServe(cfg.Ipaddress+":"+cfg.WebPort, nil))

}
