# Snowflake

Yet another Golang implementation of Twitter's Snowflake.

## Performance

TODO

## Spec

ID is a 64 bit unsigned integer composed of:
- the top bit is always zero
- time: 41 bits (millisecond precision with a custom epoch, enough to cover until the year 2083)
- host id: 10 bits, gives us up to 1024 hosts
- sequence number: 12 bits, rolls over every 4096 per host (with protection to avoid rollover in the same ms, it may block the call for some hundreds of microseconds on very heavy loads)

Differences from Twitter's Snowflake:
- unint64 instead of int64
- zero on error instead of -1
- different epoc: 2014 instead of 2010

## Lisence

Apache License 2.0

## Links

- https://github.com/sdming/gosnow
