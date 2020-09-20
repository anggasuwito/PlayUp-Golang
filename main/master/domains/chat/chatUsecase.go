package chat

import (
	"database/sql"
	"log"
)

type ChatUsecase struct {
	db       *sql.DB
	ChatRepo ChatRepository
}

type ChatUsecaseInterface interface {
	handleGetChat(chatList *[]*chat,matchID string) error
	handleSendChat(newChat *chat) error
}

func NewChatUsecase(db *sql.DB) ChatUsecaseInterface {
	return ChatUsecase{db, NewChatRepo(db)}
}
func (r ChatUsecase) handleGetChat(chatList *[]*chat,matchID string) error {
	err := r.ChatRepo.handleGetChat(chatList,matchID) 
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (r ChatUsecase) handleSendChat(newChat *chat) error {
	err := r.ChatRepo.handleSendChat(newChat) 
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}
