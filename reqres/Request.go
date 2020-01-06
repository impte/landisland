package reqres

type LoginRequest struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}