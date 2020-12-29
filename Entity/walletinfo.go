package Entity

type Walletinfo struct {
	Walletname              string
	Walletversion           int64
	Balance                 float64
	Unconfirmed_balance     float64
	Immature_balance        float64
	Txcount                 int64
	Keypoololdest           int64
	Keypoolsize             int64
	Hdseedid                string
	Keypoolsize_hd_internal int64
	Paytxfee                float64
	Private_keys_enabled    bool
	Avoid_reuse             bool
	Scanning                bool
}
