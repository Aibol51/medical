package medicalRecord

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

    "github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
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
		Id:          pointy.GetPointer(result.ID.String()),
		CreatedAt:    pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:    pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		PatientName:	&result.PatientName,
		PhoneNumber:	&result.PhoneNumber,
		Gender:	&result.Gender,
		Age:	&result.Age,
		VisitTime:	&result.VisitTime,
		Diagnosis:	&result.Diagnosis,
		TreatmentPlan:	&result.TreatmentPlan,
		Prescription:	&result.Prescription,
		ExaminationResults:	&result.ExaminationResults,
		DoctorAdvice:	&result.DoctorAdvice,
		DoctorId:	&result.DoctorID,
		Department:	&result.Department,
		AppointmentId:	&result.AppointmentID,
		Remarks:	&result.Remarks,
		UserId:	&result.UserID,
	}, nil
}

