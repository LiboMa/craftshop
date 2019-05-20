package products

import (
	"database/sql"
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

	if err != nil {
		log.Println(err)
	}

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

func UpdateProductByID(p *Products) (sql.Result, error) {

	_sql := fmt.Sprintf("UPDATE shop_products SET name='%s', model='%s', price=%f, description='%s', image_url='%s', video_url='%s', capacity=%d, created_on=%d, created_by='%s', modified_on=%d, modified_by='%s',labels='%s', state=%d WHERE id=%d",
		p.Name, p.Model, p.Price, p.Description, p.Image_url, p.Video_url, p.Capacity,
		p.Created_on, p.Created_by, p.Modified_on, p.Modified_by,
		p.Labels, p.State, p.ID,
	)
	db := common.Getdb()
	log.Println(_sql)
	result, err := db.Exec(_sql)
	// result, err := db.NamedExec(`UPDATE shop_products SET name:name, model=:model, price=:price, description=:description,
	//  image_url=:image_url, video_url=:video_url, capacity=:capacity,
	//  created_on=:created_on created_by=:created_by, modified_on=:modified_on,modified_by=:modified_by,
	//  labels=:labels,state=:state WHERE id=:id`, p)

	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

func DeleteProductByID(p *Products) (sql.Result, error) {

	_sql := fmt.Sprintf("DELETE FROM shop_products WHERE id=%d", p.ID)
	db := common.Getdb()
	log.Println(_sql)
	result, err := db.Exec(_sql)
	// result, err := db.NamedExec(`UPDATE shop_products SET name:name, model=:model, price=:price, description=:description,
	//  image_url=:image_url, video_url=:video_url, capacity=:capacity,
	//  created_on=:created_on created_by=:created_by, modified_on=:modified_on,modified_by=:modified_by,
	//  labels=:labels,state=:state WHERE id=:id`, p)

	if err != nil {
		log.Fatal(err)
	}
	return result, err
}
