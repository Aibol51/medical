package swiper

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSwiperLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSwiperLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSwiperLogic {
	return &UpdateSwiperLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSwiperLogic) UpdateSwiper(in *core.SwiperInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.Swiper.UpdateOneID(*in.Id).
			SetNotNilSort(in.Sort).
			SetNotNilTitleZh(in.TitleZh).
			SetNotNilTitleEn(in.TitleEn).
			SetNotNilTitleRu(in.TitleRu).
			SetNotNilTitleKk(in.TitleKk).
			SetNotNilBannerZh(in.BannerZh).
			SetNotNilBannerEn(in.BannerEn).
			SetNotNilBannerRu(in.BannerRu).
			SetNotNilBannerKk(in.BannerKk).
			SetNotNilContentZh(in.ContentZh).
			SetNotNilContentEn(in.ContentEn).
			SetNotNilContentRu(in.ContentRu).
			SetNotNilContentKk(in.ContentKk).
			SetNotNilJumpURL(in.JumpUrl)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
