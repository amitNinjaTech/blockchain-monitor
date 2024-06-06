package blockchain

type BlockchainConfig struct {
    NodeURL string
}

type Metrics struct {
    BlockHeight       int64
    TransactionCount  int64
    PendingTxCount    int64
    AvgBlockTime      float64
    HashRate          float64
}
