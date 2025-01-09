package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMedicalRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMedicalRecordLogic {
	return &CreateMedicalRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateMedicalRecordLogic) CreateMedicalRecord(in *core.MedicalRecordInfo) (*core.BaseUUIDResp, error) {
    result, err := l.svcCtx.DB.MedicalRecord.Create().
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
			Save(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess }, nil
}
