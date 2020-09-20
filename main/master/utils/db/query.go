package db

const (
	SELECT_USER_TOTAL_COUNT          = `select count(user_id) from master_user`
	SELECT_USER_BY_ID                = `select * from master_user where user_id =?`
	REGISTER_USER                    = `INSERT INTO master_user(user_id,user_photo_profile,user_name,user_full_name,user_gender, user_email,user_password) VALUES (?,?,?,?,?,?,?)`
	SELECT_USER_BY_USERNAME_OR_EMAIL = `SELECT * FROM master_user WHERE user_name = ? or user_email = ?`
	SELECT_USER_BY_GOOGLE            = `SELECT 
	mu.user_id,
	mu.user_photo_profile,
	mu.user_name,
	mu.user_full_name,
	mu.user_gender,
	mu.user_email,
	mu.user_password,
	mu.user_google_account,
	mr.rank_id,
	mr.rank_user_match_count,
	mr.rank_user_grade_count
FROM
	master_user AS mu
		JOIN
	master_rank AS mr ON mu.user_id = mr.rank_user_id
WHERE
	mu.user_google_account = ?`
	SELECT_USER_BY_FACEBOOK = `SELECT 
	mu.user_id,
	mu.user_photo_profile,
	mu.user_name,
	mu.user_full_name,
	mu.user_gender,
	mu.user_email,
	mu.user_password,
	mu.user_facebook_account,
	mr.rank_id,
	mr.rank_user_match_count,
	mr.rank_user_grade_count
FROM
	master_user AS mu
		JOIN
	master_rank AS mr ON mu.user_id = mr.rank_user_id
WHERE
	mu.user_facebook_account = ?`
	SELECT_USER_BY_USERNAME = `
	SELECT 
		mu.user_id,
		mu.user_photo_profile,
		mu.user_name,
		mu.user_full_name,
		mu.user_gender,
		mu.user_email,
		mu.user_password,
		mr.rank_id,
		mr.rank_user_match_count,
		mr.rank_user_grade_count
	FROM
		master_user AS mu
			JOIN
		master_rank AS mr ON mu.user_id = mr.rank_user_id
	WHERE
		mu.user_name = ?`
	SELECT_USER_BY_USER_ID = `
	SELECT 
		mu.user_id,
		mu.user_photo_profile,
		mu.user_name,
		mu.user_full_name,
		mu.user_gender,
		mu.user_email,
		mu.user_password,
		mr.rank_id,
		mr.rank_user_match_count,
		mr.rank_user_grade_count,
		mu.user_facebook_account,
		mu.user_google_account
	FROM
		master_user AS mu
			JOIN
		master_rank AS mr ON mu.user_id = mr.rank_user_id
	WHERE
		mu.user_id = ?`

	UPDATE_USER_PHOTO_PROFILE = `update master_user set user_photo_profile =? where user_id =?`

	UPDATE_USER_PROFILE = `update master_user set user_name = ?,user_full_name =?,user_gender =?,user_email =? where user_id =?`

	SELECT_ALL_USER = `select * from master_user limit ?,?`
)
const (
	SELECT_SCHEDULE_BY_USER_OPPONENT_AND_STATUS = `select * from master_schedule where (schedule_user_id =? or schedule_opponent_id = ?) and schedule_status = ?`

	SELECT_SCHEDULE_BY_ID = `select * from master_schedule where schedule_id =?`

	SELECT_SCHEDULE_ALL = `select * from master_schedule limit ?,?`

	SELECT_SCHEDULE_TOTAL_DATA = `select count(schedule_id) from master_schedule`

	SELECT_SCHEDULE_BY_MATCH_ID = `select * from master_schedule where schedule_match_id = ?`

	INSERT_SCHEDULE = `INSERT master_schedule values (?,?,?,?,?,?,?,?,?,?)`

	SELECT_SCHEDULE_BY_MATCH_ID_AND_STATUS = `select * from master_schedule where schedule_match_id = ? and schedule_status=?`

	UPDATE_SCHEDULE = `UPDATE master_schedule SET schedule_status = 'I', schedule_result = ? WHERE schedule_id = ?`
)

const (
	INSERT_BADMINTON_SINGLE_PLAYER_LIST = `insert badminton_single_player_list values (?,?,?,?)`
)
const (
	UPDATE_RESULT = `UPDATE master_result set result_winner =?, result_loser=? where result_match_id=?`
	INSERT_RESULT = `insert into master_result values (?,?,?,?)`
)

const (
	SELECT_RANK_BY_USER_ID = `select rank_user_match_count,rank_user_grade_count from master_rank where rank_user_id =?`

	UPDATE_RANK = `update master_rank set rank_user_match_count =?, rank_user_grade_count = ? where rank_user_id =?`

	INSERT_RANK = `INSERT master_rank values (?,?,?,?)`

	SELECT_RANK_BY_ID_SPORT = `SELECT DISTINCT
    mu.user_id,
    mu.user_photo_profile,
    mu.user_name,
    mu.user_full_name,
    mr.rank_id,
    mr.rank_user_match_count,
    mr.rank_user_grade_count,
    mm.match_sport_id
FROM
    master_user AS mu
        JOIN
    master_rank AS mr ON mu.user_id = mr.rank_user_id
        JOIN
    master_schedule AS ms ON mu.user_id = ms.schedule_user_id
        JOIN
    master_match AS mm ON ms.schedule_match_id = mm.match_id
WHERE
    mm.match_sport_id = ? order by mr.rank_user_grade_count desc`
)
const (
	SELECT_ALL_CATEGORY     = `SELECT * FROM master_category WHERE category_status='1'`
	SELECT_CATEGORY_BY_NAME = `SELECT * FROM master_category WHERE category_name=? AND category_status='1'`
	SELECT_CATEGORY_BY_ID   = `SELECT * FROM master_category WHERE category_id=? AND category_status='1'`
	INSERT_CATEGORY         = `INSERT INTO master_category(category_id,category_name, category_img) VALUES (?,?,?)`

	UPDATE_CATEGORY = `UPDATE master_category SET category_name=?, category_img=? WHERE category_id=?`
	DELETE_CATEGORY = `UPDATE master_category SET category_status='0' WHERE category_id=?`
)
const (
	SELECT_ALL_FEATURE     = `SELECT * FROM master_features WHERE feature_status='1'`
	SELECT_FEATURE_BY_NAME = `SELECT * FROM master_features WHERE feature_list=? AND feature_status='1'`
	SELECT_FEATURE_BY_ID   = `SELECT * FROM master_features WHERE feature_id=? AND feature_status='1'`
	INSERT_FEATURE         = `INSERT INTO master_features(feature_id,feature_list,feature_description,feature_img) VALUES (?,?,?,?)`
	UPDATE_FEATURE         = `UPDATE master_features SET feature_list=?, feature_description=?, feature_img=? WHERE feature_id=?`
	DELETE_FEATURE         = `UPDATE master_features SET feature_status='0' WHERE feature_id=?`
)
const (
	INSERT_MATCH = `insert master_match values (?,?)`
)
