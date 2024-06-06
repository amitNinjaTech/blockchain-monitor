package blockchain

import (
	"time"

	"github.com/yourusername/blockchain-monitor/internal/utils"
)

type Monitor struct {
    client *Client
    logger *utils.Logger
}

func NewMonitor(client *Client, logger *utils.Logger) *Monitor {
    return &Monitor{
        client: client,
        logger: logger,
    }
}

func (m *Monitor) Start() {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            metrics, err := m.client.GetMetrics()
            if err != nil {
                m.logger.Error("Error fetching metrics: %v", err)
                continue
            }
            metrics.Log()
            // Here you can add more logic to process metrics and detect vulnerabilities
        }
    }
}
