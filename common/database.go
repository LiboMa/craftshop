package common

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TODO load db config from conf/config.json
//func Init(db_config.json) {
type Database struct {
	*sqlx.DB
}

var DB *sqlx.DB

func InitDB() *sqlx.DB {
	db, err := sqlx.Open("mysql", "root:Desert_eagle@tcp(127.0.0.1:3306)/edushop")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}
	DB = db
	return DB
}

func FetchOne(query string, cond interface{}) *sqlx.Row {

	return DB.QueryRowx(query, cond)
}

func FetchAll(sql_query string) (*sqlx.Rows, error) {

	rows, err := DB.Queryx(sql_query)
	if err != nil {
		log.Fatal(err)
	}

	return rows, err

}

func Exec(query string) {

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func Getdb() *sqlx.DB {
	return DB
}
