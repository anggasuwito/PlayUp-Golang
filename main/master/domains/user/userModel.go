package user

// User -> user struct
type User struct {
	ID           string `json:"id"`
	PhotoProfile string `json:"photo"`
	Username     string `json:"username"`
	UserFullName string `json:"user_full_name"`
	UserGender   string `json:"gender"`
	UserEmail    string `json:"email"`
	Password     string `json:"password"`
	rank
	Created         string `json:"created"`
	Updated         string `json:"updated"`
	FacebookAccount string `json:"facebook_account"`
	GoogleAccount   string `json:"google_account"`
}
type rank struct {
	RankID             string `json:"rank_id"`
	RankUserMatchCount string `json:"rank_user_match_count"`
	RankUserMatchGrade string `json:"rank_user_grade_count"`
}
