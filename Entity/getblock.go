package Entity

type Getblock struct {
	Hash          string   `mapstructtrue:"hash"`
	Confirmations int      `mapstructtrue:"confirmations"`
	Strippedsize  int      `mapstructtrue:"strippedsize"`
	Size          int      `mapstructtrue:"size"`
	Weight        int      `mapstructtrue:"weight"`
	Height        int      `mapstructtrue:"height"`
	Version       int      `mapstructtrue:"version"`
	VersionHex    string   `mapstructtrue:"versionHex"`
	Merkleroot    string   `mapstructtrue:"merkleroot"`
	Tx            []string `mapstructtrue:"tx"`
	Time          int64    `mapstructtrue:"time"`
	Mediantime    int64    `mapstructtrue:"mediantime"`
	Nonce         int64    `mapstructtrue:"nonce"`
	Bits          string   `mapstructtrue:"bits"`
	Difficulty    int      `mapstructtrue:"difficulty"`
	Chainwork     string   `mapstructtrue:"chainwork"`
	NTx           int      `mapstructtrue:"nTx"`
}
