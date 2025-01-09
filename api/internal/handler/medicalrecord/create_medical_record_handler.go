package medicalrecord

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/medicalrecord"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /medical_record/create medicalrecord CreateMedicalRecord
//
// Create medical record information | 创建MedicalRecord信息
//
// Create medical record information | 创建MedicalRecord信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MedicalRecordInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMedicalRecordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicalRecordInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicalrecord.NewCreateMedicalRecordLogic(r.Context(), svcCtx)
		resp, err := l.CreateMedicalRecord(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
