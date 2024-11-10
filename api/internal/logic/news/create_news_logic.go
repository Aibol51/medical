package news

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateNewsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateNewsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateNewsLogic {
	return &CreateNewsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateNewsLogic) CreateNews(req *types.NewsInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateNews(l.ctx,
		&core.NewsInfo{
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
