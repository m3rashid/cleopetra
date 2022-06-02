package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/m3rashid/cleopetra/server/graph/generated"
	"github.com/m3rashid/cleopetra/server/graph/model"
)

func (r *mutationResolver) Login(
	ctx context.Context,
	input *model.LoginInput) (*model.Auth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Signup(
	ctx context.Context,
	input *model.SignupInput) (*model.Auth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) VerifyCreateAuthor(
	ctx context.Context,
	input *model.VerifyCreateAuthorInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(
	ctx context.Context,
	input *model.CreateAuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "Server is running OK", nil
}

func (r *queryResolver) GetUser(ctx context.Context) (*model.Auth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAuthors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAuthor(ctx context.Context, id string) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPosts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPost(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategories(ctx context.Context) ([]*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCategory(ctx context.Context, id string) (*model.Category, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
