// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
	"net/http"
	"strings"
)

// Error defines model for Error.
type Error struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

// NewPet defines model for NewPet.
type NewPet struct {
	Name string  `json:"name"`
	Tag  *string `json:"tag,omitempty"`
}

// Pet defines model for Pet.
type Pet struct {
	// Embedded struct due to allOf(#/components/schemas/NewPet)
	NewPet
	// Embedded fields due to inline allOf schema
	Id int64 `json:"id"`
}

// FindPetsParams defines parameters for FindPets.
type FindPetsParams struct {
	Tags  *[]string `json:"tags,omitempty"`
	Limit *int32    `json:"limit,omitempty"`
}

// addPetJSONBody defines parameters for AddPet.
type addPetJSONBody NewPet

// DeletePetParams defines parameters for DeletePet.
type DeletePetParams struct {
	Id int64 `json:"id"`
}

// FindPetByIdParams defines parameters for FindPetById.
type FindPetByIdParams struct {
	Id int64 `json:"id"`
}

// AddPetRequestBody defines body for AddPet for application/json ContentType.
type AddPetJSONRequestBody addPetJSONBody

type ServerInterface interface {
	//  (GET /pets)
	FindPets(w http.ResponseWriter, r *http.Request)
	//  (POST /pets)
	AddPet(w http.ResponseWriter, r *http.Request)
	//  (DELETE /pets/{id})
	DeletePet(w http.ResponseWriter, r *http.Request)
	//  (GET /pets/{id})
	FindPetById(w http.ResponseWriter, r *http.Request)
}

// ParamsForFindPets operation parameters from context
func ParamsForFindPets(ctx context.Context) *FindPetsParams {
	return ctx.Value("FindPetsParams").(*FindPetsParams)
}

// FindPets operation middleware
func FindPetsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		var err error

		// ------------- Optional query parameter "tags" -------------
		if paramValue := r.URL.Query().Get("tags"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "tags", r.URL.Query(), &params.Tags)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter tags: %s", err), http.StatusBadRequest)
			return
		}

		// ------------- Optional query parameter "limit" -------------
		if paramValue := r.URL.Query().Get("limit"); paramValue != "" {

		}

		err = runtime.BindQueryParameter("form", true, false, "limit", r.URL.Query(), &params.Limit)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter limit: %s", err), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(r.Context(), "FindPetsParams", &params)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AddPet operation middleware
func AddPetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// DeletePet operation middleware
func DeletePetCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// ------------- Path parameter "id" -------------
		var pathErr error
		var id int64

		pathErr = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
		if pathErr != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", pathErr), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(r.Context(), "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// FindPetById operation middleware
func FindPetByIdCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// ------------- Path parameter "id" -------------
		var pathErr error
		var id int64

		pathErr = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
		if pathErr != nil {
			http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", pathErr), http.StatusBadRequest)
			return
		}

		ctx = context.WithValue(r.Context(), "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(FindPetsCtx)
		r.Get("/pets", si.FindPets)
	})
	r.Group(func(r chi.Router) {
		r.Use(AddPetCtx)
		r.Post("/pets", si.AddPet)
	})
	r.Group(func(r chi.Router) {
		r.Use(DeletePetCtx)
		r.Delete("/pets/{id}", si.DeletePet)
	})
	r.Group(func(r chi.Router) {
		r.Use(FindPetByIdCtx)
		r.Get("/pets/{id}", si.FindPetById)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
