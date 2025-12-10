package graph

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Infamous003/graphql-service/graph/model"
)

func (r *mutationResolver) Follow(ctx context.Context, followerID string, followeeID string) (*model.FollowResponse, error) {
	// converting ID to int
	followerIDInt, err := strconv.Atoi(followerID)
	if err != nil {
		msg := "invalid follower ID"
		return &model.FollowResponse{Message: &msg}, nil
	}

	followeeIDInt, err := strconv.Atoi(followeeID)
	if err != nil {
		msg := "invalid followee ID"
		return &model.FollowResponse{Message: &msg}, nil
	}

	// creating the payload
	payload := map[string]int{
		"follower_id": followerIDInt,
		"followee_id": followeeIDInt,
	}

	var resp model.FollowResponse
	err = r.Resolver.callFollowService(http.MethodPost, "/follow", payload, &resp)
	if err != nil {
		msg := err.Error()
		return &model.FollowResponse{Message: &msg}, nil
	}

	return &resp, nil
}

func (r *mutationResolver) Unfollow(ctx context.Context, followerID string, followeeID string) (*model.FollowResponse, error) {
	followerIDInt, err := strconv.Atoi(followerID)
	if err != nil {
		msg := "invalid follower ID"
		return &model.FollowResponse{Message: &msg}, nil
	}

	followeeIDInt, err := strconv.Atoi(followeeID)
	if err != nil {
		msg := "invalid followee ID"
		return &model.FollowResponse{Message: &msg}, nil
	}

	payload := map[string]int{
		"follower_id": followerIDInt,
		"followee_id": followeeIDInt,
	}

	var resp model.FollowResponse
	err = r.Resolver.callFollowService(http.MethodPost, "/unfollow", payload, &resp)
	if err != nil {
		msg := err.Error()
		return &model.FollowResponse{Message: &msg}, nil
	}

	return &resp, nil
}

func (r *queryResolver) Followers(ctx context.Context, userID string) ([]*model.User, error) {
	var resp struct {
		Followers []*model.User `json:"followers"`
	}

	endpoint := fmt.Sprintf("/users/%s/followers", userID)
	err := r.Resolver.callFollowService(http.MethodGet, endpoint, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch followers: %v", err)
	}

	return convertUsers(resp.Followers), nil
}

func (r *queryResolver) Following(ctx context.Context, userID string) ([]*model.User, error) {
	var resp struct {
		Following []*model.User `json:"following"`
	}

	endpoint := fmt.Sprintf("/users/%s/following", userID)
	err := r.Resolver.callFollowService(http.MethodGet, endpoint, nil, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch following: %v", err)
	}

	return convertUsers(resp.Following), nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
