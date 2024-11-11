package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppointmentLogic {
	return &CreateAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAppointmentLogic) CreateAppointment(in *core.AppointmentInfo) (*core.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Appointment.Create().
		SetNotNilPatientName(in.PatientName).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilIDCard(in.IdCard).
		SetNotNilGender(in.Gender).
		SetNotNilAge(in.Age).
		SetNotNilAppointmentTime(in.AppointmentTime).
		SetNotNilSymptoms(in.Symptoms).
		SetNotNilStatus(in.Status).
		SetNotNilRemarks(in.Remarks).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
