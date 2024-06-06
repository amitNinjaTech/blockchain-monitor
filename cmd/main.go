package main

import (
    "blockchain-monitor/config"
    "blockchain-monitor/internal/blockchain"
    "blockchain-monitor/internal/utils"
    "log"
)

func main() {
    cfg, err := config.LoadConfig("config.json")
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    client := blockchain.NewClient(cfg.Blockchain)
    monitor := blockchain.NewMonitor(client)
    
    // Start monitoring
    go monitor.Start()

    // Prevent main from exiting
    select {}
}
