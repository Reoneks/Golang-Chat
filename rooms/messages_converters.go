package rooms

func FromMessagesDto(messagesDto MessagesDto) Messages {
	return Messages(messagesDto)
}

func FromMessagesDtos(messagesDto []MessagesDto) (messages []Messages) {
	for _, dto := range messagesDto {
		messages = append(messages, FromMessagesDto(dto))
	}
	return
}

func ToMessagesDto(messages Messages) MessagesDto {
	return MessagesDto(messages)
}

func ToMessagesDtos(messages []Messages) (messagesDto []MessagesDto) {
	for _, dto := range messages {
		messagesDto = append(messagesDto, ToMessagesDto(dto))
	}
	return
}
