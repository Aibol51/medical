package medicalrecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMedicalRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateMedicalRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMedicalRecordLogic {
	return &UpdateMedicalRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateMedicalRecordLogic) UpdateMedicalRecord(req *types.MedicalRecordInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.CoreRpc.UpdateMedicalRecord(l.ctx,
		&core.MedicalRecordInfo{
			Id:                 req.Id,
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
