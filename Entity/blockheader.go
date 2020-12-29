package Entity

type BlockHeader struct {
	Hash          string   `mapstructtrue:"hash"`
	Confirmations int      `mapstructtrue:"confirmations"`
	Height        int      `mapstructtrue:"height"`
	Version       int      `mapstructtrue:"version"`
	VersionHex    string   `mapstructtrue:"versionHex"`
	Merkleroot    string   `mapstructtrue:"merkleroot"`
	Time          int64    `mapstructtrue:"time"`
	Mediantime    int64    `mapstructtrue:"mediantime"`
	Nonce         int64    `mapstructtrue:"nonce"`
	Bits          string   `mapstructtrue:"bits"`
	Difficulty    int      `mapstructtrue:"difficulty"`
	Chainwork     string   `mapstructtrue:"chainwork"`
	NTx           int      `mapstructtrue:"nTx"`
}
