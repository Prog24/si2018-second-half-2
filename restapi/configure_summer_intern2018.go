// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"

	"github.com/eure/si2018-second-half-2/controllers/message"
	"github.com/eure/si2018-second-half-2/controllers/tempmessage"
	"github.com/eure/si2018-second-half-2/controllers/token"
	"github.com/eure/si2018-second-half-2/controllers/user"
	"github.com/eure/si2018-second-half-2/controllers/userimage"
	"github.com/eure/si2018-second-half-2/controllers/userlike"
	"github.com/eure/si2018-second-half-2/controllers/usermatch"
	"github.com/eure/si2018-second-half-2/controllers/usertempmatch"
	"github.com/eure/si2018-second-half-2/restapi/summerintern"
)

//go:generate swagger generate server --target .. --name  --spec ../si2018.yml --api-package summerintern

func configureFlags(api *summerintern.SummerIntern2018API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *summerintern.SummerIntern2018API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Implemented
	api.GetProfileByUserIDHandler = summerintern.GetProfileByUserIDHandlerFunc(user.GetProfileByUserID)
	api.GetTokenByUserIDHandler = summerintern.GetTokenByUserIDHandlerFunc(token.GetTokenByUserID)
	api.PutProfileHandler = summerintern.PutProfileHandlerFunc(user.PutProfile)
	api.GetUsersHandler = summerintern.GetUsersHandlerFunc(user.GetUsers)
	api.PostMessageHandler = summerintern.PostMessageHandlerFunc(message.PostMessage)
	api.GetMessagesHandler = summerintern.GetMessagesHandlerFunc(message.GetMessages)
	api.GetLikesHandler = summerintern.GetLikesHandlerFunc(userlike.GetLikes)
	api.PostLikeHandler = summerintern.PostLikeHandlerFunc(userlike.PostLike)
	api.GetMatchesHandler = summerintern.GetMatchesHandlerFunc(usermatch.GetMatches)
	api.PostImagesHandler = summerintern.PostImagesHandlerFunc(userimage.PostImage)
	api.GetTempMatchHandler = summerintern.GetTempMatchHandlerFunc(usertempmatch.GetTempMatch)
	api.PostTempMatchHandler = summerintern.PostTempMatchHandlerFunc(usertempmatch.PostTempMatch)
	api.PutTempMatchHandler = summerintern.PutTempMatchHandlerFunc(usertempmatch.PutTempMatch)
	api.GetTempMessagesHandler = summerintern.GetTempMessagesHandlerFunc(tempmessage.GetTempMessages)
	api.PostTempMessageHandler = summerintern.PostTempMessageHandlerFunc(tempmessage.PostTempMessage)

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
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
