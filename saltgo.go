package saltgo

import (
	"net/http"
	"encoding/json"
	"bytes"
	"log"
	"io/ioutil"
)

type Client struct {
	token string
	cfg Cfg
}

type Cfg struct {
	Base string  `json:"base"`
	Auth Auth
}

type Auth struct{
	Eauth string `json:"eauth"`
	User string  `json:"username"`
	Passwd string  `json:"password"`
}

func New(cfg *Cfg) *Client{
	client := Client{}
	client.cfg = *cfg
	return &client
}



func (self *Client) Post(perfix string, data interface{}) (*http.Response, error) {

	urls := self.cfg.Base + perfix
	json_data, err := json.Marshal(data)
	if err != nil {
		panic(err)
		return nil, err
	}
	res,err := http.Post(urls, "application/json", bytes.NewBuffer(json_data))
	if err != nil {
		return nil, err
	}
	return res,err
}


func (salt *Client) Auth() (error) {

	res, err := salt.Post("/login", salt.cfg.Auth)
	if err!= nil {
		return err
	}
	resbyte, err := ioutil.ReadAll(res.Body)
	log.Println(string(resbyte))
	if err != nil {
		return err
	}
	return err

}

func (st *Client) RunCmdSync()  {

}
