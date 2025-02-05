package expert

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetExpertByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExpertByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExpertByIdLogic {
	return &GetExpertByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetExpertByIdLogic) GetExpertById(in *core.IDReq) (*core.ExpertInfo, error) {
	result, err := l.svcCtx.DB.Expert.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.ExpertInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Sort:	&result.Sort,
		NameZh:	&result.NameZh,
		NameEn:	&result.NameEn,
		NameRu:	&result.NameRu,
		NameKk:	&result.NameKk,
		ContentZh:	&result.ContentZh,
		ContentEn:	&result.ContentEn,
		ContentRu:	&result.ContentRu,
		ContentKk:	&result.ContentKk,
		CoverUrl:	&result.CoverURL,
	}, nil
}

