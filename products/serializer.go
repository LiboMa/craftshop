package products

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ProductSerializer struct {
	C        *gin.Context
	Products Products
}
type ProductsSerializer struct {
	C           *gin.Context
	ProductList []Products
}

type ProductResponse struct {
	ID          int     `json:"-"`
	Name        string  `json:"name"`
	Model       string  `json:"model"`
	Price       float32 `json:"price"`
	Description string  `json:"desc"`
	ImageUrl    string  `json:"image_url"`
	VideoUrl    string  `json:"video_url"`
	Capacity    int     `json: "capacity"`
	created_on  int     `json:"create_at"`
	created_by  string  `json:"created_by"`
	modified_on int     `json:"modified_on"`
	modified_by string  `json:"modified_by"`
	labels      string  `json:"labels"`
	state       int     `json:"state"`
}

func (p *ProductSerializer) Response() ProductResponse {

	response := ProductResponse{
		ID:          p.Products.id,
		Name:        p.Products.name,
		Price:       p.Products.price,
		Description: p.Products.description,
		ImageUrl:    p.Products.image_url,
		VideoUrl:    p.Products.video_url,
		Capacity:    p.Products.capacity,
	}

	return response

}
func (p *ProductsSerializer) Response() []ProductResponse {

	response := []ProductResponse{}

	for _, product := range p.ProductList {
		serializer := ProductSerializer{p.C, product}
		response = append(response, serializer.Response())
	}

	log.Println("Serialized done!", response)
	return response

}
