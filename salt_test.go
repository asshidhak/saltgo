package saltgo

import (
	"log"
	"testing"
)

var config = Cfg{
	"http://192.168.1.100:8000",
	Auth{"pam","saltapi","123123"},
}
func TestName(t *testing.T) {

	salt :=New(&config)
	if err := salt.Auth(); err != nil {
		panic(err)
	}
	jid,_ := salt.RunCmdAsync("test.ping", "", []string{"Zbp10vzq9iu3rf2blm0r7Z"})
	log.Println(jid)
	//salt.Abc()
}
