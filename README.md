# Snowflake

Yet another Golang implementation of Twitter's Snowflake. This implementation is used by https://github.com/tinode/chat and as such it's up to date and supported.

## Performance

Maximum theoretical performance is limited by the wait time on the sequence number. I.e. minimum time for a value to be generated is 1 ms / 4096 ~ 244 ns.
Actual performance on average hardware is 246 ns.

## Spec

ID is a 64 bit unsigned integer composed of:
- the top bit is always zero for compatibility; for instance, Go's sql implementation requires top bit of uint64 to be 0
- time - 41 bits (millisecond precision with a custom epoch, enough to cover until the year 2083)
- configured machine id - 10 bits - gives us up to 1024 machines
- sequence number - 12 bits - rolls over every 4096 per machine (with protection to avoid rollover in the same ms, and as such it may block for some hundreds of microseconds)

Differences from Twitter's Snowflake:
- uint64 instead of int64
- zero on error instead of -1
- different epoc: 2014 instead of 2010

## License

Apache License 2.0

## Links

- https://github.com/sdming/gosnow
