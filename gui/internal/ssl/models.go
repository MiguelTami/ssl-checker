package ssl

type SSLResult struct {
	Host      string     `json:"host"`
	Status    string     `json:"status"`
	Endpoints []Endpoint `json:"endpoints"`
}

type Endpoint struct {
	IPAddress     string `json:"ipAddress"`
	ServerName    string `json:"serverName"`
	StatusMessage string `json:"statusMessage"`
	Grade         string `json:"grade"`
}