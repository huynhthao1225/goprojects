package config

type ConnectionProperty struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type ConnectionProperties struct {
	ConnectionProperties []ConnectionProperty `json:"connectionProperties"`
}
