package appy

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

// The HttpProvider interface is used to define the methods that are required for a http server, these are implemented by drivers
type HttpServer interface {
	Initialize(*Appy, HttpOptions) error
	Run() error

	// Get the root group for the http server i.e. /
	RootGroup() HttpEndpointGroup
}

// Utility struct for mapping errors to http responses
type HttpErrorMapper interface {
	Map(error) HttpResult
}

// Options when creating a new http server
type HttpOptions struct {
	Provider HttpServer

	// Error mapper
	ErrorMapper HttpErrorMapper

	// The address to bind the server to
	Address string

	// SSL settings for HTTPS, runs on HTTP if nil
	SSL *SSLSettings
}

// SSLSettings is used to define the settings for a https server
type SSLSettings struct {
	CertFile string
	KeyFile  string
}

// HttpEndpointGroup is used to group http methods together
type HttpEndpointGroup interface {
	// Create a subgroup of the current group i.e. /api -> /api/v1
	Subgroup(string) HttpEndpointGroup

	// Attach pre-handle middleware
	Pre(HttpMiddleware)

	// Attach post-handle middleware
	Post(HttpMiddleware)

	StaticFile(string, string)
	StaticDir(string, http.FileSystem)

	GET(string, HttpHandler)
	POST(string, HttpHandler)
	PATCH(string, HttpHandler)
	PUT(string, HttpHandler)
	DELETE(string, HttpHandler)
}

// HttpHandler is a function that handles an http request
type HttpHandler func(c HttpContext) HttpResult

// A http handler context
type HttpContext struct {
	App    *Appy
	Header HeaderParser
	Query  QueryParameterParser
	Path   PathParameterParser
	Body   BodyParser

	Writer  http.ResponseWriter
	Request *http.Request

	// Temporary storage to pass from middleware to handler
	tempStorage map[string]any
}

// A result of a http request
type HttpResult struct {
	StatusCode int
	Body       interface{}
	Error      error

	failed bool
}

// HeaderParser is used to parse headers from a http request
type HeaderParser interface {
	// Get a header by name
	ExpectSingleString(string) (string, error)
}

// QueryParameterParser is used to parse query parameters from a http request
type QueryParameterParser interface {
	// A utility function to get a page from a query parameter, same as calling GetInt("page")
	Page() int

	// Get a query paramter by name
	GetString(string) string

	// Get a query paramter by name and convert to int
	GetInt(string) int

	ExpectPage() (int, error)
	ExpectString(string) (string, error)
	ExpectInt(string) (int, error)
}

// PathParameterParser is used to parse path parameters from a http request
type PathParameterParser interface {
	// Get a path paramter by name and convert to int
	GetInt(string) int

	ExpectInt(string) (int, error)
}

// BodyParser is used to parse the body of a http request
type BodyParser interface {
	// Parse the body into a struct
	ParseSingle(any) error

	// Parse an array of structs from the body
	ParseArray(any) error
}

// This is a middleware function that can be used to add functionality that runs before the main handler,
// if the returned result is not Nil it will be passed down the chain
type HttpMiddleware func(c HttpContext) HttpResult

func (c *HttpContext) Set(key string, value any) {
	if c.tempStorage == nil {
		c.tempStorage = make(map[string]any)
	}

	c.tempStorage[key] = value
}

func (c *HttpContext) Get(key string) (any, HttpResult) {
	if c.tempStorage == nil {
		return nil, c.Error(errors.New("no values in temporary storage"))
	}

	value, ok := c.tempStorage[key]
	if !ok {
		return nil, c.Error(errors.New("key not found in temporary storage"))
	}

	return value, c.Nil()
}

func (c *HttpContext) StoreMultipartFile(key string, outDir string) (string, HttpResult) {
	file, header, err := c.Request.FormFile(key)
	if err != nil {
		return "", c.Error(err)
	}

	// Create a name uuid
	extension := strings.Split(header.Filename, ".")[1]
	filename := uuid.New().String() + "." + extension

	// Store locally
	out, err := os.Create(outDir + filename)
	if err != nil {
		return "", c.Error(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", c.Error(err)
	}

	// File route
	return fmt.Sprintf("/images/%v", filename), c.Nil()
}

func (c *HttpContext) Nil() HttpResult {
	return HttpResult{
		failed: false,
	}
}

func (c *HttpContext) Ok(statusCode int, body interface{}) HttpResult {
	return HttpResult{
		StatusCode: statusCode,
		Body:       body,
		failed:     false,
	}
}

func (c *HttpContext) NotFound() HttpResult {
	return HttpResult{
		StatusCode: http.StatusNotFound,
		failed:     true,
	}
}

func (c *HttpContext) BadRequest(body interface{}) HttpResult {
	return HttpResult{
		StatusCode: http.StatusBadRequest,
		Body:       body,
		failed:     true,
	}
}

func (c *HttpContext) Fail(statusCode int, body interface{}) HttpResult {
	return HttpResult{
		StatusCode: statusCode,
		Body:       body,
		failed:     true,
	}
}

func (c *HttpContext) Error(err error) HttpResult {
	return HttpResult{
		StatusCode: http.StatusInternalServerError,
		Error:      err,
		failed:     true,
	}
}

func (hr HttpResult) IsFailed() bool {
	return hr.failed || hr.Error != nil
}

func (hr HttpResult) HasError() bool {
	return hr.Error != nil
}
