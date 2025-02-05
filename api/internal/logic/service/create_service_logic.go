package service

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateServiceLogic {
	return &CreateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateServiceLogic) CreateService(req *types.ServiceInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateService(l.ctx,
		&core.ServiceInfo{
			Status:        req.Status,
			Sort:          req.Sort,
			NameZh:        req.NameZh,
			NameEn:        req.NameEn,
			NameRu:        req.NameRu,
			NameKk:        req.NameKk,
			DescriptionZh: req.DescriptionZh,
			DescriptionEn: req.DescriptionEn,
			DescriptionRu: req.DescriptionRu,
			DescriptionKk: req.DescriptionKk,
			CoverUrl:      req.CoverUrl,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
