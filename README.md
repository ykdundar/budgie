# budgie

## API Key

First of all, obtain an API key from the [marketstack](https://marketstack.com/product).

## Usage

### Configuration

```sh
budgie config --token "YOUR_API_TOKEN"
```

### Portfolio Operations

#### Create Portfolio

```sh
budgie portfolio create --name "European Stocks" --currency "USD" --active true
```

#### Update Portfolio

```sh
budgie portfolio update --name "European Stocks" --new_name "German Stocks" --currency "EUR" --active false
```

#### Delete Portfolio

```sh
budgie portfolio destroy --name "European Stocks"
```

#### List All Portfolios

```sh
budgie portfolio list
```

#### Show a Single Portfolio

```sh
budgie portfolio show # shows the active portfolio
budgie portfolio show --name "European Stocks"
```

### Stock Operations

#### Add Stock to Portfolio

```sh
# Runs on search and asks "did you mean?"
budgie stock add --portfolio "European Stocks" --ticker "MSFT" 
```

#### Remove Stock from Portfolio *

```sh
budgie stock remove --portfolio "European Stocks" --ticker "MSFT"
```

### Buy and Sell Operations

#### Buy a Stock

```sh
budgie stock buy --portfolio "European Stocks" --ticker "MSFT" --date "06.02.2022" --price "180" --shares "20" --currency "USD"
```

#### Sell a Stock

```sh
# sell and remove are aliases
budgie stock remove --portfolio "European Stocks" --ticker "MSFT" --date "06.02.2020" --price "180" --shares "20"
budgie stock sell --portfolio "European Stocks" --ticker "MSFT" --date "06.02.2020" --price "180" --shares "20"
```

### Reporting

```sh
budgie stock report
budgie stock report --today # since today
budgie stock report --day 30
bugdie stock report --week 2
budgie stock report --month 5
budgie stock report --year 2
```
