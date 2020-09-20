package features

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

func InitFeatureRoute(app *model.App, mainRoute string) {
	featureController := NewController(app.DB)
	feature := app.Router.PathPrefix(mainRoute).Subrouter()
	
	feature.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)

	// feature routes untuk admin
	feature.HandleFunc("", featureController.GetFeatureList).Methods("GET")
	feature.HandleFunc("/{id}", featureController.FindFeature).Methods("GET")
	feature.HandleFunc("/create", featureController.PostFeature).Methods("POST")
	feature.HandleFunc("/update", featureController.PutFeature).Methods("PUT")
	feature.HandleFunc("/delete/{id}", featureController.DeleteFeature).Methods("DELETE")

}
