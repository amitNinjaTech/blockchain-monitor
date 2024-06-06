package blockchain

import (
    "log"
    "time"
    "blockchain-monitor/internal/utils"
)

type Monitor struct {
    Client *Client
}

func NewMonitor(client *Client) *Monitor {
    return &Monitor{Client: client}
}

func (m *Monitor) Start() {
    ticker := time.NewTicker(10 * time.Second)
    for range ticker.C {
        metrics, err := m.Client.GetMetrics()
        if err != nil {
            log.Printf("Failed to get metrics: %v", err)
            continue
        }

        log.Printf("Metrics: %+v", metrics)

        // Example alert conditions
        if metrics.TxCount > 1000 {  // Example condition, adjust as needed
            utils.SendAlert("High Transaction Volume", "Transaction count exceeded threshold: "+string(metrics.TxCount))
        }

        if metrics.PendingTx > 500 {  // Example condition, adjust as needed
            utils.SendAlert("High Pending Transactions", "Pending transactions exceeded threshold: "+string(metrics.PendingTx))
        }

        // Add more conditions as needed
    }
}
