package medicalrecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMedicalRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMedicalRecordLogic {
	return &DeleteMedicalRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMedicalRecordLogic) DeleteMedicalRecord(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.DeleteMedicalRecord(l.ctx, &core.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
