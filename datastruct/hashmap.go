package datastruct

// import (
//     "fmt"
// )

// See: http://www.cse.yorku.ca/~oz/hash.html
func sdbmHash(text string) uint64 {
	textBytes := []byte(text)
	var hash uint64 = 0
	for _, c := range textBytes {
		hash = uint64(c) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}


type HashMap struct {
	items [1024]*LinkedList
}

func (hm *HashMap) Add(key string, value int) {
	return
}

func (hm *HashMap) Get(key string) int {
	return 3
}