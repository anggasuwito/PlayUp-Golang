package user

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

// InitUserRoute -> Init User Route
func InitUserRoute(app *model.App, mainRoute string) {
	userController := NewController(app.DB)

	user := app.Router.PathPrefix(mainRoute).Subrouter()
	admin := app.Router.PathPrefix(mainRoute).Subrouter()
	user.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	admin.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	// pake goroutines
	// untuk fungsi yg pake goroutine ngga ada nilai return (menggunakan pointer)
	user.HandleFunc("/login", userController.UserLogin).Methods("POST")
	user.HandleFunc("/register", userController.UserRegister).Methods("POST")
	user.HandleFunc("/update", userController.UserUpdate).Methods("PUT")
	user.HandleFunc("/detail/{id}", userController.GetUserByID).Methods("GET")
	
	// route login untuk oAuth(google/facebook)
	user.HandleFunc("/login/google", userController.UserLoginOAuth).Methods("POST")
	user.HandleFunc("/login/facebook", userController.UserLoginOAuth).Methods("POST")
	
	user.HandleFunc("/update/google", userController.UserUpdateGoogle).Methods("PUT")
	user.HandleFunc("/update/facebook", userController.UserUpdateFacebook).Methods("PUT")

	// get untuk admin
	admin.HandleFunc("/users", userController.GetUserList).Queries("page", "{page}", "limit", "{limit}").Methods("GET")
	admin.HandleFunc("/{id}", userController.GetUserByID).Methods("GET")

	// untuk upload dan tampilkan gambar
	user.HandleFunc("/post-image", userController.PostImage).Methods("POST")
	user.HandleFunc("/get-image/{id}", userController.GetImage).Methods("GET")

	// user.HandleFunc("/$2a$14$u9J55byw83Euka56VADxJewY8C3eT88PoLQJlN0DEf6rZPbs5j152/admin", userController.GetUserList).Methods("GET")

}
