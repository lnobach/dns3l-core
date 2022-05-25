package apiv1

type ServerInfo struct {
	Version *ServerInfoVersion `json:"version"`
	Contact *ServerInfoContact `json:"contact"`
}

type ServerInfoVersion struct {
	Daemon string `json:"daemon"`
	API    string `json:"api"`
}

type ServerInfoContact struct {
	URL   string   `json:"url"`
	EMail []string `json:"email"`
}

type DNSHandlerInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Feature     []string `json:"feature"`
	ZoneNesting bool     `json:"zoneNesting"`
}

type DNSRootzoneInfo struct {
	Root    string   `json:"root"`
	AutoDNS string   `json:"autodns"`
	AcmeDNS string   `json:"acmedns"`
	CA      []string `json:"ca"`
}

type CAInfo struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"desc"`
	LogoPath    string   `json:"logo"`
	URL         string   `json:"url"`
	Roots       string   `json:"roots"`
	TotalValid  int      `json:"totalValid"`
	TotalIssued int      `json:"totalIssued"`
	Type        string   `json:"type"`
	IsAcme      bool     `json:"acme"`
	Rootzones   []string `json:"rtzn"`
	Enabled     bool
}