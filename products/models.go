package products

import (
	"fmt"
	_ "fmt"
	"log"

	"github.com/LiboMa/craftshop/common"
)

type Products struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Model       string  `db:"model"`
	Price       float64 `db:"price"`
	Description string  `db:"description"`
	Image_url   string  `db:"image_url"`
	Video_url   string  `db:"video_url"`
	Capacity    int     `db:"capacity"`
	Created_on  int64   `db:"created_on"`
	Created_by  string  `db:"created_by"`
	Modified_on int64   `db:"modified_on"`
	Modified_by string  `db:"modified_by"`
	Labels      string  `db:"labels"`
	State       int     `db:"state"`
}

func GetProductList() ([]Products, error) {

	_sql := "SELECT id, name, model, price, description, image_url, video_url, capacity from shop_products"
	rows, err := common.FetchAll(_sql)

	productList := make([]Products, 0)
	for rows.Next() {
		var p Products
		rows.Scan(&p.ID, &p.Name, &p.Model, &p.Price, &p.Description, &p.Image_url, &p.Video_url, &p.Capacity)
		productList = append(productList, p)
	}

	return productList, err

}

//func GetProduct(p Products) ProductModel {
//}

func GetProductByID(id int) (Products, error) {
	//_sql := "SELECT id, name, price, model, description, image_url, video_url, capacity FROM shop_products WHERE id = ?"
	_sql := "SELECT * FROM shop_products WHERE id = ? Limit 1"

	db := common.Getdb()
	var p Products
	err := db.Get(&p, _sql, id)

	// rows, err := common.FetchOne(_sql, p.ID)
	// defer rows.Close()

	// for rows.Next() {
	// 	rows.StructScan(p)
	// }
	fmt.Printf("p: %T, %v\n", p, p)
	return p, err
}

func CreateProduct(p *Products) {

	_sql, err := common.CreateQuery(*p, "shop_products")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(_sql)
	db := common.Getdb()
	db.Exec(_sql)
}
