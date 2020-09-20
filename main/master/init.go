package master

import (
	"database/sql"
	"playUp/main/master/domains/category"
	"playUp/main/master/domains/chat"
	"playUp/main/master/domains/features"
	"playUp/main/master/domains/rank"

	"playUp/main/master/domains/badmintonSingle"
	"playUp/main/master/domains/schedule"
	"playUp/main/master/domains/user"
	"playUp/main/master/model"

	"github.com/gorilla/mux"
)

const (
	MAIN_USER_ROUTE = "/user"
	MAIN_PLAY_ROUTE = "/play"
	MAIN_SCHEDULE_ROUTE = "/schedule"
	MAIN_CATEGORY_ROUTE = "/category"
	MAIN_FEATURE_ROUTE = "/feature"
	MAIN_RANK_ROUTE = "/rank"
	MAIN_CHAT_ROUTE = "/chat"

)

func Init(db *sql.DB, r *mux.Router) {
	newAppConfig := &model.App{
		DB:     db,
		Router: r,
	}
	user.InitUserRoute(newAppConfig, MAIN_USER_ROUTE)
	badmintonSingle.InitBSRoute(newAppConfig, MAIN_PLAY_ROUTE)
	schedule.InitScheduleRoute(newAppConfig,MAIN_SCHEDULE_ROUTE)
	category.InitCategoryRoute(newAppConfig,MAIN_CATEGORY_ROUTE)
	features.InitFeatureRoute(newAppConfig,MAIN_FEATURE_ROUTE)
	rank.InitRankRoute(newAppConfig,MAIN_RANK_ROUTE)
	chat.InitChatRoute(newAppConfig,MAIN_CHAT_ROUTE)
}
