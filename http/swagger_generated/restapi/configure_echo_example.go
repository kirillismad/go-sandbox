// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"sandbox/http/swagger_generated/restapi/operations"
	"sandbox/http/swagger_generated/restapi/operations/files"
	"sandbox/http/swagger_generated/restapi/operations/products"
)

//go:generate swagger generate server --target ../../swagger_generated --name EchoExample --spec ../api/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.EchoExampleAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EchoExampleAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.BinProducer = runtime.ByteStreamProducer()
	api.JSONProducer = runtime.JSONProducer()

	if api.ProductsCreateProductHandler == nil {
		api.ProductsCreateProductHandler = products.CreateProductHandlerFunc(func(params products.CreateProductParams) middleware.Responder {
			return middleware.NotImplemented("operation products.CreateProduct has not yet been implemented")
		})
	}
	if api.ProductsDeleteProductHandler == nil {
		api.ProductsDeleteProductHandler = products.DeleteProductHandlerFunc(func(params products.DeleteProductParams) middleware.Responder {
			return middleware.NotImplemented("operation products.DeleteProduct has not yet been implemented")
		})
	}
	if api.FilesDownloadFileHandler == nil {
		api.FilesDownloadFileHandler = files.DownloadFileHandlerFunc(func(params files.DownloadFileParams) middleware.Responder {
			return middleware.NotImplemented("operation files.DownloadFile has not yet been implemented")
		})
	}
	if api.ProductsGetProductHandler == nil {
		api.ProductsGetProductHandler = products.GetProductHandlerFunc(func(params products.GetProductParams) middleware.Responder {
			return middleware.NotImplemented("operation products.GetProduct has not yet been implemented")
		})
	}
	if api.ProductsListProductsHandler == nil {
		api.ProductsListProductsHandler = products.ListProductsHandlerFunc(func(params products.ListProductsParams) middleware.Responder {
			return middleware.NotImplemented("operation products.ListProducts has not yet been implemented")
		})
	}
	if api.ProductsUpdateProductHandler == nil {
		api.ProductsUpdateProductHandler = products.UpdateProductHandlerFunc(func(params products.UpdateProductParams) middleware.Responder {
			return middleware.NotImplemented("operation products.UpdateProduct has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
