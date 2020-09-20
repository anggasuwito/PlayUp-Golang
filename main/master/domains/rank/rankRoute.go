package rank

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

// InitRankRoute -> Init Rank Route
func InitRankRoute(app *model.App, mainRoute string) {
	rankController := NewController(app.DB)

	rank := app.Router.PathPrefix(mainRoute).Subrouter()
	rank.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	
	// untuk tampilkan rank berdasarkan sport id
	// pake goroutine
	// untuk fungsi yg pake goroutine ngga ada nilai return (menggunakan pointer)
	rank.HandleFunc("/{id}", rankController.HandleGetRankBySportID).Methods("GET")

}
