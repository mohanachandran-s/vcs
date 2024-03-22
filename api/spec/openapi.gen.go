// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	externalRef0 "github.com/trustbloc/vcs/pkg/restapi/v1/common"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "./common.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
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
