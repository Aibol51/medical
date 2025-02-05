package service

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetServiceByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetServiceByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetServiceByIdLogic {
	return &GetServiceByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetServiceByIdLogic) GetServiceById(in *core.IDReq) (*core.ServiceInfo, error) {
	result, err := l.svcCtx.DB.Service.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.ServiceInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Sort:	&result.Sort,
		NameZh:	&result.NameZh,
		NameEn:	&result.NameEn,
		NameRu:	&result.NameRu,
		NameKk:	&result.NameKk,
		DescriptionZh:	&result.DescriptionZh,
		DescriptionEn:	&result.DescriptionEn,
		DescriptionRu:	&result.DescriptionRu,
		DescriptionKk:	&result.DescriptionKk,
		CoverUrl:	&result.CoverURL,
	}, nil
}

