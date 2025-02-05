package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExpertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExpertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExpertLogic {
	return &UpdateExpertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateExpertLogic) UpdateExpert(in *core.ExpertInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.Expert.UpdateOneID(*in.Id).
			SetNotNilSort(in.Sort).
			SetNotNilNameZh(in.NameZh).
			SetNotNilNameEn(in.NameEn).
			SetNotNilNameRu(in.NameRu).
			SetNotNilNameKk(in.NameKk).
			SetNotNilContentZh(in.ContentZh).
			SetNotNilContentEn(in.ContentEn).
			SetNotNilContentRu(in.ContentRu).
			SetNotNilContentKk(in.ContentKk).
			SetNotNilCoverURL(in.CoverUrl)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
