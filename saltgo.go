package saltgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
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

func New(cfg *Cfg) (*Client, error) {
	client := Client{}
	client.cfg = *cfg
	if err := client.Auth(); err != nil {
		return nil,err
	}
	return &client, nil
}

func (self *Client) Post(perfix string, data interface{}) (*http.Response, error) {

	urls := self.cfg.Base + perfix
	json_data, err := json.Marshal(data)
	if err != nil {
		panic(err)
		return nil, err
	}
	log.Println("Post data: ",string(json_data))
	r := &http.Client{}
	req, _ := http.NewRequest("POST",urls, bytes.NewBuffer(json_data))
	req.Header.Set("X-Auth-Token",self.token)
	req.Header.Set("Content-Type","application/json")
	res,err := r.Do(req)
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
	if res.StatusCode != 200 {
		return errors.New("auth failed.")
	}
	resbyte, err := ioutil.ReadAll(res.Body)
	resault := AuthResponse{}
	if err := json.Unmarshal(resbyte,&resault); err !=nil {
		return err
	}
	new_token := resault.Return[0].Token
	salt.token = new_token
	return nil

}


func (salt *Client) RunCmdAsync(fun string, args string, tgt []string) (string, error) {

	saltRes := Response{}
	client := "local_async"
	run := Request{client, fun,args, tgt, "list"}
	res, err := salt.Post("/", run)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if res.StatusCode != 200 {
		return "", errors.New("failed to run job.")
	}
	resbyte,err := ioutil.ReadAll(res.Body)
	if err := json.Unmarshal(resbyte, &saltRes); err != nil {
		return "", err
	}
	return saltRes.Return[0].Jid, nil
}


//salt "Zbp10vzq9iu3rf2blm0r7Z" state.sls cfg.nginx
func (salt *Client) StateAsync(tgt []string, state string) (string, error) {
	saltRes := Response{}
	client := "local_async"
	fun := "state.sls"
	job := Request{client, fun, state, tgt, "list"}
	res,err := salt.Post("/", job)
	if err != nil {
		return "", err
	}
	resbyte,err:= ioutil.ReadAll(res.Body)

	if err!= nil {
		return "",err
	}
	if err := json.Unmarshal(resbyte, &saltRes); err != nil {
		return "", nil
	}
	return  saltRes.Return[0].Jid, err
}


func (salt *Client) GetJob(jid string) ([]map[string]interface{}, error) {
	client := "runner"
	job := RunnerRequest{client, "jobs.lookup_jid", jid}
	res, err := salt.Post("/",job)
	if err != nil {
		return nil, err
	}
	bytes,_ := ioutil.ReadAll(res.Body)
	result := RunnerResponse{}
	if err := json.Unmarshal(bytes,&result); err != nil {
		return nil,err
	}
	return result.Return, nil
}