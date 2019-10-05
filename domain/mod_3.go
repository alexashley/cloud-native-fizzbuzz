package domain

type Mod3Query struct {
	Value int `json:"value"`
}

type Mod3Result struct {
	IsDivisibleByThree bool `json:"divisibleByThree"`
}
