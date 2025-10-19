package dto

type ResultDetail struct {
	PlayerID  string `json:"player_id" binding:"required,uuid4"`
	GoalTime  string `json:"goal_time" binding:"required,goal_time_format"`
	IsPenalty bool   `json:"is_penalty"`
}

type CreateResult struct {
	ScheduleID string         `json:"schedule_id" binding:"required,uuid4"`
	ScoreHome  int64          `json:"score_home" binding:"required,min=0"`
	ScoreAway  int64          `json:"score_away" binding:"required,min=0"`
	Details    []ResultDetail `json:"details" binding:"required,dive"`
}
