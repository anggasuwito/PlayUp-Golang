package badmintonSingle

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

func InitBSRoute(app *model.App, mainRoute string) {

	BadmintonSController := NewController(app.DB)

	bs := app.Router.PathPrefix(mainRoute).Subrouter()

	bs.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	
	bs.HandleFunc("/badmintonS/find", BadmintonSController.HandleFindOpponent).Methods("POST")

	bs.HandleFunc("/badmintonS/cancel/{id}", BadmintonSController.HandleCancelMatch).Methods("POST")

	bs.HandleFunc("/badmintonS/reset", BadmintonSController.HandleResetRoom).Methods("POST")

}
