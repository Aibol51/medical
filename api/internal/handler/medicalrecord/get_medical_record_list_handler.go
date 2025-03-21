package medicalrecord

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/medicalrecord"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /medical_record/list medicalrecord GetMedicalRecordList
//
// Get medical record list | 获取MedicalRecord信息列表
//
// Get medical record list | 获取MedicalRecord信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MedicalRecordListReq
//
// Responses:
//  200: MedicalRecordListResp

func GetMedicalRecordListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicalRecordListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicalrecord.NewGetMedicalRecordListLogic(r.Context(), svcCtx)
		resp, err := l.GetMedicalRecordList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
