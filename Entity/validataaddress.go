package Entity

type Validata struct {
	Isvalid         bool   `mapstructtrue:"isvalid"`
	Address         string `mapstructtrue:"address"`
	ScriptPubKey    string `mapstructtrue:"scriptPubKey"`
	Isscript        bool   `mapstructtrue:"isscript"`
	Iswitness       bool   `mapstructtrue:"iswitness"`
	Witness_version int    `mapstructtrue:"witness_version"`
	Witness_program string `mapstructtrue:"witness_program"`
}
