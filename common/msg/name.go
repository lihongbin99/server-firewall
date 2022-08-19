package msg

type NameMessage struct {
	Name string `json:"name"`
}

func (t *NameMessage) GetMessageType() uint32 {
	return NameMessageType
}

type NameResultMessageDetails struct {
	Name string `json:"name"`
	Ip   string `json:"ip"`
	Msg  string `json:"msg"`
}
type NameResultMessage struct {
	Details []NameResultMessageDetails `json:"details"`
}

func (t *NameResultMessage) GetMessageType() uint32 {
	return NameResultMessageType
}
