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
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
