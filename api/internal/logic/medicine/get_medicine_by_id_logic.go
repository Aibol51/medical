package medicine

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicineByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMedicineByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicineByIdLogic {
	return &GetMedicineByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMedicineByIdLogic) GetMedicineById(req *types.IDReq) (resp *types.MedicineInfoResp, err error) {
	data, err := l.svcCtx.CoreRpc.GetMedicineById(l.ctx, &core.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.MedicineInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.MedicineInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:      data.Status,
			Sort:        data.Sort,
			Name:        data.Name,
			Quantity:    data.Quantity,
			Description: data.Description,
			Remarks:     data.Remarks,
			Images:      data.Images,
		},
	}, nil
}
