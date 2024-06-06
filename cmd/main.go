package main

import (
	"log"

	"github.com/amitNinjaTech/blockchain-monitor/config"
	"github.com/amitNinjaTech/blockchain-monitor/internal/blockchain"
	"github.com/amitNinjaTech/blockchain-monitor/internal/utils"
)

func main() {
    // Load configuration
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Could not load configuration: %v", err)
    }

    // Initialize logger
    logger := utils.NewLogger(cfg.LogLevel)

    // Initialize blockchain client
    client := blockchain.NewClient(cfg.Blockchain)

    // Start monitoring
    monitor := blockchain.NewMonitor(client, logger)
    monitor.Start()
}
