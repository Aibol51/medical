package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/medicalrecord"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetMedicalRecordListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMedicalRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMedicalRecordListLogic {
	return &GetMedicalRecordListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMedicalRecordListLogic) GetMedicalRecordList(in *core.MedicalRecordListReq) (*core.MedicalRecordListResp, error) {
	var predicates []predicate.MedicalRecord
	if in.PatientName != nil {
		predicates = append(predicates, medicalrecord.PatientNameContains(*in.PatientName))
	}
	if in.IdCardNumber != nil {
		predicates = append(predicates, medicalrecord.IDCardNumberContains(*in.IdCardNumber))
	}
	if in.PhoneNumber != nil {
		predicates = append(predicates, medicalrecord.PhoneNumberContains(*in.PhoneNumber))
	}
	result, err := l.svcCtx.DB.MedicalRecord.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.MedicalRecordListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.MedicalRecordInfo{
			Id:                 pointy.GetPointer(v.ID.String()),
			CreatedAt:          pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:          pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			PatientName:        &v.PatientName,
			Gender:             &v.Gender,
			Age:                &v.Age,
			IdCardNumber:       &v.IDCardNumber,
			PhoneNumber:        &v.PhoneNumber,
			ChiefComplaint:     &v.ChiefComplaint,
			PresentIllness:     &v.PresentIllness,
			PastHistory:        &v.PastHistory,
			SmokingHistory:     &v.SmokingHistory,
			DrinkingHistory:    &v.DrinkingHistory,
			AllergyHistory:     &v.AllergyHistory,
			HeartRate:          &v.HeartRate,
			BloodPressure:      &v.BloodPressure,
			OxygenSaturation:   &v.OxygenSaturation,
			BloodGlucose:       &v.BloodGlucose,
			Weight:             &v.Weight,
			WaistCircumference: &v.WaistCircumference,
			BodyFat:            &v.BodyFat,
			Diagnosis:          &v.Diagnosis,
			DietTherapy:        &v.DietTherapy,
			ExerciseTherapy:    &v.ExerciseTherapy,
			MedicationTherapy:  &v.MedicationTherapy,
			TreatmentPlan:      &v.TreatmentPlan,
			DoctorId:           &v.DoctorID,
			AppointmentId:      &v.AppointmentID,
			Remarks:            &v.Remarks,
			UserId:             &v.UserID,
		})
	}

	return resp, nil
}
