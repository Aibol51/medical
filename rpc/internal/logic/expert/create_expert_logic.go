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

type CreateExpertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateExpertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExpertLogic {
	return &CreateExpertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateExpertLogic) CreateExpert(in *core.ExpertInfo) (*core.BaseIDResp, error) {
    query := l.svcCtx.DB.Expert.Create().
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

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
