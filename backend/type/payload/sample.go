package payload

type Foo struct {
	Bar *int `json:"bar" validate:"required"`
}

type TargetUrl struct {
	Url *string `json:"url" validate:"required"`
}

type PlusBody struct {
	No1 *int `json:"no1" validate:"required"`
	No2 *int `json:"no2" validate:"gte=2"`
}

type CalculateResult struct {
	Result *int `json:"result"`
}
