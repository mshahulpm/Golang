package handlers

import (
	"context"
	"sample-server/cmd/internal/db"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID primitive.ObjectID  `json:"_id" bson:"_id"`
	CreatedAt time.Time    `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt" bson:"updated_at"`
	Title string           `json:"title" bson:"title"`
}

type ErrorResponse struct {
	FailedField string 
	Tag string 
	Value string 

}

func ValidateProductStruct(p Product) []*ErrorResponse{
	var errors []*ErrorResponse 

	validate := validator.New() 

	err := validate.Struct(p) 

	if err != nil {
		for _,err := range err.(validator.ValidationErrors) {
			var element ErrorResponse 
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag() 
			element.Value = err.Param() 

			errors = append(errors, &element) 
		}
	}
  return nil 
}

type Collection string 

const (
	productsCollection Collection = "product" 
)

func CreateProduct(c *fiber.Ctx) error {
	product := Product  {
		ID:  primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// err := c.BodyParser(&product) 

	// if err != nil {
	// 	return err 
	// }
	//  the above code is the same for below code 

	if err := c.BodyParser(&product) ; err !=nil {
		return err 
	}

	errors := ValidateProductStruct(product) 

	if errors != nil {
		return c.JSON(errors) 
	}

	client,err := db.GetMongoClient()

	if err!=nil { 
		return err 
	}

	collection := client.Database(db.DataBase).Collection(string(productsCollection)) 

	_,err = collection.InsertOne(context.TODO(),product) 

	if err !=nil {
		return err 
	}

	return c.JSON(product)
}