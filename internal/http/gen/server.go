// Package gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package gen

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /sushi)
	FindSushis(ctx echo.Context, params FindSushisParams) error

	// (POST /sushi)
	AddSushi(ctx echo.Context) error

	// (GET /sushi/{id})
	FindSushiById(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindSushis converts echo context to params.
func (w *ServerInterfaceWrapper) FindSushis(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindSushisParams
	// ------------- Optional query parameter "asc" -------------

	err = runtime.BindQueryParameter("form", true, false, "asc", ctx.QueryParams(), &params.Asc)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asc: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindSushis(ctx, params)
	return err
}

// AddSushi converts echo context to params.
func (w *ServerInterfaceWrapper) AddSushi(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddSushi(ctx)
	return err
}

// FindSushiById converts echo context to params.
func (w *ServerInterfaceWrapper) FindSushiById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindSushiById(ctx, id)
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
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/sushi", wrapper.FindSushis)
	router.POST(baseURL+"/sushi", wrapper.AddSushi)
	router.GET(baseURL+"/sushi/:id", wrapper.FindSushiById)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xUQW8TOxD+K9W8d1w1ee0Th721Eki9ABLHKgd3PUld7dqu7QAh2sOmqEEUqahSW0BI",
	"9IAqIaSqJy5Q+mPcROVfIHuzSZpsGxAcOJBLHHvyzTfffDNtiEQiBUduNIRt0NE6JsQfbysllDtIJSQq",
	"w9BfR4Ki+64LlRADITBuFhcgANOSmP/EBipIA0hQa9Lw0YNHbRTjDUjTABRuNplCCuFqjjmKrw3BxNoG",
	"RsZh3cVHD5p6nU0T4iQpSxGAVCwafxljpsUT4qGYwUSX/ntwQZQirSnCPmeRoYAroz3kTOL4Xh3C1Tb8",
	"q7AOIfxTGSlfGcheGVaZBpNlMjqp+q3/S1SfIMpoCa1a6sIYrwsHSVFHiknDBIcwJzy3dH/FYTMTe1Hc",
	"3ZJkEMBDVDoP/G++Ol91JQqJnEgGISz6qwAkMeuec0UX5TfQTOfyr7azd/m+298/7e0e9L4eggdUxIWs",
	"UAjhDuPUk9IeWZEEDSrtpbwK13/V/Xa0bbOPNnttOzs227HZB5tt22wHXLkQwmYTVQuCgWeA6Mh1z4s/",
	"ZoI1IWIk3Ik5mSMnmSe4+Pypv396DXTMEmaugM8cmLTmeqel4Dpv+UK1mk8cN8i9fkTKmEVenMqGdoza",
	"YxmGZr7JYoW/Jg2eBmXdmSsIgX+vk2ZsforTTVTyDVOSusnxscTIIJ3DUYwU+loTZSf9g9PL492Ls7f9",
	"Zy+nTLREcw9BPh2ozbKgrd9WyGhsp2u5PD/rPT/KHaPHOOQTalQT019s+w90+4/vbhoMlkWlzWg6c2Mw",
	"ajt7/Rfd3skbmx3a7HjW8lhurdBZ+6OY7UObvbPZU7vVtVtfbOfcZieMFmPudttoyv391W5eN/Ll27r2",
	"t/fu8z0AAP//gyvjcocIAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}

