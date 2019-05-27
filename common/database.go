package common

import (
	"fmt"
	"log"
	"os"
	"reflect"

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
	db, err := sqlx.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	DB = db
	return DB
}

func FetchOne(query string, cond interface{}) (*sqlx.Rows, error) {

	rows, err := DB.Queryx(query, cond)

	if err != nil {
		log.Println(err)
	}
	return rows, err
}

func FetchAll(query string) (*sqlx.Rows, error) {

	log.Println(query)
	rows, err := DB.Queryx(query)
	if err != nil {
		log.Println(err)
	}

	return rows, err

}

func CreateQuery(q interface{}, tname string) (string, error) {

	log.Println("etnry query:", q)
	var err error

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		//tname := reflect.TypeOf(q).Name()
		fmt.Println(q)
		query := fmt.Sprintf("INSERT INTO `%s` VALUES(", tname)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.Int64:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.Float64:
				if i == 0 {
					query = fmt.Sprintf("%s%f", query, v.Field(i).Float())
				} else {
					query = fmt.Sprintf("%s, %f", query, v.Field(i).Float())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return "", err
			}
		}
		query = fmt.Sprintf("%s)", query)
		log.Println(query)
		fmt.Println(query)
		return query, nil

	}
	log.Println("unsupported type")
	//err := errors.New("unsupported type")
	return "", err

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
