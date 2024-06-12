package config

// PostgreSQL driver
import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

type ConfigDB struct {
	DbName   string `toml:"dbName"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

// init data base
func InitDB() {
	var config ConfigDB
	var err error
	_, err = toml.DecodeFile("config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		config.User, config.Password, config.Host, config.Port, config.DbName)

	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("failed to conncet db", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("failed to ping db", err.Error())
	}
	log.Println("connected")
}
