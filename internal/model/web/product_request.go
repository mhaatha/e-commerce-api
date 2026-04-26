package web

type CreateProductRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=150"`
	Quantity    int    `json:"quantity" validate:"required,number,minimum=0"`
	Price       int    `json:"price" validate:"required,number,minimum=0"`
	Description string `json:"description" validate:"min=3,max=2000"`
}
