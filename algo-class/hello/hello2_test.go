package hello

import (
	"strconv"
	"testing"
	"utils"
)

func TestHello2World(t *testing.T) {
	i := utils.RandomInt(100)
	t.Log("hello world" + strconv.Itoa(i))
}
