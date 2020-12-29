package Entity

type Txoutsetinfo struct {
	Height            int     `mapstructtrue:"height"`
	Bestblock         string  `mapstructtrue:"bestblock"`
	Transactions      int     `mapstructtrue:"transactions"`
	Txouts            int     `mapstructtrue:"txouts"`
	Bogosize          int     `mapstructtrue:"bogosize"`
	Hash_serialized_2 string  `mapstructtrue:"hash_serialized_2"`
	Disk_size         int     `mapstructtrue:"disk_size"`
	Total_amount      float64 `mapstructtrue:"total_amount"`
}
