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
			PhoneNumber:        req.PhoneNumber,
			Gender:             req.Gender,
			Age:                req.Age,
			VisitTime:          req.VisitTime,
			Diagnosis:          req.Diagnosis,
			TreatmentPlan:      req.TreatmentPlan,
			Prescription:       req.Prescription,
			ExaminationResults: req.ExaminationResults,
			DoctorAdvice:       req.DoctorAdvice,
			DoctorId:           req.DoctorId,
			Department:         req.Department,
			AppointmentId:      req.AppointmentId,
			Remarks:            req.Remarks,
			UserId:             req.UserId,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
