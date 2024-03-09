package commands

import "strings"

func Echo(args []string) string {
	str := strings.Join(args, "")

	return str
}
