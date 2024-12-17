package models


// CanvasData 用来存储画板数据
type CanvasData struct {
	ID         uint   `json:"id" gorm:"primaryKey"` // 主键
	UserID     uint   `json:"user_id"`              // 外键：与用户关联
	CanvasData string `json:"canvas_data"`          // 画板的 JSON 数据
	CreatedAt  string `json:"created_at"`           // 创建时间
	UpdatedAt  string `json:"updated_at"`           // 更新时间
}
