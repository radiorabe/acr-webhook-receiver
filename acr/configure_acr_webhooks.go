// This file is safe to edit. Once it exists it will not be overwritten

package acr

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/radiorabe/acr-webhook-receiver/acr/operations"
	apiop "github.com/radiorabe/acr-webhook-receiver/acr/operations/api"
	"github.com/radiorabe/acr-webhook-receiver/acr/operations/webhook"
	"github.com/radiorabe/acr-webhook-receiver/models"
)

//go:generate swagger generate server --spec ../swagger.yml --target ../ --additional-initialism=ACR --principal models.Principal --default-scheme=https --name ACRWebhooks --server-package acr
//go:generate swagger generate client --spec ../swagger.yml --target ../ --additional-initialism=ACR --principal models.Principal --default-scheme=https --tags=api

var dbConn *gorm.DB

var dbFlags = struct {
	DSN string `env:"POSTGRES_DSN" long:"postgres-dsn" description:"PostgreSQL DSN for gorm.io/driver/postgresql"`
}{}

var webhookFlags = struct {
	APIKey string `env:"ACR_API_KEY" long:"api-key" description:"API key transmitted by ACRCloud with each request"`
}{}

func configureFlags(api *operations.ACRWebhooksAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Webhook Settings",
			LongDescription:  "",
			Options:          &webhookFlags,
		},
		{
			ShortDescription: "Database Settings",
			LongDescription:  "",
			Options:          &dbFlags,
		},
	}
}

func configureAPI(api *operations.ACRWebhooksAPI) http.Handler {
	api.ServeError = errors.ServeError

	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.APIKeyAuth = func(token string) (*models.Principal, error) {
		if token == webhookFlags.APIKey {
			prin := models.Principal(token)
			return &prin, nil
		}
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.WebhookAddResultHandler = webhook.AddResultHandlerFunc(func(params webhook.AddResultParams, principal *models.Principal) middleware.Responder {
		fmt.Print(principal)
		record := &models.Result{
			Result: params.Body,
		}
		result := getDatabase().Create(record)
		if result.Error != nil {
			return webhook.NewAddResultInternalServerError()
		}
		if result.RowsAffected != 1 {
			return webhook.NewAddResultInternalServerError()
		}
		return webhook.NewAddResultCreated()
	})

	api.APIGetResultsHandler = apiop.GetResultsHandlerFunc(func(params apiop.GetResultsParams) middleware.Responder {
		var result []*models.Result

		tsUTCQuery := "to_timestamp((result -> 'data' -> 'metadata' ->> 'timestamp_utc') || ' 0000', 'YYYY-MM-DD HH24:MI:SS TZH')"

		query := getDatabase().Model(&models.Result{}).Limit(int(*params.Limit)).Offset(int(*params.Offset))

		if params.From != nil {
			from, err := params.From.Value()
			if err != nil {
				return apiop.NewGetResultsInternalServerError()
			}
			query = query.Where(
				fmt.Sprintf("%s => ?", tsUTCQuery),
				from,
			)
		}
		if params.To != nil {
			to, err := params.To.Value()
			if err != nil {
				return apiop.NewGetResultsInternalServerError()
			}
			query = query.Where(
				fmt.Sprintf("%s <= ?", tsUTCQuery),
				to,
			)
		}

		query = query.Order(fmt.Sprintf("%s desc", tsUTCQuery))

		tx := query.Scan(&result)
		if tx.Error != nil {
			return apiop.NewGetResultsInternalServerError()
		}
		return apiop.NewGetResultsOK().WithPayload(result)
	})

	api.APIGetResultHandler = apiop.GetResultHandlerFunc(func(params apiop.GetResultParams) middleware.Responder {
		result := &models.Result{}
		tx := getDatabase().First(result, params.ResultID)
		if tx.Error != nil {
			return apiop.NewGetResultInternalServerError()
		}

		return apiop.NewGetResultOK().WithPayload(result)
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
	getDatabase().AutoMigrate(&models.Result{})
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	swaggerOpts := &middleware.SwaggerUIOpts{}
	return middleware.SwaggerUI(*swaggerOpts, handler)
}

func getDatabase() *gorm.DB {
	if dbConn == nil {
		conn, err := gorm.Open(postgres.Open(dbFlags.DSN), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		dbConn = conn
	}
	return dbConn
}
