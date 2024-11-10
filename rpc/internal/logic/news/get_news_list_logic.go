package news

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/news"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetNewsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNewsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNewsListLogic {
	return &GetNewsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNewsListLogic) GetNewsList(in *core.NewsListReq) (*core.NewsListResp, error) {
	var predicates []predicate.News
	if in.TitleZh != nil {
		predicates = append(predicates, news.TitleZhContains(*in.TitleZh))
	}
	if in.TitleEn != nil {
		predicates = append(predicates, news.TitleEnContains(*in.TitleEn))
	}
	if in.TitleRu != nil {
		predicates = append(predicates, news.TitleRuContains(*in.TitleRu))
	}
	result, err := l.svcCtx.DB.News.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.NewsListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.NewsInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Sort:	&v.Sort,
			TitleZh:	&v.TitleZh,
			TitleEn:	&v.TitleEn,
			TitleRu:	&v.TitleRu,
			TitleKk:	&v.TitleKk,
			ContentZh:	&v.ContentZh,
			ContentEn:	&v.ContentEn,
			ContentRu:	&v.ContentRu,
			ContentKk:	&v.ContentKk,
			CoverUrl:	&v.CoverURL,
		})
	}

	return resp, nil
}
