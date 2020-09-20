package schedule

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

// InitScheduleRoute -> Init User Route
func InitScheduleRoute(app *model.App, mainRoute string) {
	scheduleController := NewController(app.DB)

	schedule := app.Router.PathPrefix(mainRoute).Subrouter()
	schedule.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	// pake goroutine
	// untuk fungsi yg pake goroutine ngga ada nilai return (menggunakan pointer)
	schedule.HandleFunc("/create", scheduleController.HandleCreateSchedule).Methods("POST")
	schedule.HandleFunc("/update", scheduleController.HandleUpdateSchedule).Methods("PUT")

	// ini untuk fetch jadwal pertandingan
	schedule.HandleFunc("/active/{id}", scheduleController.HandleFetchSchedule).Methods("GET")
	// ini untuk fetch history pertandingan
	schedule.HandleFunc("/inactive/{id}", scheduleController.HandleFetchHistory).Methods("GET")
	

	// ini untuk fetch detail jadwal pertandingan
	schedule.HandleFunc("/detail/{id}", scheduleController.HandleFetchDetailSchedule).Methods("GET")

	// ini untuk fetch semua schedule
	schedule.HandleFunc("/all", scheduleController.HandleFetchAllSchedule).Queries("page", "{page}", "limit", "{limit}").Methods("GET")


}
