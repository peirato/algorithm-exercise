package hello 

import ("testing"
	"utils"
"strconv")

func TestHelloWorld(t *testing.T) {
	i := utils.RandomInt(100)
    t.Log("hello world" + strconv.Itoa(i))
}