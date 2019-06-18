package gen

import (
	"fmt"
	"hash/fnv"
)

func unique(filename string) string {
	hash := fnv.New32()
	hash.Write([]byte(filename))
	return fmt.Sprintf("%x", hash.Sum32())
}
