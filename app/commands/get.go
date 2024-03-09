package commands

func Get(store map[string]string, key string) (string, bool) {
	v, f := store[key]

	return v, f
}
