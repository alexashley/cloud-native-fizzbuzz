package domain

type Mod3OracleRequest struct {
	Value int `json:"value"`
}

type Mod3OracleResponse struct {
	IsDivisibleByThree bool `json:"divisibleByThree"`
}
