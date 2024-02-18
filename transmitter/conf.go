package transmitter

type Conf struct {
	DeviceName   string         `json:"device_name"`
	ServicesConf map[string]any `json:"services"`
}
