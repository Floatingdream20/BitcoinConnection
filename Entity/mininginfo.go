package Entity

type MiningInfo struct {
	Blocks        int64
	Difficulty    float64
	Networkhashps float64
	Pooledtx      int64
	Chain         string
	Warnings      string
}
