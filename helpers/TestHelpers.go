package helpers

import (
	"encoding/base64"
	"fmt"
	"math/rand"
)

//CreateKeys - new create keys
func CreateKeys(n int, seed int64, len int) []string {
	rand := rand.New(rand.NewSource(seed))
	res := make([]string, n)
	b := make([]byte, len)
	for i := range res {
		rand.Read(b)
		res[i] = base64.RawStdEncoding.EncodeToString(b)
	}
	return res
}

//CreateStringSlice create string slice with the input len  with each diffrent elements
func CreateStringSlice(prefix string, firstIndex int, len int, suffix string) []string {
	res := make([]string, len)
	for i := 0; i < len; i++ {
		res[i] = fmt.Sprintf("%v%v%v", prefix, i+firstIndex, suffix)
	}
	return res
}
