package saltgo

type Request struct {
	Client string    `json:"client"`
	Module string    `json:"fun"`
	Args  string	 `json:"arg"`
	Target []string  `json:"tgt"`
	Form   string	 `json:"expr_form"`
}

type Response struct {
	Return []struct{
		Jid string 		  `json:"jid"`
		Minions []string  `json:"minions"`
	}
}

type AuthResponse struct {
	Return []struct{
		Token string
	}
}

type RunnerRequest struct {
	Client string `json:"client"`
	Module string `json:"fun"`
	Jid   string `json:"jid"`
}

type RunnerResponse struct {
	Return []map[string]interface{}
}

