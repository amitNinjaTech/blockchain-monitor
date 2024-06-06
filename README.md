# Blockchain Monitor

This project is an external agent written in Go to monitor a blockchain network. The agent extracts relevant operational metrics to identify potential vulnerabilities, threats, and attacks in real-time.

## Setup Instructions

### Prerequisites
- Docker
- Docker Compose (optional)
- Go (for local development)

### Configuration

Create a `config.json` file in the root directory with the following structure:
'''
{
    "blockchain": {
        "node_url": "YOUR_BLOCKCHAIN_NODE_URL" (Sample api key for Infura is added for testing ethereum nodes)
    },
    "log_level": "info"
}
'''

Replace YOUR_BLOCKCHAIN_NODE_URL with the URL of your blockchain node. This can be an open node or a local blockchain setup.

Using Open Nodes
Identify Open Nodes:

For Ethereum, use services like Infura or Alchemy.
For Bitcoin, use services like Blockstream or Blockchain.info.
Get Access:

Sign up for the service and obtain the necessary API keys or project IDs.
Update Configuration:

Update config.json with the node URL, for example:

{
    "blockchain": {
        "node_url": "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID"
    },
    "log_level": "info"
}


## Build and Run the Docker Container:
docker build -t blockchain-monitor .
docker run -d -p 8080:8080 --name blockchain-monitor -v $(pwd)/config.json:/app/config.json blockchain-monitor

## Verify:

Check the logs to ensure the application is running and fetching metrics:

docker logs blockchain-monitor

### Using a Local Blockchain Setup

Ethereum (Using Ganache)

Install and Run Ganache:

Download and install Ganache.
Start Ganache and note the RPC server URL (e.g., http://127.0.0.1:7545).
Update Configuration:

{
    "blockchain": {
        "node_url": "http://127.0.0.1:7545"
    },
    "log_level": "info"
}


## Build and run the Docker Containers 
docker build -t blockchain-monitor .
docker run -d -p 8080:8080 --name blockchain-monitor -v $(pwd)/config.json:/app/config.json blockchain-monitor


#3 Verify:
Check the logs to ensure the application is running and fetching metrics:

docker logs blockchain-monitor


Metric Justification

Metrics Monitored
Block Height:

Justification: The current height of the blockchain can help detect chain reorganization attacks (reorgs) and ensure the node is synchronized with the network.
Usage: Sudden drops or anomalies in block height can indicate a reorg or potential fork.
Transaction Count:

Justification: The number of transactions can indicate network activity and potential spam attacks.
Usage: Unusually high transaction counts can signal spam attacks or stress tests on the network.
Pending Transactions:

Justification: Monitoring pending transactions helps identify network congestion and potential denial-of-service (DoS) attacks.
Usage: A high number of pending transactions may indicate a DoS attack or network congestion.
Average Block Time:

Justification: The average time to produce a block is crucial for assessing network health and performance.
Usage: Variations in average block time can indicate network issues or mining attacks.
Hash Rate:

Justification: The hash rate represents the computational power of the network, critical for detecting mining attacks.
Usage: A sudden drop in hash rate can indicate a mining attack or loss of network security.
Approach Explanation
Chosen Approach: Open Nodes and Local Blockchain Setup

Open Nodes:

Rationale: Using open nodes like Infura or Alchemy provides access to real-time data from live blockchain networks without the overhead of running and maintaining a full node. This is suitable for monitoring the actual state of the network and for long-term monitoring solutions.
Local Blockchain Setup:

Rationale: A local blockchain setup, such as Ganache for Ethereum or Bitcoin Core in regtest mode, allows for controlled environments where specific scenarios can be simulated. This is ideal for development, testing, and simulating attack vectors or network conditions without impacting real networks.
Both approaches are suitable depending on the use case:

Open Nodes: Ideal for real-time, live network monitoring and for use cases where setting up a local node is not feasible.
Local Blockchain: Ideal for development, testing, and scenarios where controlled experiments are necessary.
By combining both approaches, the application can be robustly tested and deployed for real-world monitoring and development purposes.


