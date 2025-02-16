package swaggergenerated

import (
	"log"
	"sandbox/http/swagger_generated/models"
	"sandbox/http/swagger_generated/restapi"
	"sandbox/http/swagger_generated/restapi/operations"
	"sandbox/http/swagger_generated/restapi/operations/products"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/samber/lo"
)

func CreateProductHandler(p products.CreateProductParams) middleware.Responder {
	return products.NewCreateProductCreated().WithPayload(&models.EchoExampleProduct{
		ID:       gofakeit.Int64(),
		Interest: gofakeit.Float64(),
		IsActive: gofakeit.Bool(),
		Title:    gofakeit.Sentence(3),
		Total:    gofakeit.Int64(),
	})
}

func GetProductHandler(p products.GetProductParams) middleware.Responder {
	return products.NewGetProductOK().WithPayload(&models.EchoExampleProduct{
		ID:       p.ID,
		Interest: gofakeit.Float64(),
		IsActive: gofakeit.Bool(),
		Title:    gofakeit.Sentence(3),
		Total:    gofakeit.Int64(),
	})
}

func ListProductsHandler(p products.ListProductsParams) middleware.Responder {
	return products.NewListProductsOK().WithPayload(lo.Times(3, func(_ int) *models.EchoExampleProduct {
		return &models.EchoExampleProduct{
			ID:       gofakeit.Int64(),
			Interest: gofakeit.Float64(),
			IsActive: gofakeit.Bool(),
			Title:    gofakeit.Sentence(3),
			Total:    gofakeit.Int64(),
		}
	}))
}

func NewServer() (*restapi.Server, func() error) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewEchoExampleAPI(swaggerSpec)

	api.ProductsCreateProductHandler = products.CreateProductHandlerFunc(CreateProductHandler)
	api.ProductsGetProductHandler = products.GetProductHandlerFunc(GetProductHandler)
	api.ProductsListProductsHandler = products.ListProductsHandlerFunc(ListProductsHandler)

	server := restapi.NewServer(api)

	server.ConfigureAPI()
	return server, server.Shutdown
}
