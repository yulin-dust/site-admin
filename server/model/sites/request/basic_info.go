
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BasicInfoSearch struct{
    CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
      Domain  *string `json:"domain" form:"domain"` 
      CmsType  *string `json:"cmsType" form:"cmsType"` 
      Status  *string `json:"status" form:"status"` 
    request.PageInfo
}
