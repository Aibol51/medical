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

type UpdateMedicineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMedicineLogic {
	return &UpdateMedicineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMedicineLogic) UpdateMedicine(in *core.MedicineInfo) (*core.BaseResp, error) {
	query:= l.svcCtx.DB.Medicine.UpdateOneID(*in.Id).
			SetNotNilSort(in.Sort).
			SetNotNilNameZh(in.NameZh).
			SetNotNilNameEn(in.NameEn).
			SetNotNilNameRu(in.NameRu).
			SetNotNilNameKk(in.NameKk).
			SetNotNilQuantity(in.Quantity).
			SetNotNilDescriptionZh(in.DescriptionZh).
			SetNotNilDescriptionEn(in.DescriptionEn).
			SetNotNilDescriptionRu(in.DescriptionRu).
			SetNotNilDescriptionKk(in.DescriptionKk).
			SetNotNilRemarks(in.Remarks).
			SetNotNilImages(in.Images)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}

	 err := query.Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
