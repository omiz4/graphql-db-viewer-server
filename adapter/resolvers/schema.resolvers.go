package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"
	"gqlgen-db-viewer/application/model"
	"gqlgen-db-viewer/graphql/generated"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context) (model.Node, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
