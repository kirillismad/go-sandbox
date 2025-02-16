package swaggergenerated

import (
	"os"
	"sandbox/http/swagger_gen/models"
	"sandbox/http/swagger_gen/restapi/operations/files"
	"sandbox/http/swagger_gen/restapi/operations/products"

	"github.com/brianvoe/gofakeit/v7"
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

func UpdateProductHandler(p products.UpdateProductParams) middleware.Responder {
	return products.NewUpdateProductOK().WithPayload(&models.EchoExampleProduct{
		ID:       p.ID,
		Interest: gofakeit.Float64(),
		IsActive: gofakeit.Bool(),
		Title:    gofakeit.Sentence(3),
		Total:    gofakeit.Int64(),
	})
}

func DeleteProductHandler(p products.DeleteProductParams) middleware.Responder {
	return products.NewDeleteProductNoContent()
}

func DownloadFileHandler(f files.DownloadFileParams) middleware.Responder {
	file, err := os.Open(os.Getenv("IMAGE_FILE"))
	if err != nil {
		return files.NewDownloadFileInternalServerError().WithPayload(map[string]string{"error": err.Error()})
	}
	return files.NewDownloadFileOK().WithPayload(file).WithContentType("image/png").WithContentDisposition("inline; filename=image.png")
}
