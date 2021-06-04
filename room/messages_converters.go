package room

func FromMessagesDto(messagesDto MessagesDto) Message {
	return Message{
		Id:        messagesDto.Id,
		Text:      messagesDto.Text,
		Status:    NewStatusType(messagesDto.Status),
		RoomID:    messagesDto.RoomID,
		CreatedBy: messagesDto.CreatedBy,
		CreatedAt: messagesDto.CreatedAt,
		UpdatedAt: messagesDto.UpdatedAt,
	}
}

func FromMessagesDtos(messagesDto []MessagesDto) (messages []Message) {
	for _, dto := range messagesDto {
		messages = append(messages, FromMessagesDto(dto))
	}
	return
}

func ToMessagesDto(messages Message) MessagesDto {
	return MessagesDto{
		Id:        messages.Id,
		Text:      messages.Text,
		Status:    messages.Status.ToInt64(),
		RoomID:    messages.RoomID,
		CreatedBy: messages.CreatedBy,
		CreatedAt: messages.CreatedAt,
		UpdatedAt: messages.UpdatedAt,
	}
}

func ToMessagesDtos(messages []Message) (messagesDto []MessagesDto) {
	for _, dto := range messages {
		messagesDto = append(messagesDto, ToMessagesDto(dto))
	}
	return
}
