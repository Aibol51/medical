package medicine

import (
	"context"

    "github.com/suyuan32/simple-admin-core/rpc/ent/medicine"
    "github.com/suyuan32/simple-admin-core/rpc/internal/svc"
    "github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteMedicineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMedicineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMedicineLogic {
	return &DeleteMedicineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMedicineLogic) DeleteMedicine(in *core.IDsReq) (*core.BaseResp, error) {
	_, err := l.svcCtx.DB.Medicine.Delete().Where(medicine.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
