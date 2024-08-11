package adaptor

import (
	"context"

	"github.com/xh-polaris/service-idl-gen-go/kitex_gen/meowcloud/content"

	"github.com/xh-polaris/meowcloud-content/biz/application/service"
	"github.com/xh-polaris/meowcloud-content/biz/infrastructure/config"
)

type ContentServerImpl struct {
	*config.Config
	AlbumService service.AlbumService
	PhotoService service.PhotoService
}

func (s *ContentServerImpl) CreateAlbum(ctx context.Context, req *content.CreateAlbumReq) (res *content.CreateAlbumResp, err error) {
	return s.AlbumService.CreateAlbum(ctx, req)
}

func (s *ContentServerImpl) RetrieveAlbum(ctx context.Context, req *content.RetrieveAlbumReq) (res *content.RetrieveAlbumResp, err error) {
	return s.AlbumService.RetrieveAlbum(ctx, req)
}

func (s *ContentServerImpl) UpdateAlbum(ctx context.Context, req *content.UpdateAlbumReq) (res *content.UpdateAlbumResp, err error) {
	return s.AlbumService.UpdateAlbum(ctx, req)
}

func (s *ContentServerImpl) DeleteAlbum(ctx context.Context, req *content.DeleteAlbumReq) (res *content.DeleteAlbumResp, err error) {
	return s.AlbumService.DeleteAlbum(ctx, req)
}

func (s *ContentServerImpl) ListAlbum(ctx context.Context, req *content.ListAlbumReq) (res *content.ListAlbumResp, err error) {
	return s.AlbumService.ListAlbum(ctx, req)
}

func (s *ContentServerImpl) CreatePhoto(ctx context.Context, req *content.CreatePhotoReq) (res *content.CreatePhotoResp, err error) {
	return s.PhotoService.CreatePhoto(ctx, req)
}

func (s *ContentServerImpl) RetrievePhoto(ctx context.Context, req *content.RetrievePhotoReq) (res *content.RetrievePhotoResp, err error) {
	return s.PhotoService.RetrievePhoto(ctx, req)
}

func (s *ContentServerImpl) UpdatePhoto(ctx context.Context, req *content.UpdatePhotoReq) (res *content.UpdatePhotoResp, err error) {
	return s.PhotoService.UpdatePhoto(ctx, req)
}

func (s *ContentServerImpl) DeletePhoto(ctx context.Context, req *content.DeletePhotoReq) (res *content.DeletePhotoResp, err error) {
	return s.PhotoService.DeletePhoto(ctx, req)
}

func (s *ContentServerImpl) ListPhoto(ctx context.Context, req *content.ListPhotoReq) (res *content.ListPhotoResp, err error) {
	return s.PhotoService.ListPhoto(ctx, req)
}
