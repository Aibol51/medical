package appointment

import (
	"context"

	"github.com/suyuan32/simple-admin-core/rpc/ent/appointment"
	"github.com/suyuan32/simple-admin-core/rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppointmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppointmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppointmentListLogic {
	return &GetAppointmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppointmentListLogic) GetAppointmentList(in *core.AppointmentListReq) (*core.AppointmentListResp, error) {
	var predicates []predicate.Appointment
	if in.PatientName != nil {
		predicates = append(predicates, appointment.PatientNameContains(*in.PatientName))
	}
	if in.PhoneNumber != nil {
		predicates = append(predicates, appointment.PhoneNumberContains(*in.PhoneNumber))
	}
	if in.IdCard != nil {
		predicates = append(predicates, appointment.IDCardContains(*in.IdCard))
	}
	result, err := l.svcCtx.DB.Appointment.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &core.AppointmentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.AppointmentInfo{
			Id:              pointy.GetPointer(v.ID.String()),
			CreatedAt:       pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:       pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			PatientName:     &v.PatientName,
			PhoneNumber:     &v.PhoneNumber,
			IdCard:          &v.IDCard,
			Gender:          &v.Gender,
			Age:             &v.Age,
			AppointmentTime: &v.AppointmentTime,
			Symptoms:        &v.Symptoms,
			Status:          &v.Status,
			Remarks:         &v.Remarks,
			UserId:          &v.UserID,
		})
	}

	return resp, nil
}
