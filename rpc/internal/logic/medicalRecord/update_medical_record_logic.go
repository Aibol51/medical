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
	err := l.svcCtx.DB.MedicalRecord.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilPatientName(in.PatientName).
		SetNotNilGender(in.Gender).
		SetNotNilAge(in.Age).
		SetNotNilIDCardNumber(in.IdCardNumber).
		SetNotNilPhoneNumber(in.PhoneNumber).
		SetNotNilChiefComplaint(in.ChiefComplaint).
		SetNotNilPresentIllness(in.PresentIllness).
		SetNotNilPastHistory(in.PastHistory).
		SetNotNilSmokingHistory(in.SmokingHistory).
		SetNotNilDrinkingHistory(in.DrinkingHistory).
		SetNotNilAllergyHistory(in.AllergyHistory).
		SetNotNilHeartRate(in.HeartRate).
		SetNotNilBloodPressure(in.BloodPressure).
		SetNotNilOxygenSaturation(in.OxygenSaturation).
		SetNotNilBloodGlucose(in.BloodGlucose).
		SetNotNilWeight(in.Weight).
		SetNotNilWaistCircumference(in.WaistCircumference).
		SetNotNilBodyFat(in.BodyFat).
		SetNotNilDiagnosis(in.Diagnosis).
		SetNotNilDietTherapy(in.DietTherapy).
		SetNotNilExerciseTherapy(in.ExerciseTherapy).
		SetNotNilMedicationTherapy(in.MedicationTherapy).
		SetNotNilTreatmentPlan(in.TreatmentPlan).
		SetNotNilDoctorID(in.DoctorId).
		SetNotNilAppointmentID(in.AppointmentId).
		SetNotNilRemarks(in.Remarks).
		SetNotNilUserID(in.UserId).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
