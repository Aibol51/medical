package swiper

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/swiper"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetSwiperListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSwiperListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSwiperListLogic {
	return &GetSwiperListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSwiperListLogic) GetSwiperList(in *core.SwiperListReq) (*core.SwiperListResp, error) {
	var predicates []predicate.Swiper
	if in.TitleZh != nil {
		predicates = append(predicates, swiper.TitleZhContains(*in.TitleZh))
	}
	if in.TitleEn != nil {
		predicates = append(predicates, swiper.TitleEnContains(*in.TitleEn))
	}
	if in.TitleRu != nil {
		predicates = append(predicates, swiper.TitleRuContains(*in.TitleRu))
	}
	result, err := l.svcCtx.DB.Swiper.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.SwiperListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.SwiperInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Sort:	&v.Sort,
			TitleZh:	&v.TitleZh,
			TitleEn:	&v.TitleEn,
			TitleRu:	&v.TitleRu,
			TitleKk:	&v.TitleKk,
			BannerZh:	&v.BannerZh,
			BannerEn:	&v.BannerEn,
			BannerRu:	&v.BannerRu,
			BannerKk:	&v.BannerKk,
			ContentZh:	&v.ContentZh,
			ContentEn:	&v.ContentEn,
			ContentRu:	&v.ContentRu,
			ContentKk:	&v.ContentKk,
			JumpUrl:	&v.JumpURL,
		})
	}

	return resp, nil
}
