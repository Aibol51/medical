package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMedicineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMedicineLogic {
	return &CreateMedicineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMedicineLogic) CreateMedicine(req *types.MedicineInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateMedicine(l.ctx,
		&core.MedicineInfo{
			Status:        req.Status,
			Sort:          req.Sort,
			NameZh:        req.NameZh,
			NameEn:        req.NameEn,
			NameRu:        req.NameRu,
			NameKk:        req.NameKk,
			Quantity:      req.Quantity,
			DescriptionZh: req.DescriptionZh,
			DescriptionEn: req.DescriptionEn,
			DescriptionRu: req.DescriptionRu,
			DescriptionKk: req.DescriptionKk,
			Remarks:       req.Remarks,
			Images:        req.Images,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
