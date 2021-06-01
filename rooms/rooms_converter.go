package rooms

func FromRoomsDto(roomDto RoomsDto) Rooms {
	return Rooms(roomDto)
}

func FromRoomsDtos(roomsDtos []RoomsDto) (rooms []Rooms) {
	for _, dto := range roomsDtos {
		rooms = append(rooms, FromRoomsDto(dto))
	}
	return
}

func ToRoomsDto(room Rooms) RoomsDto {
	return RoomsDto(room)
}

func ToRoomsDtos(rooms []Rooms) (RoomsDtos []RoomsDto) {
	for _, dto := range rooms {
		RoomsDtos = append(RoomsDtos, ToRoomsDto(dto))
	}
	return
}
