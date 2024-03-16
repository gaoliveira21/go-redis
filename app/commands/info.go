package commands

import "errors"

func Info(t string) (map[string]string, error) {
	switch t {
	case "replication":
		r := make(map[string]string)

		r["role"] = "master"

		return r, nil
	default:
		return nil, errors.New("invalid argument received")
	}
}
