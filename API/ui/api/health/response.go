package health

type DBQueryResponse struct {
	Texts []string `json:"texts"`
}

type DBCommandResponse struct {
	Result string `json:"result"`
}
