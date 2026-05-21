
package sites

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
    sitesReq "github.com/flipped-aurora/gin-vue-admin/server/model/sites/request"
)

type BasicInfoService struct {}
// CreateBasicInfo 创建基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) CreateBasicInfo(ctx context.Context, basicInfo *sites.BasicInfo) (err error) {
	err = global.GVA_DB.Create(basicInfo).Error
	return err
}

// DeleteBasicInfo 删除基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService)DeleteBasicInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&sites.BasicInfo{},"id = ?",ID).Error
	return err
}

// DeleteBasicInfoByIds 批量删除基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService)DeleteBasicInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]sites.BasicInfo{},"id in ?",IDs).Error
	return err
}

// UpdateBasicInfo 更新基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService)UpdateBasicInfo(ctx context.Context, basicInfo sites.BasicInfo) (err error) {
	err = global.GVA_DB.Model(&sites.BasicInfo{}).Where("id = ?",basicInfo.ID).Updates(&basicInfo).Error
	return err
}

// GetBasicInfo 根据ID获取基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService)GetBasicInfo(ctx context.Context, ID string) (basicInfo sites.BasicInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&basicInfo).Error
	return
}
// GetBasicInfoInfoList 分页获取基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService)GetBasicInfoInfoList(ctx context.Context, info sitesReq.BasicInfoSearch) (list []sites.BasicInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sites.BasicInfo{})
    var basicInfos []sites.BasicInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
    if info.Domain != nil && *info.Domain != "" {
        db = db.Where("domain LIKE ?", "%"+ *info.Domain+"%")
    }
    if info.CmsType != nil && *info.CmsType != "" {
        db = db.Where("cms_type = ?", *info.CmsType)
    }
    if info.Status != nil && *info.Status != "" {
        db = db.Where("status = ?", *info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&basicInfos).Error
	return  basicInfos, total, err
}
func (basicInfoService *BasicInfoService)GetBasicInfoPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
