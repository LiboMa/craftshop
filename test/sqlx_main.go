package main

import (
	//"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TODO load db config from conf/config.json
//func Init(db_config.json) {
var db *sqlx.DB

func Init() {
	var err error
	db, err = sqlx.Open("mysql", "root:Desert_eagle@tcp(127.0.0.1:3306)/edushop")

	if err != nil {
		log.Fatal(err)
	}

}

type Products struct {
	id          int     `db:"id"`
	name        string  `db:"name"`
	model       string  `db:"model"`
	price       float32 `db:"price"`
	description string  `db:"description"`
	image_url   string  `db:"image_url"`
	video_url   string  `db:"video_url"`
	capacity    int     `db:"capacity"`
	created_on  int     `db:"create_on"`
	created_by  string  `db:"created_by"`
	modified_on int     `db:"modified_on"`
	modified_by string  `db:"modified_by"`
	labels      string  `db:"labels"`
	state       int     `db:"state"`
}

func main() {

	Init()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("db error")
	}

	//p := Products{}

	//err = db.QueryRow("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id = ?", 1).Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
	//rows, err := db.Queryx("select id, name, model, price, description, image_url, video_url, capacity from shop_products")
	rows, err := db.Queryx("select * from shop_products")
	//	rows, err := db.Queryx("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id < 2")
	//rows, err := db.Queryx("select id, name, model, price from shop_products where id = 1")
	//rows, err := db.Queryx("select * from shop_products")
	//rst, err := db.Exec("select * from shop_products where id=1")

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(rows)

	products := make([]Products, 0)

	var p Products
	//err = db.Get(&p, "SELECT * from shop_products")
	for rows.Next() {

		//rows.Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
		rows.StructScan(&p)

		products = append(products, p)
	}

	//fmt.Println(rows)
	fmt.Println(products)

	fmt.Println(p)

}
