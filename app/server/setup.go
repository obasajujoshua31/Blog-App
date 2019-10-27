package server

import (
	"blog-app/app/config"
	"blog-app/app/services"
	"encoding/json"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/graphql-go/graphql"
	"net/http"
)

type Server struct {
	Configuration *config.AppConfig
	GQLSchema     *graphql.Schema
	Router        *Router
}

type reqBody struct {
	Query string `json:"query"`
}

func NewServer(config *config.AppConfig, schema *graphql.Schema) *Server {
	server := &Server{
		Configuration: config,
		Router:        NewRouter(),
		GQLSchema:     schema,
	}

	return server
}

func (s *Server) startGraphqlServer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Must Provide graphql query in request body", 400)
			return
		}

		var rBody reqBody

		err := json.NewDecoder(r.Body).Decode(&rBody)

		if err != nil {
			http.Error(w, "Error parsing JSON request body", 400)
		}

		results := services.ExecuteQuery(rBody.Query, *s.GQLSchema)

		render.JSON(w, r, results)
	}
}

func (s *Server) InitMiddlewares() {
	s.Router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.StripSlashes,
		middleware.Recoverer,
	)
}

func connectToDB(appConfig config.AppConfig) (*services.DB, error) {

	db, err := services.New(
		services.ConnString(appConfig))

	if err != nil {
		return nil, err
	}

	return db, nil
}
