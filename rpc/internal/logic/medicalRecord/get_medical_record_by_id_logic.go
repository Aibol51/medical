package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicalRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicalRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicalRecordByIdLogic {
	return &GetMedicalRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMedicalRecordByIdLogic) GetMedicalRecordById(in *core.UUIDReq) (*core.MedicalRecordInfo, error) {
	result, err := l.svcCtx.DB.MedicalRecord.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &core.MedicalRecordInfo{
		Id:                 pointy.GetPointer(result.ID.String()),
		CreatedAt:          pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:          pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		PatientName:        &result.PatientName,
		Gender:             &result.Gender,
		Age:                &result.Age,
		IdCardNumber:       &result.IDCardNumber,
		PhoneNumber:        &result.PhoneNumber,
		ChiefComplaint:     &result.ChiefComplaint,
		PresentIllness:     &result.PresentIllness,
		PastHistory:        &result.PastHistory,
		SmokingHistory:     &result.SmokingHistory,
		DrinkingHistory:    &result.DrinkingHistory,
		AllergyHistory:     &result.AllergyHistory,
		HeartRate:          &result.HeartRate,
		BloodPressure:      &result.BloodPressure,
		OxygenSaturation:   &result.OxygenSaturation,
		BloodGlucose:       &result.BloodGlucose,
		Weight:             &result.Weight,
		WaistCircumference: &result.WaistCircumference,
		BodyFat:            &result.BodyFat,
		Diagnosis:          &result.Diagnosis,
		DietTherapy:        &result.DietTherapy,
		ExerciseTherapy:    &result.ExerciseTherapy,
		MedicationTherapy:  &result.MedicationTherapy,
		TreatmentPlan:      &result.TreatmentPlan,
		DoctorId:           &result.DoctorID,
		AppointmentId:      &result.AppointmentID,
		Remarks:            &result.Remarks,
		UserId:             &result.UserID,
	}, nil
}
