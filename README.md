# budgie

![release](https://github.com/ykdundar/budgie/actions/workflows/release.yml/badge.svg)
[![GoDoc](https://godoc.org/github.com/ykdundar/budgie?status.svg)](https://godoc.org/github.com/ykdundar/budgie)
![Supported Version](https://img.shields.io/badge/go%20version-%3E%3D1.14-turquoise)
[![Go Report Card](https://goreportcard.com/badge/github.com/ykdundar/budgie)](https://goreportcard.com/report/github.com/ykdundar/budgie)
[![Maintainability](https://api.codeclimate.com/v1/badges/5199df38a19f2964163a/maintainability)](https://codeclimate.com/github/ykdundar/budgie/maintainability)
[![License](https://img.shields.io/github/license/ykdundar/budgie)](https://github.com/ykdundar/budgie/blob/main/LICENSE)

`budgie` is a CLI tool that allows you to keep a track of your portfolio and your investments.
Under the hood it uses APIs of [marketstack](https://marketstack.com) to provide intraday and
end of day price data. `budgie` stores everything locally (in a locally stored SQLite database)
and doesn't share any of your financial data online.

## Install

```sh
go install github.com/ykdundar/budgie@latest
```

> budgie uses Go Modules to manage dependencies, and supports Go versions >=1.14.x.

## API Key

First of all, obtain an API key from the [marketstack](https://marketstack.com/product).

## Configuration

In order to use `budgie`, first you need configure the tool with your API token as follows:

```sh
budgie config --token="YOUR_API_TOKEN"
```

If you want to update your token, you can run the same command again. The latest token will take precedence.

## Usage

Run `budgie -h` to print help instructions.

### Portfolio Operations

#### Add a New Portfolio

```sh
budgie portfolio add --name="European Stocks"
```

#### Update a Portfolio

```sh
budgie portfolio update --name="European Stocks" --rename="US Stocks"
```

#### Delete a Portfolio

```sh
budgie portfolio delete --name="European Stocks"
```

#### List All Portfolios

```sh
budgie portfolio list
```

```
╭─────────────────────────────────╮
│          MY PORTFOLIOS          │
├───────────────┬─────────────────┤
│            ID │ NAME            │
├───────────────┼─────────────────┤
│             1 │ US Stocks       │
╰───────────────┴─────────────────╯
```

#### Show a Single Portfolio

> Prints `open`, `high`, `low`, `last`, `close`, `date` and `exchange` attributes of stocks in your portfolio

```sh
budgie portfolio show --name="US Stocks"
```

```
╭────────────────────────────────────────────────────────────────────────────╮
│                               US STOCKS                                    │
├────────┬────────┬─────────┬────────┬──────┬───────┬────────────┬───────────┤
│ SYMBOL │   OPEN │    HIGH │    LOW │ LAST │ CLOSE │ DATE       │ EXCHANGE  │
├────────┼────────┼─────────┼────────┼──────┼───────┼────────────┼───────────┤
│ AAPL   │ 161.73 │  161.73 │ 161.73 │    0 │     0 │ 2022-04-23 │ IEXG      │
├────────┼────────┼─────────┼────────┼──────┼───────┼────────────┼───────────┤
│ MSFT   │ 273.94 │  273.94 │ 273.94 │    0 │     0 │ 2022-04-23 │ IEXG      │
╰────────┴────────┴─────────┴────────┴──────┴───────┴────────────┴───────────╯
```

### Stock Operations

#### Add a Stock to a Portfolio

```sh
budgie stock add --portfolio="US Stocks" --ticker="MSFT"
```

#### Remove a Stock from a Portfolio

```sh
budgie stock remove --portfolio="US Stocks" --ticker="MSFT"
```

#### Search Stocks in Stock Exchanges

```sh
budgie stock search --name="Microsoft"
```

```
╭─────────────────────────────────────────────────────────────────────────────╮
│                                  MICROSOFT                                  │
├────────────────────────┬─────────────┬───────────┬───────────┬──────────────┤
│ NAME                   │ SYMBOL      │ EXCHANGE  │ COUNTRY   │ CITY         │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ Microsoft Corporation  │ MSFT        │ NASDAQ    │ USA       │ New York     │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT DRN          │ MSFT34.BVMF │ Bovespa   │ Brazil    │ Sao Paulo    │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFT.XSGO   │ BVS       │ Chile     │ Santiago     │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT DL-,00000625 │ MSF.XETRA   │ XETR      │ Germany   │ Frankfurt    │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT DL-,00000625 │ MSF.XFRA    │ FSX       │ Germany   │ Frankfurt    │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFT.XMIL   │ MIL       │ Italy     │ Milano       │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFT.XMEX   │ BMV       │ Mexico    │ Mexico City  │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFT.XBUE   │ BCBA      │ Argentina │ Buenos Aires │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFTD.XBUE  │ BCBA      │ Argentina │ Buenos Aires │
├────────────────────────┼─────────────┼───────────┼───────────┼──────────────┤
│ MICROSOFT CORP         │ MSFT        │ IEX       │ USA       │ New York     │
╰────────────────────────┴─────────────┴───────────┴───────────┴──────────────╯
```

#### Show Price Data of Any Stock

- You can search for price data of any stock by providing a comma separated list of symbols.

```sh
budgie stock show --ticker "MSFT, AAPL"
```

```
╭────────────────────────────────────────────────────────────────────────────╮
│                               AAPL, MSFT                                   │
├────────┬────────┬────────┬────────┬──────┬───────┬────────────┬────────────┤
│ SYMBOL │  OPEN  │   HIGH │  LOW   │ LAST │ CLOSE │ DATE       │ EXCHANGE   │
├────────┼────────┼────────┼────────┼──────┼───────┼────────────┼────────────┤
│ AAPL   │ 161.73 │ 161.73 │ 161.73 │    0 │     0 │ 2022-04-23 │ IEXG       │
├────────┼────────┼────────┼────────┼──────┼───────┼────────────┼────────────┤
│ MSFT   │ 273.94 │ 273.94 │ 273.94 │    0 │     0 │ 2022-04-23 │ IEXG       │
╰────────┴────────┴────────┴────────┴──────┴───────┴────────────┴────────────╯
```

### Transaction Operations

#### Buy/Sell a Stock

```sh
budgie transaction buy --ticker="MSFT" --price=10.5 --date="02.01.2006" --shares=50
budgie transaction sell --ticker="MSFT" --price=15 --date="03.01.2006" --shares=20
```

## List All Transactions

```sh
budgie transaction list
```

```
╭─────────────────────────────────────────────────────────────────────────────────────────╮
│                                     MY TRANSACTIONS                                     │
├────┬────────┬───────────┬─────────────┬────────┬───────┬──────────────┬─────────────────┤
│ ID │ SYMBOL │ CATEGORY  │ DATE        │ SHARES │ PRICE │ MARKET VALUE │  PURCHASE VALUE │
├────┼────────┼───────────┼─────────────┼────────┼───────┼──────────────┼─────────────────┤
│  1 │ AAPL   │ purchase  │ 2022-4-20   │     10 │   90  │ 1617.90      │             900 │
├────┼────────┼───────────┼─────────────┼────────┼───────┼──────────────┼─────────────────┤
│  2 │ AAPL   │ purchase  │ 2022-4-20   │     20 │  100  │ 3235.80      │            2000 │
├────┼────────┼───────────┼─────────────┼────────┼───────┼──────────────┼─────────────────┤
│  3 │ MSFT   │ purchase  │ 2022-4-20   │     30 │  110  │ 8220.90      │            3300 │
├────┼────────┼───────────┼─────────────┼────────┼───────┼──────────────┼─────────────────┤
│  4 │ MSFT   │ purchase  │ 2022-4-20   │     40 │  115  │ 10961.20     │            4600 │
├────┼────────┼───────────┼─────────────┼────────┼───────┼──────────────┼─────────────────┤
│  5 │ AAPL   │ sale      │ 2022-4-20   │      1 │   90  │ 161.79       │              90 │
╰────┴────────┴───────────┴─────────────┴────────┴───────┴──────────────┴─────────────────╯
```


#### Remove a Transaction from History

```sh
budgie transaction remove --id=1
```

#### Report Transactions (Earnings/Losses)

```sh
budgie transaction report 
```

```
╭───────────────────────────────────────────────────────────────╮
│                       MY EARNINGS/LOSSES                      │
├────────┬─────────┬────────────────┬──────────────┬────────────┤
│ SYMBOL │  SHARES │ PURCHASE VALUE │ MARKET VALUE │ DIFFERENCE │
├────────┼─────────┼────────────────┼──────────────┼────────────┤
│ AAPL   │       1 │             90 │ 161.79       │     -71.78 │
├────────┼─────────┼────────────────┼──────────────┼────────────┤
│ AAPL   │      30 │           2900 │ 4853.70      │   -1953.70 │
├────────┼─────────┼────────────────┼──────────────┼────────────┤
│ MSFT   │      70 │           7900 │ 19182.10     │  -11282.10 │
╰────────┴─────────┴────────────────┴──────────────┴────────────╯
```

You can time limit the `report` command by specifying a `day`, a `month` or a `year` as follows:

```sh
budgie transaction report day 30 # transaction report of the last 30 days
budgie transaction report month 5 # transaction report of the last 5 months
budgie transaction report year 2 # transaction report of the last 2 years
```

## Contributions

1. Fork the repo
2. Clone the fork (`git clone git@github.com:YOUR_USERNAME/budgie.git && cd budgie`)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Make changes and add them (`git add --all`)
5. Commit your changes (`git commit -m 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create a pull request

## License

See [LICENSE](https://github.com/ykdundar/budgie/blob/master/LICENSE).
