package service

import (
	"context"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
	albummapper "github.com/xh-polaris/meowcloud-content/biz/infrastructure/mapper/album"
	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"
)

type AlbumService interface {
	CreateAlbum(ctx context.Context, req *content.CreateAlbumReq) (res *content.CreateAlbumResp, err error)
	RetrieveAlbum(ctx context.Context, req *content.RetrieveAlbumReq) (res *content.RetrieveAlbumResp, err error)
	UpdateAlbum(ctx context.Context, req *content.UpdateAlbumReq) (res *content.UpdateAlbumResp, err error)
	DeleteAlbum(ctx context.Context, req *content.DeleteAlbumReq) (res *content.DeleteAlbumResp, err error)
	ListAlbum(ctx context.Context, req *content.ListAlbumReq) (res *content.ListAlbumResp, err error)
}

type AlbumServiceImpl struct {
	Config           *config.Config
	AlbumMongoMapper albummapper.IMongoMapper
}

var AlbumSet = wire.NewSet(
	wire.Struct(new(AlbumServiceImpl), "*"),
	wire.Bind(new(AlbumService), new(*AlbumServiceImpl)),
)

func (s *AlbumServiceImpl) CreateAlbum(ctx context.Context, req *content.CreateAlbumReq) (res *content.CreateAlbumResp, err error) {
	album := &albummapper.Album{}
	err = copier.Copy(album, req.Album)
	if err != nil {
		return nil, err
	}
	err = s.AlbumMongoMapper.Insert(ctx, album)
	if err != nil {
		return nil, err
	}
	return &content.CreateAlbumResp{Album: &content.Album{
		Id: album.ID.Hex(),
	}}, nil
}

func (s *AlbumServiceImpl) RetrieveAlbum(ctx context.Context, req *content.RetrieveAlbumReq) (res *content.RetrieveAlbumResp, err error) {
	album, err := s.AlbumMongoMapper.FindOne(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	c := &content.Album{}
	err = copier.Copy(c, album)
	if err != nil {
		return nil, err
	}
	return &content.RetrieveAlbumResp{Album: &content.Album{}}, nil
}

func (s *AlbumServiceImpl) UpdateAlbum(ctx context.Context, req *content.UpdateAlbumReq) (res *content.UpdateAlbumResp, err error) {
	album := &albummapper.Album{}
	err = copier.Copy(album, req.Album)
	if err != nil {
		return nil, err
	}
	err = s.AlbumMongoMapper.Upsert(ctx, album)
	if err != nil {
		return nil, err
	}
	return &content.UpdateAlbumResp{}, nil
}

func (s *AlbumServiceImpl) DeleteAlbum(ctx context.Context, req *content.DeleteAlbumReq) (res *content.DeleteAlbumResp, err error) {
	err = s.AlbumMongoMapper.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &content.DeleteAlbumResp{}, nil
}

func (s *AlbumServiceImpl) ListAlbum(ctx context.Context, req *content.ListAlbumReq) (res *content.ListAlbumResp, err error) {
	albums, count, err := s.AlbumMongoMapper.FindMany(ctx, req.PaginationOptions.GetOffset(), req.PaginationOptions.GetLimit())
	if err != nil {
		return nil, err
	}

	var catAlbumList []*content.Album
	for _, album := range albums {
		catAlbum := &content.Album{}
		err = copier.Copy(catAlbum, album)
		if err != nil {
			return nil, err
		}
		catAlbumList = append(catAlbumList, catAlbum)
	}

	return &content.ListAlbumResp{Albums: catAlbumList, Total: count}, nil
}
