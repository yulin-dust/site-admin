package sites

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
	sitesReq "github.com/flipped-aurora/gin-vue-admin/server/model/sites/request"
	sitesRes "github.com/flipped-aurora/gin-vue-admin/server/model/sites/response"
	"gorm.io/gorm"
)

// 站点正常状态标识，对应 sites_basic_info.status 字段
const basicInfoStatusNormal = "1"

// 新建站点时 PublishStats.PlanCount 的默认值
const defaultPlanCount int64 = 10

type BasicInfoService struct{}

// CreateBasicInfo 创建基础信息记录，同步初始化对应的 PublishStats 行
// （PlanCount 默认 defaultPlanCount，其余计数为 0）。
// 整个过程在事务中执行：PublishStats 初始化失败会回滚 BasicInfo 创建，
// 避免出现"站点存在但没有统计行"的中间态。
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) CreateBasicInfo(ctx context.Context, basicInfo *sites.BasicInfo) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(basicInfo).Error; err != nil {
			return err
		}
		return tx.Create(newInitialPublishStats(basicInfo.ID)).Error
	})
}

// newInitialPublishStats 为新建站点构造初始的发布统计行
func newInitialPublishStats(siteUintID uint) *sites.PublishStats {
	siteID := int32(siteUintID)
	var (
		today     int64 = 0
		plan            = defaultPlanCount
		total     int64 = 0
		yesterday int64 = 0
	)
	return &sites.PublishStats{
		SiteId:         &siteID,
		TodayCount:     &today,
		PlanCount:      &plan,
		TotalCount:     &total,
		YesterdayCount: &yesterday,
	}
}

// DeleteBasicInfo 删除基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) DeleteBasicInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&sites.BasicInfo{}, "id = ?", ID).Error
	return err
}

// DeleteBasicInfoByIds 批量删除基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) DeleteBasicInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]sites.BasicInfo{}, "id in ?", IDs).Error
	return err
}

// UpdateBasicInfo 更新基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) UpdateBasicInfo(ctx context.Context, basicInfo sites.BasicInfo) (err error) {
	err = global.GVA_DB.Model(&sites.BasicInfo{}).Where("id = ?", basicInfo.ID).Updates(&basicInfo).Error
	return err
}

// GetBasicInfo 根据ID获取基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) GetBasicInfo(ctx context.Context, ID string) (basicInfo sites.BasicInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&basicInfo).Error
	return
}

// GetBasicInfoInfoList 分页获取基础信息记录
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) GetBasicInfoInfoList(ctx context.Context, info sitesReq.BasicInfoSearch) (list []sites.BasicInfo, total int64, err error) {
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
		db = db.Where("domain LIKE ?", "%"+*info.Domain+"%")
	}
	if info.CmsType != nil && *info.CmsType != "" {
		db = db.Where("cms_type = ?", *info.CmsType)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	db = db.Order("id desc")

	err = db.Find(&basicInfos).Error
	return basicInfos, total, err
}
func (basicInfoService *BasicInfoService) GetBasicInfoPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// GetAllSites 获取正常状态的网站信息（不含 GVA_MODEL / IP 字段）。
// onlyPending = true 时，仅返回今日发布数还没达到计划数的站点（today_count < plan_count），
// 适合爬虫调度场景："还需要继续发的站点有哪些"。
// Author [yourname](https://github.com/yourname)
func (basicInfoService *BasicInfoService) GetAllSites(ctx context.Context, onlyPending bool) (list []sitesRes.BasicInfoBrief, err error) {
	db := global.GVA_DB.
		Model(&sites.BasicInfo{}).
		Where("sites_basic_info.status = ?", basicInfoStatusNormal)

	if onlyPending {
		// JOIN 发布统计表，过滤当日还未达到计划量的站点
		db = db.
			Joins("JOIN sites_publish_stats ps ON ps.site_id = sites_basic_info.id AND ps.deleted_at IS NULL").
			Where("COALESCE(ps.today_count, 0) < COALESCE(ps.plan_count, 0)")
	}

	err = db.
		Select("sites_basic_info.domain",
			"sites_basic_info.url",
			"sites_basic_info.api_token",
			"sites_basic_info.cms_type",
			"sites_basic_info.site_type",
			"sites_basic_info.status",
			"sites_basic_info.category_ids",
			"sites_basic_info.author_ids").
		Find(&list).Error
	return
}
