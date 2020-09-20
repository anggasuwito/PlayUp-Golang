package rank

type rankList struct {
	RankUserID           string `json:"rank_user_id"`
	RankUserPhotoProfile string `json:"rank_user_photo_profile"`
	RankUserUserName     string `json:"rank_user_name"`
	RankUserUserFullName string `json:"rank_user_full_name"`
	RankUserRankID       string `json:"rank_id"`
	RankUserMatchCount   string `json:"rank_match_count"`
	RankUserGradeCount   string `json:"rank_grade_count"`
	RankUserMatchSportID string `json:"rank_match_sport_id"`
}
