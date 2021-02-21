package counter

const max = 4294967296

var c = make(map[string]uint)

// Init conter
func Init(key string) {
	c[key] = 0
}

// Up increate Key conter and return
func Up(key string) uint {
	if c[key] == max {
		c[key] = 0
	}
	c[key]++
	return c[key]
}

// Erease Key
func Erease(key string) {
	delete(c, key)
}
