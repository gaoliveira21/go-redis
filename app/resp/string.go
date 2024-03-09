package resp

func NewRespString(s string) string {
	return "+" + s + "\r\n"
}
