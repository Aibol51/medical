package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAppointmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAppointmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAppointmentLogic {
	return &UpdateAppointmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAppointmentLogic) UpdateAppointment(in *core.AppointmentInfo) (*core.BaseResp, error) {
	err := l.svcCtx.DB.Appointment.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilPatientName(in.PatientName).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilIDCard(in.IdCard).
		SetNotNilGender(in.Gender).
		SetNotNilAge(in.Age).
		SetNotNilAppointmentTime(in.AppointmentTime).
		SetNotNilSymptoms(in.Symptoms).
		SetNotNilStatus(in.Status).
		SetNotNilRemarks(in.Remarks).
		SetNotNilUserID(in.UserId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
