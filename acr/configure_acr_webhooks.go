// This file is safe to edit. Once it exists it will not be overwritten

package acr

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/radiorabe/acr-webhook-receiver/acr/operations"
	apiop "github.com/radiorabe/acr-webhook-receiver/acr/operations/api"
	"github.com/radiorabe/acr-webhook-receiver/acr/operations/compat"
	"github.com/radiorabe/acr-webhook-receiver/acr/operations/webhook"
	"github.com/radiorabe/acr-webhook-receiver/models"
)

//go:generate swagger generate server --spec ../swagger.yml --target ../ --additional-initialism=ACR --principal models.Principal --default-scheme=https --name ACRWebhooks --server-package acr
//go:generate swagger generate client --spec ../swagger.yml --target ../ --additional-initialism=ACR --principal models.Principal --default-scheme=https --tags=api --tags=compat

type requestIDKeyType string

const (
	resultsRecordTsUTCQuery                  = "to_timestamp((result -> 'data' -> 'metadata' ->> 'timestamp_utc') || ' 0000', 'YYYY-MM-DD HH24:MI:SS TZH')"
	requestIDKey            requestIDKeyType = ""
)

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

		query := getDatabase().Model(
			&models.Result{},
		)
		if params.From != nil {
			from, err := params.From.Value()
			if err != nil {
				return apiop.NewGetResultsInternalServerError()
			}
			query = query.Where(
				fmt.Sprintf("%s >= ?", resultsRecordTsUTCQuery),
				from,
			)
		}
		if params.To != nil {
			to, err := params.To.Value()
			if err != nil {
				return apiop.NewGetResultsInternalServerError()
			}
			query = query.Where(
				fmt.Sprintf("%s <= ?", resultsRecordTsUTCQuery),
				to,
			)
		}

		tx := query.Limit(
			int(*params.Limit),
		).Offset(
			int(*params.Offset),
		).Order(
			fmt.Sprintf("%s desc", resultsRecordTsUTCQuery),
		).Scan(
			&result,
		)

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

	api.CompatGetCustomStreamHandler = compat.GetCustomStreamHandlerFunc(func(params compat.GetCustomStreamParams) middleware.Responder {
		var result []*models.Data

		date, err := time.Parse("20060102", params.Date)
		if err != nil {
			return compat.NewGetCustomStreamBadRequest()
		}

		tx := getDatabase().Model(
			&models.Result{},
		).Select([]string{
			"result -> 'data' -> 'metadata' AS metadata",
			"result -> 'result_type' AS result_type",
			"result -> 'data' -> 'status' AS status",
		}).Where(
			fmt.Sprintf("%s >= ?", resultsRecordTsUTCQuery),
			date.Format("2006-01-02T15:04:05Z"),
		).Where(
			fmt.Sprintf("%s < ?", resultsRecordTsUTCQuery),
			date.Add(24*time.Hour).Format("2006-01-02T15:04:05Z"),
		).Order(
			fmt.Sprintf("%s asc", resultsRecordTsUTCQuery),
		).Scan(
			&result,
		)

		if tx.Error != nil {
			return compat.NewGetCustomStreamInternalServerError()
		}
		return compat.NewGetCustomStreamOK().WithPayload(result)
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
	if err := getDatabase().AutoMigrate(&models.Result{}); err != nil {
		log.WithError(err).Fatal(err)
	}
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {

	nextRequestID := func() string {
		return fmt.Sprintf("%v", uuid.Must(uuid.NewV4()))
	}

	logger := log.New()
	logWriter := logger.Writer()
	defer logWriter.Close()

	return tracing(nextRequestID)(logging(logger)(handler))
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

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
