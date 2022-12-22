# Luis Coin
This project consists in a Golang REST API which exposes a simple Solidity smart contract. This contract consists in a simple implementation of a cryptocurrency. The supported operations are check the coin balance for a specific account and make a transfer between accounts. 

This project was made using Ganache as test blockchain so you need to provide a RPC_SERVER_URL for the application to connect and deploy the contract.

## How to run 
1. docker build -t luis-coin .
2. docker run -p 8080:8080 --name luis-coin -e ADMIN_ACCOUNT_KEY=\<your admin private key> -e RPC_SERVER_URL=\<your rpc server url> luis-coin. E.g: `docker run -p 8080:8080 --name luis-coin -e ADMIN_ACCOUNT_KEY=e6df1e73b2716b40141e0665433ee5fa4834e89dab533e3d505ae6a2aff61d99 -e RPC_SERVER_URL=http://host.docker.internal:7545 luis-coin`

## Required environment variables
|     Variable      |                    Description                    |                             Example                              |
|:-----------------:|:-------------------------------------------------:|:----------------------------------------------------------------:|
|       Port        |      Port where the application will respond      |                              :8080                               |
| ADMIN_ACCOUNT_KEY | Blockchain account private key to deploy contract | e6df1e73b2716b40141e0665433ee5fa4834e89dab533e3d505ae6a2aff61d99 |
|  RPC_SERVER_URL   |           Blockchain RPC server address           |                      http://127.0.0.1:7545                       |