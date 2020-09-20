package category

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

func InitCategoryRoute(app *model.App, mainRoute string) {
	categoryController := NewController(app.DB)
	category := app.Router.PathPrefix(mainRoute).Subrouter()
	category.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	
	// category routes ntuk admin
	category.HandleFunc("", categoryController.GetCategoryList).Methods("GET")
	category.HandleFunc("/{id}", categoryController.FindCategory).Methods("GET")
	category.HandleFunc("/create", categoryController.PostCategory).Methods("POST")
	category.HandleFunc("/update", categoryController.PutCategory).Methods("PUT")
	category.HandleFunc("/delete/{id}", categoryController.DeleteCategory).Methods("DELETE")

}
