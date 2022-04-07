package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

	"github.com/DavidHODs/hackernews/graph/generated"
	"github.com/DavidHODs/hackernews/graph/model"
	"github.com/DavidHODs/hackernews/internal/links"
	"github.com/DavidHODs/hackernews/internal/users"
	"github.com/DavidHODs/hackernews/pkg/jwt"
)

// func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
// 	var link model.Link
// 	var user model.User
// 	link.Address = input.Address
// 	link.Title = input.Title
// 	user.Name = "test user"
// 	link.User = &user

// 	return &link, nil
// }

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
	user.Username = input.Username
	user.Password = input.Password
	user.Create()
	token, err := jwt.GenerateToken(user.Username)
	if err != nil{
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
// 	var links []*model.Link
// 	dummyLink := model.Link {
// 		Title: "first dummy link",
// 		Address: "https://dummy.org",
// 		User: &model.User{Name: "owner"},
// 	}

// 	dummyLink2 := model.Link {
// 		Title: "second dummy link",
// 		Address: "https://dummy1.org",
// 		User: &model.User{Name: "admin"},
// 	}

// 	dummyLink3 := model.Link {
// 		Title: "third dummy link",
// 		Address: "https://dummy2.org",
// 		User: &model.User{Name: "user"},
// 	}

// 	links = append(links, &dummyLink, &dummyLink2, &dummyLink3)
// 	return links, nil
// }

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	var link links.Link
	link.Title = input.Title
	link.Address = input.Address

	linkID := link.Save()
	return &model.Link {
		ID: strconv.FormatInt(linkID, 10),
		Title: link.Title,
		Address: link.Address,
	}, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	var resultLinks []*model.Link
	var dbLinks []links.Link
	dbLinks = links.GetAll()
	for _, link := range dbLinks {
		resultLinks = append(resultLinks, &model.Link{ID: link.ID, Title: link.Title, Address: link.Address})
	}

	return resultLinks, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
