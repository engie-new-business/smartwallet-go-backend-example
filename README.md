# Smartwallet backend example

This repository contains example of backend services you can develop using Rockside. Thoses services can be used for example to manage a smartwallet on ios or android client application.

## Deploy Gnosis Safe

This function deploy a Gnosis Safe smartwallet with the provided account as owner.
The transaction is sent to a Rockside Forwarder that will be in charge to forward to the Gnosis Proxy Factory and to payback Rockside Relayer for the Tx gas fees.

First you need to deploy a forwarder using Rockside API. [Look at your docs](https://docs.rockside.io/rockside-api#deploy-a-forwarder-contract).

On this example we use an EOA managed by Rockside to sign the message. So you have to create an EOA using Rockside API.

```sh
curl --location --request POST 'https://api.rockside.io/ethereum/eoa' \
--header 'apikey: YOUR_API_KEY' \
--header 'Content-Type: application/json' \
--data-raw ''
```

To payback for the transaction fees your forwarder need to have some funds.

For more details on how Rockside can interact with Gnosis Safe, have a look to [this repository](https://github.com/rocksideio/rockside-integration-examples/tree/master/gnosis-safe). 

### Environment Variables

```
APIKEY: Your Rockside APIKEY
BASE_URL: Rockside base url - https://api.rockside.io
FORWARDER_ADDRESS: Address of the forwarder you deployed
FACTORY_ADDRESS: Address of the Gnosis Proxy factory
SMARTWALLET_IMPL_ADDR: Address of Gnosis Safe Mastercopy
ADMIN_ADDRESS: Address of the EOA managed with Rockside
```

Gnosis safe mastercopy  and factory address can be found here: [Gnosis Safe Contract Address](https://github.com/gnosis/safe-contracts/tree/v1.1.1/.openzeppelin)

For Ropsten you can use the one we deployed.
```
Gnosis Safe Proxy Factory: 0x016457118b425fe86952381eC5127F28D4248984
Gnosis Safe Master Copy: 0xB6998f4E968573534D6ea6A500323B0d1cd03767
```

### Deploy Gnosis Request example

Request:

```sh
curl --location --request POST 'DEPLOY_GNOSIS_FUNC_URL' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account": "YOUR_GNOSIS_OWNER_ACCOUNT"
}'
```

Response:

```js
{
    "transaction_hash": "0x6663a81a1a827c4bf2301eb169de900c51d2b6e4e2c26d503dce10888f8cdee9",
    "tracking_id": "01E9ZSDHMYYFMW3E1CVQ9ADVHK"
}
```


The address of the Gnosis safe can be found on log's data of the transaction receipt


## Relay Parameters

This service allow your client application to get the gas price and the Rockside relayer address to use for sending there transaction depending on the desired speed.

### Environment Variables

```
APIKEY: Your Rockside APIKEY
BASE_URL: Rockside base url - https://api.rockside.io
```

### Get Relay Parameters Transaction example

Request:

```sh
curl --location --request POST 'GET_TX_PARAMS_URL?network=mainnet' \
--header 'Content-Type: application/json' \
--data-raw '{
    "address": "0x4b435b57ab0aa32f169aba415122a4a981dad2a1"
}'
```

Response:

```
{
    {
    "speeds": {
        "fast": {
            "gas_price": "124000000000",
            "relayer": "0xb20328AA6B950986798F9755c68DEF45e2631Ec4"
        },
        "fastest": {
            "gas_price": "134000000000",
            "relayer": "0xb20328AA6B950986798F9755c68DEF45e2631Ec4"
        },
        "safelow": {
            "gas_price": "112000000000",
            "relayer": "0xb20328AA6B950986798F9755c68DEF45e2631Ec4"
            },
        "standard": {
            "gas_price": "116000000000",
            "relayer": "0xb20328AA6B950986798F9755c68DEF45e2631Ec4"
        }
    }
    }
}
```

## Relay Transaction

This service take the encoted data you want to send to your Gnosis Safe Wallet and execute the transaction with rockside relay API.

### Environment Variables

```
APIKEY: Your Rockside APIKEY
BASE_URL: Rockside base url - https://api.rockside.io
```

### Relay transactions request example

Request:

```sh
curl --location --request POST 'RELAY_TX_URL?network=mainnet' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to": "YOUR_GNOSIS_ADDRESS",
    "data": "ENCODED_DATA"
}'
```

Response:

```js
{
    "transaction_hash": "0x6663a81a1a827c4bf2301eb169de900c51d2b6e4e2c26d503dce10888f8cdee9",
    "tracking_id": "01E9ZSDHMYYFMW3E1CVQ9ADVHK"
}
```

## Transaction Infos

Because Rockside can replay your transaction to make it validated faster, the tx hash can change. Use this service to follow the status of your transaction and get the transaction receipt.

### Transaction infos request example

Request:

```sh
curl --location --request POST 'RELAY_TX_URL?network=mainnet' \
--header 'Content-Type: application/json' \
--data-raw '{
    "to": "YOUR_GNOSIS_ADDRESS",
    "data": "ENCODED_DATA"
}'
```

Response:

```js
{
    "transaction_hash": "0x6663a81a1a827c4bf2301eb169de900c51d2b6e4e2c26d503dce10888f8cdee9",
    "tracking_id": "01E9ZSDHMYYFMW3E1CVQ9ADVHK"
}
```


### Environment Variables

```
APIKEY: Your Rockside APIKEY
BASE_URL: Rockside base url - https://api.rockside.io
```
