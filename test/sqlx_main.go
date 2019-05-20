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
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Model       string  `db:"model"`
	Price       float32 `db:"price"`
	Description string  `db:"description"`
	Image_url   string  `db:"image_url"`
	Video_url   string  `db:"video_url"`
	Capacity    int     `db:"capacity"`
	Created_on  int     `db:"created_on"`
	Created_by  string  `db:"created_by"`
	Modified_on int     `db:"modified_on"`
	Modified_by string  `db:"modified_by"`
	Labels      string  `db:"labels"`
	State       int     `db:"state"`
}

func main() {

	Init()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		fmt.Println("db error")
	}

	//p := Products{}
	var p Products

	db.Get(&p, "select  * from shop_products where id = ? Limit 1", 1)

	fmt.Println("---->", p)

	//err = db.QueryRow("select id, name, model, price, description, image_url, video_url, capacity from shop_products where id = ?", 1).Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
	//err = db.QueryRow("select id, name, model, price, description, image_url, video_url, capacity, created_on, created_by  from shop_products where id = ?", 1).Scan(&p.ID, &p.Name, &p.Model, &p.Price, &p.Description, &p.Image_url, &p.Video_url, &p.Capacity, &p.Created_on, &p.Created_by)
	//fmt.Println(p)
	//err = db.QueryRowx("select * from shop_products where id = ?", 1).Scan(&p.ID, &p.Name, &p.Model, &p.Price, &p.Description, &p.Image_url, &p.Video_url, &p.Capacity)
	//fmt.Println(p)
	//rows, err := db.Queryx("select id, name, model, price, description, image_url, video_url, capacity from shop_products")
	//rows, err := db.Queryx("select id, name, model, price, description, image_url, video_url, capacity from shop_products")
	//fmt.Println(p)
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

	//var p Products
	//err = db.Get(&p, "SELECT * from shop_products")
	for rows.Next() {

		//rows.Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
		err = rows.StructScan(&p)

		products = append(products, p)
	}

	//fmt.Println(rows)
	fmt.Println(products)

	//fmt.Println(p)

	//var input_product string = `{"name":"english_A2","model":"","price":99,"desc":"english lessons for children age of 3-6","image_url":"http://s3.edushop.com/static/images/en_a2.jepg","video_url":"http://s3.edushop.com/static/images/en_a2.jepg","Capacity":99}`

	//fmt.Println(input_product)

}
