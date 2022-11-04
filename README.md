# ohlc-resampler

ohlc-resampler is a simple utility to transform aggregated trade data 
from Binance Cryptocurrency market data into OHLCV candles to simulate
market behavior and test trading algorithms.

## Word of warning!

At the moment this is very raw alpha version and cannot give any guarantee
in data correctness. Also no any further backwards compatibility guaranteed.

## Compilation and installation

There is nothing special required to compile and run this software. 
It has no dependencies and uses only the standard library. Executable 
file will be saved in your `$GOPATH/bin` directory.

```sh
go build
go install
```

That's all.

## Usage

ohlc-resampler reads data from `stdin`, similarly, output will be written
to `stdout`. Errors are written to `stderr`.

```sh
ohlc-resampler timeframe N < input.csv
```

* frame - timeframe in string format like: 1m, 5m, 15m, 1h. See Go's
[time#ParseDuration](https://pkg.go.dev/time#ParseDuration).
* N - number of candles, cannot be less than zero, otherwise
fallback 20 will be used.

### Standard example

```sh
ohlc-resampler 5m 20 < files/DOGEUSDT-aggTrades-2022-10-26.csv
```

### Redirect output to file
```sh
ohlc-resampler 5m 20 < files/DOGEUSDT-aggTrades-2022-10-26.csv > output.csv
```

## Testing

```sh
go test
```

### Coverage

```sh
go test -cover
```

With cover profile.

```sh
go test -coverprofile c.out && go tool cover -html c.out
```

## Contribution

Any help is appreciated. Found a bug, typo, inaccuracy, etc?
Please do not hesitate and make pull request or issue.

## License

MIT 2022
