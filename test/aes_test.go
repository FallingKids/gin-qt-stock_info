package test

import (
	"fmt"
	"testing"

	"github.com/gin-qt-business/app/utils"
)

func TestAes(t *testing.T) {
	text := "123456"
	key := "jufeng6668987364"
	iv := "jufeng1234567890"
	res, _ := utils.Encrypt([]byte(text), []byte(key), []byte(iv))
	fmt.Print(res)
}
