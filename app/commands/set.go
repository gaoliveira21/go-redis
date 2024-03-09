package commands

func Set(store map[string]string, key string, value string) {
	store[key] = value
}
