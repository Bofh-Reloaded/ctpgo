package common

type Exchange struct {
	Name    string
	Router  Address
	Factory Address
	Fee     int
	Etype   int
}

func NewExchange(
	name string,
	router string,
	factory string,
	fee int,
	etype int,
) *Exchange {
	ex := &Exchange{
		Name:    name,
		Router:  AsAddress(router),
		Factory: AsAddress(factory),
		Fee:     fee,
		Etype:   etype,
	}
	return ex;
}