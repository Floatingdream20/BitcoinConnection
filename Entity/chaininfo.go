package Entity

type ChainInfo struct {
	Chain                string  `mapstructtrue:"chain"`
	Blocks               int     `mapstructtrue:"blocks"`
	Headers              int     `mapstructtrue:"headers"`
	Bestblockhash        string  `mapstructtrue:"bestblockhash"`
	Difficulty           int     `mapstructtrue:"difficulty"`
	Mediantime           int64   `mapstructtrue:"mediantime"`
	Verificationprogress float64 `mapstructtrue:"verificationprogress"`
	Initialblockdownload bool    `mapstructtrue:"initialblockdownload"`
	Chainwork            string  `mapstructtrue:"chainwork"`
	Size_on_disk         int     `mapstructtrue:"size_on_disk"`
	Pruned               bool    `mapstructtrue:"prunde"`
	Softforks            bip     `mapstructtrue:"softforks"`
	Warnings             string
}
type bip struct {
	Bip34  bipdata `mapstructtrue:"bip34"`
	Bip66  bipdata `mapstructtrue:"bip66"`
	Bip65  bipdata `mapstructtrue:"bip65"`
	Csv    bipdata `mapstructtrue:"csv"`
	Segwit bipdata `mapstructtrue:"segwit"`
}
type bipdata struct {
	Type   string `mapstructtrue:"type"`
	Active bool   `mapstructtrue:"active"`
	Height int64  `mapstructtrue:"height"`
}
