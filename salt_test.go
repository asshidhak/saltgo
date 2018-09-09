package saltgo

import (
	"testing"
	"log"
)

var config = Cfg{
	"http://127.0.0.1:8080",
	Auth{"pam","xxx","xxx"},
}
func TestName(t *testing.T) {

	salt :=New(&config)
	log.Println(salt.Auth())
	//salt.Abc()
}
