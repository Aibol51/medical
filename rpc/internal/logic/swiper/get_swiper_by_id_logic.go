package swiper

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSwiperByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSwiperByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSwiperByIdLogic {
	return &GetSwiperByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSwiperByIdLogic) GetSwiperById(in *core.IDReq) (*core.SwiperInfo, error) {
	result, err := l.svcCtx.DB.Swiper.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.SwiperInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Sort:	&result.Sort,
		TitleZh:	&result.TitleZh,
		TitleEn:	&result.TitleEn,
		TitleRu:	&result.TitleRu,
		TitleKk:	&result.TitleKk,
		BannerZh:	&result.BannerZh,
		BannerEn:	&result.BannerEn,
		BannerRu:	&result.BannerRu,
		BannerKk:	&result.BannerKk,
		ContentZh:	&result.ContentZh,
		ContentEn:	&result.ContentEn,
		ContentRu:	&result.ContentRu,
		ContentKk:	&result.ContentKk,
		JumpUrl:	&result.JumpURL,
	}, nil
}

