package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/oapi-codegen/runtime"
	"github.com/sawitpro/technical_test/config"
	"github.com/sawitpro/technical_test/handler"
	"github.com/sawitpro/technical_test/repository"
	"github.com/sawitpro/technical_test/usecase"
)

func main() {
	InitServer()
}

func InitServer() {
	echoServer := echo.New()

	cfg := config.NewConfig()

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Panic(err)
	}

	repo := repository.NewRepository(cfg)
	uc := usecase.NewUserUsecase(cfg, repo)
	hand := handler.NewHandler(uc)

	RegisterHandlers(echoServer, hand, cfg)

	if err := echoServer.Start(fmt.Sprintf(":%d", serverPort)); err != nil {
		log.Println(err)
	}
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler handler.ServerInterface
}

// Login converts echo context to params.
func (w *ServerInterfaceWrapper) Login(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Login(ctx)
	return err
}

// GetUserProfile converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetUserProfile(ctx, id)
	return err
}

// UpdateUserProfile converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateUserProfile(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.UpdateUserProfile(ctx, id)
	return err
}

// Registration converts echo context to params.
func (w *ServerInterfaceWrapper) Registration(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.Registration(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si handler.ServerInterface, cfg *config.Config) {
	RegisterHandlersWithBaseURL(router, si, cfg, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si handler.ServerInterface, cfg *config.Config, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/login", wrapper.Login)
	router.GET(baseURL+"/profile/:id", wrapper.GetUserProfile, config.JWTVerify(cfg.PublicKey))
	router.PUT(baseURL+"/profile/:id", wrapper.UpdateUserProfile, config.JWTVerify(cfg.PublicKey))
	router.POST(baseURL+"/registration", wrapper.Registration)

}
