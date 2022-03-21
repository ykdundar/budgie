# budgie

## API Key

First of all, obtain an API key from the [marketstack](https://marketstack.com/product).

## Configuration

In order to use budgie, first you need to enter your API token as follows:

```sh
budgie config --token="YOUR_API_TOKEN"
```

If you want to update your token, you can run the same command again. The latest token will be used while using budgie.

## Usage

### Portfolio Operations

#### Create Portfolio

```sh
budgie portfolio create --name="European Stocks" --currency="USD" --active=true
```

#### Update Portfolio

```sh
budgie portfolio update --name="European Stocks" --rename="German Stocks" --currency="EUR" --active=false
```

#### Delete Portfolio

```sh
budgie portfolio delete --name="European Stocks"
```

#### List All Portfolios

```sh
budgie portfolio list
```

#### Show a Single Portfolio

```sh
budgie portfolio show # shows the active portfolio
budgie portfolio show --name="European Stocks"
```

### Stock Operations

#### Add Stock to Portfolio

```sh
# Runs on search and asks "did you mean?"
budgie stock add --portfolio="European Stocks" --ticker="MSFT" 
```

#### Remove Stock from Portfolio *

```sh
budgie stock remove --portfolio="European Stocks" --ticker="MSFT"
```

### Transaction Operations

#### Buy Stock to Transactions

```sh
budgie transaction buy --ticker="MSFT" --price=10.5 --date="02.01.2006" --shares=5 
```

#### Sell Stock to Transactions

```sh
budgie transaction sell --ticker="MSFT" --price=10.5 --date="02.01.2006" --shares=5 
```

#### Remove Stock from Transactions

```sh
budgie transaction remove --id=1
```

#### Report Transaction Earnings/Losses

##### Report All Transactions

```sh
budgie transaction report 
```
##### Report Transaction Earnings/Losses for The Given Number of Days

```sh
budgie transaction report day 30
```

##### Report Transaction Earnings/Losses for The Given Number of Months

```sh
budgie transaction report month 5
```

##### Report Transaction Earnings/Losses for The Given Number of Years

```sh
budgie transaction report year 5
```

#### List All Transactions Records

```sh
budgie transaction list
```