package medicine

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-core/api/internal/logic/medicine"
	"github.com/suyuan32/simple-admin-core/api/internal/svc"
	"github.com/suyuan32/simple-admin-core/api/internal/types"
)

// swagger:route post /medicine/create medicine CreateMedicine
//
// Create medicine information | 创建Medicine信息
//
// Create medicine information | 创建Medicine信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: MedicineInfo
//
// Responses:
//  200: BaseMsgResp

func CreateMedicineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MedicineInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := medicine.NewCreateMedicineLogic(r.Context(), svcCtx)
		resp, err := l.CreateMedicine(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
