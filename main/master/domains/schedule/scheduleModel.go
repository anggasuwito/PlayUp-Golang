package schedule

type schedule struct {
	ScheduleID            string `json:"schedule_id"`
	ScheduleMatchID       string `json:"schedule_match_id"`
	ScheduleUserID        string `json:"schedule_user_id"`
	ScheduleUserName      string `json:"schedule_user_name"`
	SceduleUserPhoto      string `json:"schedule_user_photo"`
	SchedeleLocation      string `json:"schedule_location"`
	ScheduleTime          string `json:"schedule_time"`
	ScheduleStatus        string `json:"schedule_status"`
	ScheduleResult        string `json:"schedule_result"`
	ScheduleOpponent      string `json:"schedule_opponent"`
	ScheduleOpponentID    string `json:"schedule_opponent_id"`
	ScheduleOpponentPhoto string `json:"schedule_opponent_photo"`
}
type resultMatch struct {
	Winner   string `json:"winner"`
	WinnerID string `json:"winner_id"`
	Loser    string `json:"loser"`
	LoserID  string `json:"loser_id"`
}
type rank struct {
	rankMatchCount int
	rankGradeCount int
}
