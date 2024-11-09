package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMedicineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMedicineLogic {
	return &CreateMedicineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMedicineLogic) CreateMedicine(in *core.MedicineInfo) (*core.BaseIDResp, error) {
    query := l.svcCtx.DB.Medicine.Create().
			SetNotNilSort(in.Sort).
			SetNotNilName(in.Name).
			SetNotNilQuantity(in.Quantity).
			SetNotNilDescription(in.Description).
			SetNotNilRemarks(in.Remarks).
			SetNotNilImages(in.Images)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	result, err := query.Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess }, nil
}
