package user

import (
	"context"
	"errors"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/consts"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	prefixPhotoCacheKey = "cache:photo:"
	CollectionName      = "user"
)

type (
	// IMongoMapper is an interface to be customized, add more methods here,
	// and implement the added methods in MongoMapper.
	IMongoMapper interface {
		Insert(ctx context.Context, data *Photo) error
		FindOne(ctx context.Context, id string) (*Photo, error)
		Upsert(ctx context.Context, data *Photo) error
		List(ctx context.Context, albumId string, skip int64, count int64, onlyFeatured bool) ([]*Photo, int64, error)
		Delete(ctx context.Context, id string) error
	}

	MongoMapper struct {
		conn *monc.Model
	}

	Photo struct {
		ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
		CreatedAt         time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
		UpdatedAt         time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
		DeletedAt         time.Time          `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
		AlbumId           string             `bson:"albumId,omitempty" json:"albumId,omitempty"`
		OwnerId           string             `bson:"ownerId,omitempty" json:"ownerId,omitempty"`
		IsFeatured        bool               `bson:"isFeatured,omitempty" json:"isFeatured,omitempty"`
		Location          string             `bson:"location,omitempty" json:"location,omitempty"`
		LocationLongitude float64            `bson:"locationLongitude,omitempty" json:"locationLongitude,omitempty"`
		LocationLatitude  float64            `bson:"locationLatitude,omitempty" json:"locationLatitude,omitempty"`
		Description       string             `bson:"description,omitempty" json:"description,omitempty"`
		URL               string             `bson:"url,omitempty" json:"url,omitempty"`
	}
)

func NewMongoMapper(config *config.Config) IMongoMapper {
	conn := monc.MustNewModel(config.Mongo.URL, config.Mongo.DB, CollectionName, config.CacheConf)
	return &MongoMapper{
		conn: conn,
	}
}

func (m *MongoMapper) Insert(ctx context.Context, data *Photo) error {
	if data.ID.IsZero() {
		data.ID = primitive.NewObjectID()
		data.CreatedAt = time.Now()
		data.UpdatedAt = time.Now()
	}

	key := prefixPhotoCacheKey + data.ID.Hex()
	_, err := m.conn.InsertOne(ctx, key, data)
	return err
}

func (m *MongoMapper) FindOne(ctx context.Context, id string) (*Photo, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, consts.ErrInvalidObjectId
	}

	var data Photo
	key := prefixPhotoCacheKey + id
	err = m.conn.FindOne(ctx, key, &data, bson.M{consts.ID: oid})
	switch {
	case err == nil:
		return &data, nil
	case errors.Is(err, monc.ErrNotFound):
		return nil, consts.ErrNotFound
	default:
		return nil, err
	}
}

func (m *MongoMapper) Upsert(ctx context.Context, data *Photo) error {
	data.UpdatedAt = time.Now()
	key := prefixPhotoCacheKey + data.ID.Hex()
	filter := bson.M{"_id": data.ID}
	update := bson.M{
		"$set": data,
		"$setOnInsert": bson.M{
			"createdAt": time.Now(),
		},
	}
	option := options.UpdateOptions{}
	option.SetUpsert(true)

	_, err := m.conn.UpdateOne(ctx, key, filter, update, &option)
	return err
}

func (m *MongoMapper) List(ctx context.Context, albumId string, skip int64, count int64, onlyFeatured bool) ([]*Photo, int64, error) {
	var result []*Photo
	filter := bson.M{"albumId": albumId}
	if onlyFeatured {
		filter["isFeatured"] = true
	}
	err := m.conn.Find(ctx, &result, filter, &options.FindOptions{
		Skip:  &skip,
		Limit: &count,
		Sort:  bson.M{"createdAt": -1},
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := m.conn.CountDocuments(ctx, bson.M{"albumId": albumId})
	if err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func (m *MongoMapper) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return consts.ErrInvalidObjectId
	}
	key := prefixPhotoCacheKey + id
	filter := bson.M{"_id": oid}
	set := bson.M{"deletedAt": time.Now()}
	_, err = m.conn.UpdateOne(ctx, key, filter, bson.M{"$set": set})
	return err
}
