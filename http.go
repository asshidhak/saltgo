package saltgo

type Request struct {
	Client string    `json:"client"`
	Module string    `json:"fun"`
	Args  string	 `json:"args"`
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
	Args   string `json:"args"`

}
