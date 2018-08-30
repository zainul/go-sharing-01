package internal

import (
	"database/sql"
	"fmt"

	"github.com/BurntSushi/toml"
)

var conf Config

// DB is map instance of db
var DB map[string]*sql.DB

// Config ...
type Config struct {
	Database database
}

type database struct {
	Server   string
	Port     string
	Name     string
	User     string
	Password string
	Provider string
}

func init() {
	dbConnectionMap := make(map[string]*sql.DB)

	if _, err := toml.DecodeFile("./configs/db.toml", &conf); err != nil {
		panic(err)
	}

	KAIMaster, err := sql.Open(conf.Database.Provider,
		fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Server,
			conf.Database.Port,
		),
	)

	if err != nil {
		panic(err)
	}

	dbConnectionMap["db_kai.master"] = KAIMaster

	db := KAIMaster

	errPing := db.Ping()

	if errPing != nil {
		fmt.Println("database un reachable ...")
		panic(errPing)
	}

	fmt.Println("database is running ....")
	DB = dbConnectionMap
}

// GetDB ...
func GetDB() map[string]*sql.DB {
	return DB
}
