package Entity

type Balances struct {
	Mine mine `mapstructtrue:"mine"`
}
type mine struct {
	Trusted           float64 `mapstructtrue:"trusted"`
	Untrusted_pending float64 `mapstructtrue:"untrusted_pending"`
	Immature          float64 `mapstructtrue:"immature"`
}
