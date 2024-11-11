package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppointmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppointmentLogic {
	return &CreateAppointmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAppointmentLogic) CreateAppointment(req *types.AppointmentInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateAppointment(l.ctx,
		&core.AppointmentInfo{
			PatientName:     req.PatientName,
			PhoneNumber:     req.PhoneNumber,
			IdCard:          req.IdCard,
			Gender:          req.Gender,
			Age:             req.Age,
			AppointmentTime: req.AppointmentTime,
			Symptoms:        req.Symptoms,
			Status:          req.Status,
			Remarks:         req.Remarks,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
