package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mark-by/golang/try_graphql/graph/api/db"
	"github.com/mark-by/golang/try_graphql/graph/api/dberrors"
	"github.com/mark-by/golang/try_graphql/graph/generated"
	"github.com/mark-by/golang/try_graphql/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	newVideo := &model.Video{
		URL:       input.URL,
		Name:      input.Name,
		Description: input.Description,
		UserID: input.UserID,
		CreatedAt: time.Now().UTC(),
	}

	rows, err := db.LogAndQuery(r.Db, "INSERT INTO videos (name, url, user_id, created_at) VALUES($1, $2, $3, $4) RETURNING id",
		input.Name, input.URL, input.UserID, newVideo.CreatedAt)
	defer rows.Close()

	if err != nil || !rows.Next() {
		return &model.Video{}, err
	}
	if err := rows.Scan(&newVideo.ID); err != nil {
		dberrors.DebugPrintf(err)
		if dberrors.IsForeignKeyError(err) {
			return &model.Video{}, dberrors.UserNotExist
		}
		return &model.Video{}, dberrors.InternalServerError
	}

	return newVideo, nil
}

func (r *queryResolver) Videos(ctx context.Context, limit *int, offset *int) ([]*model.Video, error) {
	video := &model.Video{}
	var videos []*model.Video

	rows, err := db.LogAndQuery(r.Db, "SELECT id, name, url, created_at, user_id FROM videos ORDER BY created_at desc limit $1 offset $2", limit, offset)
	defer rows.Close()

	if err != nil {
		dberrors.DebugPrintf(err)
		return nil, dberrors.InternalServerError
	}
	for rows.Next() {
		if err := rows.Scan(&video.ID, &video.Name, &video.URL, &video.CreatedAt, &video.UserID); err != nil {
			dberrors.DebugPrintf(err)
			return nil, dberrors.InternalServerError
		}
		videos = append(videos, video)
	}

	return videos, nil
}

func (r *videoResolver) User(ctx context.Context, obj *model.Video) (*model.User, error) {
	rows, _ := db.LogAndQuery(r.Db, "SELECT id, name, email FROM users where id = $1", obj.UserID)
	defer rows.Close()

	if !rows.Next() {
		return &model.User{}, nil
	}
	user := &model.User{}
	if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		dberrors.DebugPrintf(err)
		return &model.User{}, dberrors.InternalServerError
	}

	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Video returns generated.VideoResolver implementation.
func (r *Resolver) Video() generated.VideoResolver { return &videoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type videoResolver struct{ *Resolver }
