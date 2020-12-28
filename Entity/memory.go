package Entity

type Locked struct {
	Lockde lockde `mapstructtrue:"lockde"`
}
type lockde struct {
	Used        int   `mapstructtrue:"used"`
	Free        int64 `mapstructtrue:"free"`
	Total       int64 `mapstructtrue:"total"`
	Lockde      int64 `mapstructtrue:"lockde"`
	Chunks_used int   `mapstructtrue:"chunks_used"`
	Chunks_free int   `mapstructtrue:"chunks_free"`
}
