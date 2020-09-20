package chat

import (
	"playUp/main/master/middleware"
	"playUp/main/master/model"
)

// InitChatRoute -> Init User Route
func InitChatRoute(app *model.App, mainRoute string) {
	chatController := NewController(app.DB)

	chat := app.Router.PathPrefix(mainRoute).Subrouter()
	chat.Use(middleware.LoggerMiddleware,middleware.SetJSONMiddleware)
	
	chat.HandleFunc("", chatController.HandleSendChat).Methods("POST")
	chat.HandleFunc("/{id}", chatController.HandleGetChat).Methods("GET")

}
