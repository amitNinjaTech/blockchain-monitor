package blockchain

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Client struct {
    NodeURL    string
    httpClient *http.Client
}

func NewClient(cfg BlockchainConfig) *Client {
    return &Client{
        NodeURL:    cfg.NodeURL,
        httpClient: &http.Client{},
    }
}

type Metrics struct {
    BlockHeight int64   `json:"blockHeight"`
    TxCount     int64   `json:"txCount"`
    PendingTx   int64   `json:"pendingTx"`
    AvgBlockTime float64 `json:"avgBlockTime"`
    HashRate    float64 `json:"hashRate"`
}

func (c *Client) GetMetrics() (Metrics, error) {
    // Example request to get the latest block number
    resp, err := c.httpClient.Get(fmt.Sprintf("%s/api?module=proxy&action=eth_blockNumber", c.NodeURL))
    if err != nil {
        return Metrics{}, err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return Metrics{}, err
    }

    blockHeightHex, ok := result["result"].(string)
    if !ok {
        return Metrics{}, fmt.Errorf("invalid response from node")
    }

    var blockHeight int64
    fmt.Sscanf(blockHeightHex, "0x%x", &blockHeight)

    // Placeholder for other metrics, set dummy values for now
    metrics := Metrics{
        BlockHeight: blockHeight,
        TxCount:     0,   // Implement actual logic
        PendingTx:   0,   // Implement actual logic
        AvgBlockTime: 0.0, // Implement actual logic
        HashRate:    0.0, // Implement actual logic
    }

    return metrics, nil
}
