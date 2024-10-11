package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Anuolu-2020/hackernews-api-clone/graph/model"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/auth"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/links"
	"github.com/Anuolu-2020/hackernews-api-clone/internal/users"
	"github.com/Anuolu-2020/hackernews-api-clone/pkg/token"
)

// CreateLink is the resolver for the createLink field.
func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	var link links.Link
	link.Title = input.Title
	link.Address = input.Address

	link.User = user

	linkID := link.Save()
	graphqlUser := &model.User{ID: user.ID, Name: user.Username}

	return &model.Link{ID: strconv.FormatInt(linkID, 10), Title: link.Title, Address: link.Address, User: graphqlUser}, nil
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	tokenStr, err := token.GenerateToken(user.Username)
	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	return tokenStr, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password

	correct := user.Authenticate()
	if !correct {
		return "", &users.WrongUsernameOrPasswordErr{}
	}
	tokenStr, err := token.GenerateToken(user.Username)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := token.ParseToken(input.Token)
	if err != nil {
		return "", fmt.Errorf("access denied")
	}
	tokenStr, err := token.GenerateToken(username)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

// Links is the resolver for the links field.
func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link

	dbLinks := links.GetAll()

	for _, link := range dbLinks {
		graphqlUser := &model.User{
			ID:   link.User.ID,
			Name: link.User.Username,
		}
		resultLinks = append(resultLinks, &model.Link{Title: link.Title, Address: link.Address, ID: strconv.FormatUint(uint64(link.ID), 10), User: graphqlUser})
	}

	return resultLinks, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type (
	mutationResolver struct{ *Resolver }
	queryResolver    struct{ *Resolver }
)

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
*/
