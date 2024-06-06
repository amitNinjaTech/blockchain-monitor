package blockchain

import (
	"log"
	"net/http"

	"github.com/amitNinjaTech/blockchain-monitor/pkg/blockchain"
)

type Client struct {
    NodeURL string
    httpClient *http.Client
}

func NewClient(cfg blockchain.BlockchainConfig) *Client {
    return &Client{
        NodeURL: cfg.NodeURL,
        httpClient: &http.Client{},
    }
}

func (c *Client) GetMetrics() (blockchain.Metrics, error) {
    // Placeholder for actual implementation to get metrics from blockchain
    log.Println("Fetching blockchain metrics...")
    // Return dummy data for now
    return blockchain.Metrics{}, nil
}
