package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMedicalRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMedicalRecordLogic {
	return &UpdateMedicalRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMedicalRecordLogic) UpdateMedicalRecord(in *core.MedicalRecordInfo) (*core.BaseResp, error) {
	err:= l.svcCtx.DB.MedicalRecord.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
			SetNotNilPatientName(in.PatientName).
			SetNotNilPhoneNumber(in.PhoneNumber).
			SetNotNilGender(in.Gender).
			SetNotNilAge(in.Age).
			SetNotNilVisitTime(in.VisitTime).
			SetNotNilDiagnosis(in.Diagnosis).
			SetNotNilTreatmentPlan(in.TreatmentPlan).
			SetNotNilPrescription(in.Prescription).
			SetNotNilExaminationResults(in.ExaminationResults).
			SetNotNilDoctorAdvice(in.DoctorAdvice).
			SetNotNilDoctorID(in.DoctorId).
			SetNotNilDepartment(in.Department).
			SetNotNilAppointmentID(in.AppointmentId).
			SetNotNilRemarks(in.Remarks).
			SetNotNilUserID(in.UserId).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseResp{Msg: i18n.UpdateSuccess }, nil
}
