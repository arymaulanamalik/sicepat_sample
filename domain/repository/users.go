package repository

import (
	"context"
	"time"

	"github.com/maul/sicepat_sample/domain/model"
	"github.com/maul/sicepat_sample/shared"
	uuid "github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepositoryImpl struct {
	mongoDb *mongo.Database
}

type UsersRepository interface {
	InsertUser(context.Context, model.User) (string, error)
}

func NewUsersRepository(mdb *mongo.Database) UsersRepository {
	return &UsersRepositoryImpl{mongoDb: mdb}
}

func (br *UsersRepositoryImpl) InsertUser(ctx context.Context, User model.User) (id string, err error) {
	UseresColl := br.mongoDb.Collection(shared.CollectionUsers)

	timeNow := time.Now()

	User.ID = uuid.NewV4().String()
	User.IsDeleted = false
	User.CreatedAt = timeNow
	User.UpdatedAt = timeNow
	User.IsActive = true

	_id, err := UseresColl.InsertOne(ctx, User)
	if err != nil {
		return "", err
	}

	return _id.InsertedID.(string), nil
}
