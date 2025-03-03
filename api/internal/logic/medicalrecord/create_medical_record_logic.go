package medicalrecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMedicalRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMedicalRecordLogic {
	return &CreateMedicalRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMedicalRecordLogic) CreateMedicalRecord(req *types.MedicalRecordInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.CreateMedicalRecord(l.ctx,
		&core.MedicalRecordInfo{
			PatientName:        req.PatientName,
			Gender:             req.Gender,
			Age:                req.Age,
			IdCardNumber:       req.IdCardNumber,
			PhoneNumber:        req.PhoneNumber,
			ChiefComplaint:     req.ChiefComplaint,
			PresentIllness:     req.PresentIllness,
			PastHistory:        req.PastHistory,
			SmokingHistory:     req.SmokingHistory,
			DrinkingHistory:    req.DrinkingHistory,
			AllergyHistory:     req.AllergyHistory,
			HeartRate:          req.HeartRate,
			BloodPressure:      req.BloodPressure,
			OxygenSaturation:   req.OxygenSaturation,
			BloodGlucose:       req.BloodGlucose,
			Weight:             req.Weight,
			WaistCircumference: req.WaistCircumference,
			BodyFat:            req.BodyFat,
			Diagnosis:          req.Diagnosis,
			DietTherapy:        req.DietTherapy,
			ExerciseTherapy:    req.ExerciseTherapy,
			MedicationTherapy:  req.MedicationTherapy,
			TreatmentPlan:      req.TreatmentPlan,
			DoctorId:           req.DoctorId,
			AppointmentId:      req.AppointmentId,
			Remarks:            req.Remarks,
			UserId:             req.UserId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
