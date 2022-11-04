# ohlc-resampler

ohlc-resampler is a simple utility to transform trade data
from [Binance Cryptocurrency market](https://www.binance.com/) data into 
OHLCV candles to simulate market behavior and test trading algorithms.

Binance data archives can be found [here](https://data.binance.vision/?prefix=data/).

## Word of warning!

At the moment this is very raw alpha version and cannot provide any guarantee
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
ohlc-resampler timeframe n < input.csv
```

* timeframe - time frame in string format like: 1m, 5m, 15m, 1h. See Go's
[time#ParseDuration](https://pkg.go.dev/time#ParseDuration).
* n - number of candles, cannot be less than zero, otherwise
fallback 20 will be used.

These two parameters are mandatory.

### Output to stdout

Be aware that stream data can be very large. It can mess with your terminal
and priting can take a lot of time.

```sh
ohlc-resampler 5m 20 < DOGEUSDT-trades-2022-10-26.csv
```

### Redirect output to file
```sh
ohlc-resampler 5m 20 < DOGEUSDT-trades-2022-10-26.csv > output.csv
```

### cURL example

You can combile cURL and other command line utilities to make interesing
combinations like dowloading data from the web.

#### GNU/Linux

```
curl -s --output "-" https://data.binance.vision/data/spot/daily/trades/DOGEUSDT/DOGEUSDT-trades-2022-10-22.zip | gunzip | ohlc-resampler 5m 10 > out.csv
```

#### MacOS

```
curl -s --output "-" https://data.binance.vision/data/spot/daily/trades/DOGEUSDT/DOGEUSDT-trades-2022-10-22.zip | tar -x -O | ohlc-resampler 5m 10 > out.csv
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
go test -coverprofile c.out
go tool cover -html c.out
```

## Contribution

Any help is appreciated. Found a bug, typo, inaccuracy, etc?
Please do not hesitate and make pull request or issue.

## License

MIT 2022
