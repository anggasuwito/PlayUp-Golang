package chat

type chat struct {
	ChatID         string `json:"chat_id"`
	ChatMatchID    string `json:"chat_match_id"`
	ChatSenderID   string `json:"chat_sender_id"`
	ChatReceiverID string `json:"chat_receiver_id"`
	ChatMessage    string `json:"chat_message"`


}
