package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
    rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(min, max int64) int64{
    return min + rand.Int63n(max - min + 1)
}

func RandString(n int) string{
    var sb strings.Builder
    k := len(alphabet)
    
    for i := 0; i < n; i++{
        c := alphabet[rand.Intn(k)]
        sb.WriteByte(c)
    }
    return sb.String()
}

func RandUsername() string{
    return RandString(6)
}

func RandomEmail() string{
    return fmt.Sprintf("%s@%s.%s", RandString(6), RandString(4), RandString(3))
}
