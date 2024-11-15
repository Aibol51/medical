package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppointmentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentByIdLogic {
	return &GetAppointmentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppointmentByIdLogic) GetAppointmentById(in *core.UUIDReq) (*core.AppointmentInfo, error) {
	result, err := l.svcCtx.DB.Appointment.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.AppointmentInfo{
		Id:              pointy.GetPointer(result.ID.String()),
		CreatedAt:       pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:       pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		PatientName:     &result.PatientName,
		PhoneNumber:     &result.PhoneNumber,
		IdCard:          &result.IDCard,
		Gender:          &result.Gender,
		Age:             &result.Age,
		AppointmentTime: &result.AppointmentTime,
		Symptoms:        &result.Symptoms,
		Status:          &result.Status,
		Remarks:         &result.Remarks,
		UserId:          &result.UserID,
	}, nil
}
