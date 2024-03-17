package resp

type RespString struct {
	Len   int
	Value string
}

func NewRespString(l int, s string) *RespString {
	return &RespString{
		Len:   l,
		Value: s,
	}
}

func (s *RespString) Encode() string {
	return "+" + s.Value + "\r\n"
}
