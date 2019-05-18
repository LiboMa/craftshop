package products

import (
	_ "fmt"

	"github.com/LiboMa/craftshop/common"
)

type Products struct {
	id          int     `id`
	name        string  `product_name`
	model       string  `model`
	price       float32 `price`
	description string  `desc`
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

func GetProductList() ([]Products, error) {

	_sql := "SELECT id, name, model, price, description, image_url, video_url, capacity from shop_products"
	rows, err := common.FetchAll(_sql)

	productList := make([]Products, 0)
	for rows.Next() {
		var p Products
		rows.Scan(&p.id, &p.name, &p.model, &p.price, &p.description, &p.image_url, &p.video_url, &p.capacity)
		productList = append(productList, p)
	}

	return productList, err

}

func GetProductListShort() ([]Products, error) {

	_sql := "SELECT id, name from shop_products"
	rows, err := common.FetchAll(_sql)

	productList := make([]Products, 0)
	for rows.Next() {
		var p Products
		rows.Scan(&p.id, &p.name)
		productList = append(productList, p)
	}

	return productList, err

}
