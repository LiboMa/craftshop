package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// TODO load db config from conf/config.json
//func Init(db_config.json) {
func Init() *sql.DB {
	db, err := sql.Open("mysql", "root:Desert_eagle@tcp(127.0.0.1:3306)/edushop")

	if err != nil {
		log.Fatal(err)
	}

	return db

}

type Products struct {
	id          int     `id`
	name        string  `product_name`
	model       string  `model`
	price       float32 `price`
	description string  `description`
	image_url   string  `image_url`
	video_url   string  `video_url`
	capacity    int     `capacity`
	created_on  int     `create_at`
	created_by  string  `created_by`
	modified_on int     `modified_on`
	modified_by string  `modified_by`
	labels      string  `labels`
	state       int     `state`
}

func main() {

	db := Init()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("db error")
	}

	p := Products{}

	//err = db.QueryRow("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id = ?", 1).Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
	//rows, err := db.Query("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id = ?", 1)
	rows, err := db.Query("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id < 2")
	//rst, err := db.Exec("select * from shop_products where id=1")

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	products := make([]Products, 0)

	for rows.Next() {

		rows.Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)

		products = append(products, p)
	}

	//fmt.Println(rows)
	fmt.Println(products)

	fmt.Println(p)

}
