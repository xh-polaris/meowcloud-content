package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
	photo "github.com/xh-polaris/meowcloud-content/biz/infrastructure/mapper/photo"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
)

type PhotoService interface {
	CreatePhoto(ctx context.Context, req *content.CreatePhotoReq) (*content.CreatePhotoResp, error)
	RetrievePhoto(ctx context.Context, req *content.RetrievePhotoReq) (*content.RetrievePhotoResp, error)
	UpdatePhoto(ctx context.Context, req *content.UpdatePhotoReq) (*content.UpdatePhotoResp, error)
	DeletePhoto(ctx context.Context, req *content.DeletePhotoReq) (*content.DeletePhotoResp, error)
	ListPhoto(ctx context.Context, req *content.ListPhotoReq) (*content.ListPhotoResp, error)
}

type PhotoServiceImpl struct {
	Config           *config.Config
	PhotoMongoMapper photo.MongoMapper
}

var PhotoSet = wire.NewSet(
	wire.Struct(new(PhotoServiceImpl), "*"),
	wire.Bind(new(PhotoService), new(*PhotoServiceImpl)),
)

func (s *PhotoServiceImpl) CreatePhoto(ctx context.Context, req *content.CreatePhotoReq) (*content.CreatePhotoResp, error) {
	p := &photo.Photo{}
	err := copier.Copy(p, req.Photo)
	if err != nil {
		return nil, err
	}
	err = s.PhotoMongoMapper.Insert(ctx, p)
	if err != nil {
		return nil, err
	}
	p1 := &content.Photo{}
	err = copier.Copy(p1, p)
	if err != nil {
		return nil, err
	}
	return &content.CreatePhotoResp{Photo: p1}, nil
}

func (s *PhotoServiceImpl) RetrievePhoto(ctx context.Context, req *content.RetrievePhotoReq) (*content.RetrievePhotoResp, error) {
	p, err := s.PhotoMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	p1 := &content.Photo{}
	err = copier.Copy(p1, p)
	if err != nil {
		return nil, err
	}
	return &content.RetrievePhotoResp{Photo: p1}, nil
}

func (s *PhotoServiceImpl) UpdatePhoto(ctx context.Context, req *content.UpdatePhotoReq) (*content.UpdatePhotoResp, error) {
	p := &photo.Photo{}
	err := copier.Copy(p, req.Photo)
	if err != nil {
		return nil, err
	}
	err = s.PhotoMongoMapper.Upsert(ctx, p)
	if err != nil {
		return nil, err
	}
	return &content.UpdatePhotoResp{}, nil
}

func (s *PhotoServiceImpl) DeletePhoto(ctx context.Context, req *content.DeletePhotoReq) (*content.DeletePhotoResp, error) {
	err := s.PhotoMongoMapper.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &content.DeletePhotoResp{}, nil
}

func (s *PhotoServiceImpl) ListPhoto(ctx context.Context, req *content.ListPhotoReq) (*content.ListPhotoResp, error) {
	var onlyFeatured bool
	if req.OnlyFeatured != nil {
		onlyFeatured = *req.OnlyFeatured
	} else {
		onlyFeatured = false
	}
	photos, cnt, err := s.PhotoMongoMapper.List(ctx, req.AlbumId, req.PaginationOptions.GetOffset(), req.PaginationOptions.GetLimit(), onlyFeatured)
	if err != nil {
		return nil, err
	}
	p := make([]*content.Photo, 0, len(photos))
	for _, p0 := range photos {
		p1 := &content.Photo{}
		err = copier.Copy(p1, p0)
		if err != nil {
			return nil, err
		}
		p = append(p, p1)
	}
	return &content.ListPhotoResp{Photos: p, Total: cnt}, nil
}
