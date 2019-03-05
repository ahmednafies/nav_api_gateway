package roothandler

import (
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/graphql-go-handler"
	"github.com/hem-nav-gateway/config"
	"github.com/hem-nav-gateway/errorhandler"
	"github.com/hem-nav-gateway/fields"
	"github.com/hem-nav-gateway/session"
	"log"
	"net/http"
)

func pathVariables(vars map[string]string) (string, error) {
	companyPath := vars["company"]
	if vars["company"] == "test" {
		return config.TestCompanyName, nil
	}

	return "", errorhandler.CompanyDoesNotExist(companyPath)
}

func Handler() *gqlhandler.Handler {

	query := fields.QueryType()
	mutation := fields.MutationType()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	handler := gqlhandler.New(&gqlhandler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: false,
	})

	return handler
}

func RootEndpoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName, err := pathVariables(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	session.SetSession(r, companyName)
	handler := Handler()
	handler.ServeHTTP(w, r)
}
