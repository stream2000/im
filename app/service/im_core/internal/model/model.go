package model

type Message struct {
	SenderId   int64  `json:"sender_id"`
	ReceiverId int64  `json:"receiver_id"`
	SendSeq    int32  `json:"send_seq"`
	MsgType    int32  `json:"msg_type"`
	Content    string `json:"content"`
}

type OnlineRecord struct {
	DeviceId   string `json:"device_id"`
	DeviceType int    `json:"device_type"`
	Server     string `json:"server"`
}

type TimelineItem struct {
	Guid         int64  `json:"guid"`
	RedisId      int64  `json:"redis_id"`
	ServerSeq    int64  `json:"server_seq"`
	SenderId     int64  `json:"sender_id"`
	ReceiverId   int64  `json:"receiver_id"`
	SendSeq      int32  `json:"send_seq"`
	MsgType      int32  `json:"msg_type"`
	Content      string `json:"content"`
	BusinessType int    `json:"business_type"`
}

func TimelineItemMap(msg Message) (item TimelineItem) {
	item.SenderId = msg.SenderId
	item.ReceiverId = msg.ReceiverId
	item.SendSeq = msg.SendSeq
	item.MsgType = msg.MsgType
	item.Content = msg.Content
	return
}

const (
	C2CMessage = iota + 1
	C2GMessage
)
