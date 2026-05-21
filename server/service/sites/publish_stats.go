
package sites

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
    sitesReq "github.com/flipped-aurora/gin-vue-admin/server/model/sites/request"
)

type PublishStatsService struct {}
// CreatePublishStats 创建网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService) CreatePublishStats(ctx context.Context, publishStats *sites.PublishStats) (err error) {
	err = global.GVA_DB.Create(publishStats).Error
	return err
}

// DeletePublishStats 删除网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService)DeletePublishStats(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&sites.PublishStats{},"id = ?",ID).Error
	return err
}

// DeletePublishStatsByIds 批量删除网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService)DeletePublishStatsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]sites.PublishStats{},"id in ?",IDs).Error
	return err
}

// UpdatePublishStats 更新网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService)UpdatePublishStats(ctx context.Context, publishStats sites.PublishStats) (err error) {
	err = global.GVA_DB.Model(&sites.PublishStats{}).Where("id = ?",publishStats.ID).Updates(&publishStats).Error
	return err
}

// GetPublishStats 根据ID获取网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService)GetPublishStats(ctx context.Context, ID string) (publishStats sites.PublishStats, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&publishStats).Error
	return
}
// GetPublishStatsInfoList 分页获取网站更新数据记录
// Author [yourname](https://github.com/yourname)
func (publishStatsService *PublishStatsService)GetPublishStatsInfoList(ctx context.Context, info sitesReq.PublishStatsSearch) (list []sites.PublishStats, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sites.PublishStats{})
    var publishStatss []sites.PublishStats
    // 如果有条件搜索 下方会自动创建搜索语句
    if len(info.CreatedAtRange) == 2 {
     db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
    }
    
	if info.StartTodayCount != nil && info.EndTodayCount != nil {
		db = db.Where("today_count BETWEEN ? AND ? ", *info.StartTodayCount, *info.EndTodayCount)
	}
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&publishStatss).Error
	return  publishStatss, total, err
}
func (publishStatsService *PublishStatsService)GetPublishStatsDataSource(ctx context.Context) (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   siteId := make([]map[string]any, 0)
	   
       
       global.GVA_DB.Table("sites_basic_info").Where("deleted_at IS NULL").Select("domain as label,id as value").Scan(&siteId)
	   res["siteId"] = siteId
	return
}
func (publishStatsService *PublishStatsService)GetPublishStatsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
