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
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	VideoUrl    string  `json:"video_url"`
	Capacity    int     `json: "capacity"`
	Created_on  int64   `json:"create_at"`
	Created_by  string  `json:"created_by"`
	Modified_on int64   `json:"modified_on"`
	Modified_by string  `json:"modified_by"`
	Labels      string  `json:"labels"`
	State       int     `json:"state"`
}

func (p *ProductSerializer) Response() ProductResponse {

	response := ProductResponse{
		ID:          p.Products.ID,
		Name:        p.Products.Name,
		Model:       p.Products.Model,
		Price:       p.Products.Price,
		Description: p.Products.Description,
		ImageUrl:    p.Products.Image_url,
		VideoUrl:    p.Products.Video_url,
		Capacity:    p.Products.Capacity,
		Created_on:  p.Products.Created_on,
		Created_by:  p.Products.Created_by,
		Modified_on: p.Products.Modified_on,
		Modified_by: p.Products.Modified_by,
		Labels:      p.Products.Labels,
		State:       p.Products.State,
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
