package param

type PostDanmakuParam struct {
	Token    string `json:"token"`
	Av       int64  `json:"video_id" binding:"required"`
	Value    string `json:"value"`
	Color    string `json:"color" binding:"required"`
	Type     string `json:"type"`
	Location int64  `json:"location" binding:"required"`
}
