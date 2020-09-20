package chat

import (
	"database/sql"
	"errors"
	"log"
	"playUp/main/master/utils/generator"
)

type ChatRepo struct {
	db *sql.DB
}

type ChatRepository interface {
	handleGetChat(chatList *[]*chat, matchID string) error

	handleSendChat(newChat *chat) error
}

func NewChatRepo(db *sql.DB) ChatRepository {
	return ChatRepo{db: db}
}

func (d ChatRepo) handleGetChat(chatList *[]*chat, matchID string) error {
	tx, err := d.db.Begin()

	if err != nil {
		log.Println("ERROR 1 -> %v", err)
		return err

	}

	rows, err := tx.Query(`select * from master_chat where chat_match_id = ?`, matchID)

	if err != nil {
		log.Println("ERROR 1")
		log.Println(err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		var each chat
		err := rows.Scan(&each.ChatID, &each.ChatMatchID, &each.ChatSenderID, &each.ChatReceiverID, &each.ChatMessage)
		if err != nil {
			log.Printf("ERROR 2 -> %v", err)
			tx.Rollback()
			return errors.New("Failed in here")
		}
		*chatList = append(*chatList, &each)
	}
	tx.Commit()
	return nil
}
func (d ChatRepo) handleSendChat(newChat *chat) error {

	tx, err := d.db.Begin()
	
	chatID := generator.GenerateUUID()
	if err != nil {
		log.Println("ERROR 1 -> %v", err)
		return err

	}

	stmt, _ := tx.Prepare(`insert master_chat values(?,?,?,?,?)`)

	log.Println("-> sampe sini 2 di send chat")
	log.Println(newChat.ChatMatchID)
	_, err = stmt.Exec(chatID, newChat.ChatMatchID, newChat.ChatSenderID, newChat.ChatReceiverID, newChat.ChatMessage)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
