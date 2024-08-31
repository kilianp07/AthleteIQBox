package buttons

type Configuration struct {
	Switches map[string]ButtonConfiguration `json:"switches"`
}

type ButtonConfiguration struct {
	Gpio string `json:"gpio"`
	name string
}
