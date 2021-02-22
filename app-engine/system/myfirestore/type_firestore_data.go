package myfirestore

type ConfigCollection struct {
	YoutubeLive YoutubeLiveDoc	`firestore:"youtube-live"`
}

type YoutubeLiveDoc struct {
	LiveChatId string `firestore:"live-chat-id"`
	SleepIntervalMilli int `firestore:"sleep-interval-milli"`
	NextPageToken string `firestore:"next-page-token"`
}

type DefaultRoomDoc struct {
	Seats []Seat `firestore:"seats"`
}

type Seat struct {
	SeatId int `firestore:"seat-id"`
	UserId string `firestore:"user-id"`
}

type NoSeatRoomDoc struct {
	Users []string `firestore:"users"`
}