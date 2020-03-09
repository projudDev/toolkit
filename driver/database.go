package driver

import (
	"fmt"
	"os"

	"database/sql"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
)

type DB struct {
	SQL       *sql.DB
	RedisConn redis.Conn
}

var dbConn = &DB{}

func NewConnectionMysql() (*DB, error) {
	dbSource, err := GetDBSource()
	if err != nil {
		return nil, err
	}
	url, ok := os.LookupEnv("CLEARDB_DATABASE_URL")
	if ok {
		return connectionMysql(url)
	}
	return connectionMysql(dbSource)
}
func connectionMysql(url string) (*DB, error) {
	d, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	dbConn.SQL = d
	return dbConn, err
}
func ConnectRedis() (*DB, error) {
	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", ":6379")
		},
	}
	dbConn.RedisConn = redisPool.Get()
	return dbConn, nil
}

func GetDBSource() (string, error) {
	type database struct {
		Server   string
		Port     string
		Database string
		User     string
		Password string
	}
	type output struct {
		Directory string
		Format    string
	}
	type Config struct {
		Output   output
		Database database
	}
	var conf Config
	if _, err := toml.DecodeFile("./../config.toml", &conf); err != nil {
		return "", err
	}
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Server,
		conf.Database.Port,
		conf.Database.Database,
	)
	return dbSource, nil
}
