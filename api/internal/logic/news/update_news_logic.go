package news

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNewsLogic {
	return &UpdateNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNewsLogic) UpdateNews(req *types.NewsInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateNews(l.ctx,
		&core.NewsInfo{
			Id:        req.Id,
			Status:    req.Status,
			Sort:      req.Sort,
			TitleZh:   req.TitleZh,
			TitleEn:   req.TitleEn,
			TitleRu:   req.TitleRu,
			TitleKk:   req.TitleKk,
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
