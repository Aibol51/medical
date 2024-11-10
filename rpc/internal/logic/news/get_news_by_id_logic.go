package news

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetNewsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNewsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsByIdLogic {
	return &GetNewsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNewsByIdLogic) GetNewsById(in *core.IDReq) (*core.NewsInfo, error) {
	result, err := l.svcCtx.DB.News.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.NewsInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Sort:	&result.Sort,
		TitleZh:	&result.TitleZh,
		TitleEn:	&result.TitleEn,
		TitleRu:	&result.TitleRu,
		TitleKk:	&result.TitleKk,
		ContentZh:	&result.ContentZh,
		ContentEn:	&result.ContentEn,
		ContentRu:	&result.ContentRu,
		ContentKk:	&result.ContentKk,
		CoverUrl:	&result.CoverURL,
	}, nil
}

