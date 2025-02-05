package service

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateServiceLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateServiceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateServiceLogic {
	return &UpdateServiceLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateServiceLogic) UpdateService(req *types.ServiceInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateService(l.ctx,
		&core.ServiceInfo{
			Id:            req.Id,
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
