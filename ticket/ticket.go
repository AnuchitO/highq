package ticket

type Order struct {
	Age int `json:"age"`
}

func Price(o Order) float64 {
	if o.Age < 4 {
		return 0
	} else if o.Age < 13 {
		return 12
	} else if o.Age <= 65 {
		return 25
	} else {
		return 5
	}
}
