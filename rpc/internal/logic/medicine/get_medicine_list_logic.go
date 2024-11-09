package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/medicine"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetMedicineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicineListLogic {
	return &GetMedicineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMedicineListLogic) GetMedicineList(in *core.MedicineListReq) (*core.MedicineListResp, error) {
	var predicates []predicate.Medicine
	if in.Name != nil {
		predicates = append(predicates, medicine.NameContains(*in.Name))
	}
	if in.Description != nil {
		predicates = append(predicates, medicine.DescriptionContains(*in.Description))
	}
	if in.Remarks != nil {
		predicates = append(predicates, medicine.RemarksContains(*in.Remarks))
	}
	result, err := l.svcCtx.DB.Medicine.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MedicineListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MedicineInfo{
			Id:          &v.ID,
			CreatedAt:   pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:   pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:	pointy.GetPointer(uint32(v.Status)),
			Sort:	&v.Sort,
			Name:	&v.Name,
			Quantity:	&v.Quantity,
			Description:	&v.Description,
			Remarks:	&v.Remarks,
			Images:	&v.Images,
		})
	}

	return resp, nil
}
