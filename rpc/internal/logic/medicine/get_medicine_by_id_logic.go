package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicineByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicineByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicineByIdLogic {
	return &GetMedicineByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMedicineByIdLogic) GetMedicineById(in *core.IDReq) (*core.MedicineInfo, error) {
	result, err := l.svcCtx.DB.Medicine.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MedicineInfo{
		Id:          &result.ID,
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:	pointy.GetPointer(uint32(result.Status)),
		Sort:	&result.Sort,
		Name:	&result.Name,
		Quantity:	&result.Quantity,
		Description:	&result.Description,
		Remarks:	&result.Remarks,
		Images:	&result.Images,
	}, nil
}

