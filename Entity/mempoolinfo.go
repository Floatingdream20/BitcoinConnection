package Entity

type Mempoollinfo struct {
	Loaded        bool    `mapstructtrue:"loaded"`
	Size          int     `mapstructtrue:"size"`
	Bytes         int     `mapstructtrue:"bytes"`
	Usage         int     `mapstructtrue:"usage"`
	Maxmempool    int64   `mapstructtrue:"maxmempool"`
	Mempoolminfee float64 `mapstructtrue:"mempoolminfee"`
	Minrelaytxfee float64 `mapstructtrue:"minrelaytxfee"`
}
