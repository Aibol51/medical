package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExpertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateExpertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExpertLogic {
	return &UpdateExpertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExpertLogic) UpdateExpert(req *types.ExpertInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateExpert(l.ctx,
		&core.ExpertInfo{
			Id:        req.Id,
			Status:    req.Status,
			Sort:      req.Sort,
			NameZh:    req.NameZh,
			NameEn:    req.NameEn,
			NameRu:    req.NameRu,
			NameKk:    req.NameKk,
			ContentZh: req.ContentZh,
			ContentEn: req.ContentEn,
			ContentRu: req.ContentRu,
			ContentKk: req.ContentKk,
			CoverUrl:  req.CoverUrl,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
