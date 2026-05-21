package task

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/model/sites"
	"gorm.io/gorm"
)

// ResetPublishStats 每日 00:00 重置 sites_publish_stats 全表：
// yesterday_count 置为当前 today_count，today_count 清零。
// 使用单条 UPDATE 保证原子性，避免读-写之间被新发布写入挤掉数据。
func ResetPublishStats(db *gorm.DB) error {
	if db == nil {
		return errors.New("db Cannot be empty")
	}
	return db.Model(&sites.PublishStats{}).
		Where("1 = 1").
		Updates(map[string]interface{}{
			"yesterday_count": gorm.Expr("today_count"),
			"today_count":     0,
		}).Error
}
