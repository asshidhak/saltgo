package saltgo

import (
	"log"
	"testing"
)

var config = Cfg{
	"http://ip:8000",
	Auth{"pam","saltuser","password"},
}
func TestName(t *testing.T) {

	salt,_ := New(&config)
	if err := salt.Auth(); err != nil {
		panic(err)
	}
	//jid,_ := salt.RunCmdAsync("test.ping", "", []string{"Zbp10vzq9iu3rf2blm0r7Z","Zbp10vzq9iu3rf2blm0r7Z"})
	//salt.GetJob(jid)

	jid,_ := salt.StateAsync([]string{"Zbp10vzq9iu3rf2blm0r7Z"}, "cfg.nginxx")
	//salt.GetJob(jid)
	log.Println(jid)
	log.Println(salt.GetJob(jid))

	//salt.Abc()
}
