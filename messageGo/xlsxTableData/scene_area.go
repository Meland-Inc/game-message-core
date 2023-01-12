package xlsxTable

import "time"

type SceneAreaRow struct {
	Id        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	SceneType string    `json:"sceneType"`
	CreatedAt time.Time `json:"createdAt"` // 过期判断条件
}
