package vo

// 管理data 验证modle字段
type CreateCategoryRequest struct {
	Name string `json:"name" binding:"required"` // Name 不为空
}
