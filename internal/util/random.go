package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/rogsiel/quizzer/internal/model"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
    rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandInt(min, max int32) int32{
    return min + rand.Int31n(max - min + 1)
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

func RandQuestions(n int32) model.Question{
    qustions := make(model.Question)
    for q := 0; q < int(n); q++ {
	prompt := RandString(100)
	var options []string
	for i :=0; i < 3; i++ {
	    options = append(options, RandString(30))
	}
	qustions[prompt] = options
    }
    return qustions
}

func RandAnswers(n int) []int32 {
    var answers []int32
    for i := 0; i < n; i++ {
	answers = append(answers, RandInt(0,2))
    }
    return answers
}
