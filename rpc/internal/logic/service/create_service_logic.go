package service

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateServiceLogic) CreateService(in *core.ServiceInfo) (*core.BaseIDResp, error) {
    query := l.svcCtx.DB.Service.Create().
			SetNotNilSort(in.Sort).
			SetNotNilNameZh(in.NameZh).
			SetNotNilNameEn(in.NameEn).
			SetNotNilNameRu(in.NameRu).
			SetNotNilNameKk(in.NameKk).
			SetNotNilDescriptionZh(in.DescriptionZh).
			SetNotNilDescriptionEn(in.DescriptionEn).
			SetNotNilDescriptionRu(in.DescriptionRu).
			SetNotNilDescriptionKk(in.DescriptionKk).
			SetNotNilCoverURL(in.CoverUrl)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
