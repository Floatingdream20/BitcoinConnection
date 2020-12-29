package Entity

type Chaintxstats struct {
	Time                      int64  `mapstructtrue:"time"`
	Txcount                   int    `mapstructtrue:"txcount"`
	Window_final_block_hash   string `mapstructtrue:"window_final_block_hash"`
	Window_final_block_height int    `mapstructtrue:"window_final_block_height"`
	Window_block_count        int    `mapstructtrue:"window_block_count"`
}
