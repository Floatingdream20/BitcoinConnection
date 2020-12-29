package Entity

type Networkinfo struct {
	Version            int64
	Subversion         string
	Protocolversion    int64
	Localservices      int64
	Localservicesnames []string
	Localrelay         bool
	Timeoffset         int64
	Networkactive      bool
	Connections        int64
	Networks           []netWork
	Relayfee           float64
	Incrementalfee     float64
	Localaddresses     []localaddresse
	Warnings           string
}

type localaddresse struct {
	Address string
	Port    int64
	Score   int64
}
type netWork struct {
	Name                        string
	Limited                     bool
	Reachable                   bool
	Proxy                       string
	Proxy_randomize_credentials string
}
