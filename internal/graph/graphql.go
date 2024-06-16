package graph

import (
	"github.com/graphql-go/graphql"
	"github.com/msaufi2325/go-movies-back-end/internal/models"
)

type Graph struct {
	Movies      []*models.Movie
	QueryString string
	Config      graphql.SchemaConfig
	fields      graphql.Fields
	movieType   *graphql.Object
}
