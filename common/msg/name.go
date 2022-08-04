package msg

type NameMessage struct {
	Name string `json:"name"`
}

func (t *NameMessage) GetMessageType() uint32 {
	return NameMessageType
}

type NameResultMessage struct {
	Ip  string `json:"ip"`
	Msg string `json:"msg"`
}

func (t *NameResultMessage) GetMessageType() uint32 {
	return NameResultMessageType
}
