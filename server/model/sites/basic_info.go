// 自动生成模板BasicInfo
package sites

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 基础信息 结构体  BasicInfo
type BasicInfo struct {
	global.GVA_MODEL
	Domain      *string `json:"domain" form:"domain" gorm:"uniqueIndex;comment:域名;column:domain;size:128;" binding:"required"`                //域名
	Url         *string `json:"url" form:"url" gorm:"comment:完整的链接;column:url;size:256;" binding:"required"`                                  //链接
	ApiToken    *string `json:"apiToken" form:"apiToken" gorm:"comment:api 的token;column:api_token;size:128;"`                                //token
	CmsType     *string `json:"cmsType" form:"cmsType" gorm:"comment:cms分类;column:cms_type;size:32;"`                                         //cms
	SiteType    *string `json:"siteType" form:"siteType" gorm:"comment:站点是做什么的，一些关键字，用于有可能的友链;column:site_type;size:128;" binding:"required"` //站点类型
	Status      *string `json:"status" form:"status" gorm:"index;default:1;comment:站点的状态;column:status;size:12;" binding:"required"`          //状态
	CategoryIds *string `json:"categoryIds" form:"categoryIds" gorm:"comment:站点的分类，以,进行划分;column:category_ids;size:64;"`                      //分类
	AuthorIds   *string `json:"authorIds" form:"authorIds" gorm:"comment:作者，站点的作者，用于文章更新，以 , 进行划分;column:author_ids;size:64;"`                //作者
	IP          *string `json:"iP" form:"iP" gorm:"comment:ip;column:ip;size:64;"`                                                            //ip
}

// TableName 基础信息 BasicInfo自定义表名 sites_basic_info
func (BasicInfo) TableName() string {
	return "sites_basic_info"
}
