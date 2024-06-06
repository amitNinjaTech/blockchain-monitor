package blockchain

import (
	"log"
)

type Metrics struct {
    BlockHeight       int64
    TransactionCount  int64
    PendingTxCount    int64
    AvgBlockTime      float64
    HashRate          float64
}

func (m Metrics) Log() {
    log.Printf("BlockHeight: %d, TransactionCount: %d, PendingTxCount: %d, AvgBlockTime: %f, HashRate: %f",
        m.BlockHeight, m.TransactionCount, m.PendingTxCount, m.AvgBlockTime, m.HashRate)
}
