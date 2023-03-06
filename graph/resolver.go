package graph

import "github.com/kendo_ow/gographql/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct{

	videos []*model.Video
}
