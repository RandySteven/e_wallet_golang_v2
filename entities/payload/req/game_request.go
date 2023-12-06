package req

type ChooseReward struct {
	GameID uint
	BoxID  uint `json:"box_id" binding:"required"`
}
