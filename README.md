### Command line currency converter

This tool is designed to convert the amount (*provided as the 1st argument*) in one currency (*provided as the 2nd argument*) to another currency (*provided as the 3rd argument*). The conversion is based on the data provided by [this API](https://pro.coinmarketcap.com/api/v1).

### Building

To build the project, run

```
go build ./cmd/converter
```

in the root of the repository.

### API Key
To use the API, you need to create an account on [CoinMarketCap](https://coinmarketcap.com/) and get an API key.

After you create your account, copy the API key and run the following command:
```
export COINMARKET_API_KEY=<your key>
```
where <your key> is replaced with your key.

### Usage

To run the project, use

```
./converter <amount> <from> <to>
```

command, where 
- <amount> is the amount you want to convert
- <from> is the base currency
- <to> is the desired currency

*Example*:
```
./converter 100 ETH USD
```
This command will convert 100 ETH to USD.

The tool is **case-insensitive**, so
```
./converter 100 eth usd
```
works too.
