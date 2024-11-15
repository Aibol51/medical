package swiper

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSwiperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSwiperLogic {
	return &CreateSwiperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSwiperLogic) CreateSwiper(req *types.SwiperInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateSwiper(l.ctx,
		&core.SwiperInfo{
			Status:    req.Status,
			Sort:      req.Sort,
			TitleZh:   req.TitleZh,
			TitleEn:   req.TitleEn,
			TitleRu:   req.TitleRu,
			TitleKk:   req.TitleKk,
			BannerZh:  req.BannerZh,
			BannerEn:  req.BannerEn,
			BannerRu:  req.BannerRu,
			BannerKk:  req.BannerKk,
			ContentZh: req.ContentZh,
			ContentEn: req.ContentEn,
			ContentRu: req.ContentRu,
			ContentKk: req.ContentKk,
			JumpUrl:   req.JumpUrl,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
