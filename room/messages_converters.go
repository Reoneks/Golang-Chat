package room

func FromMessagesDto(messagesDto MessagesDto) Message {
	return Message(messagesDto)
}

func FromMessagesDtos(messagesDto []MessagesDto) (messages []Message) {
	for _, dto := range messagesDto {
		messages = append(messages, FromMessagesDto(dto))
	}
	return
}

func ToMessagesDto(messages Message) MessagesDto {
	return MessagesDto(messages)
}

func ToMessagesDtos(messages []Message) (messagesDto []MessagesDto) {
	for _, dto := range messages {
		messagesDto = append(messagesDto, ToMessagesDto(dto))
	}
	return
}
