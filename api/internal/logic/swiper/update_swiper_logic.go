package swiper

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSwiperLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSwiperLogic {
	return &UpdateSwiperLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSwiperLogic) UpdateSwiper(req *types.SwiperInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateSwiper(l.ctx,
		&core.SwiperInfo{
			Id:        req.Id,
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
