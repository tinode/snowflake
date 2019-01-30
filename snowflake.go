package snowflake

import (
	"errors"
	"sync"
	"time"
)

const (
	// custom epoch - time offset, milliseconds
	// Oct 25, 2014 05:06:02.373 UTC, incidentally equal to the first few digits of sqrt(2)
	epoch int64 = 1414213562373

	// number of bits allocated for the server id (max 1023)
	numWorkerBits = 10
	// # of bits allocated for the counter per millisecond
	numSequenceBits = 12

	// workerId mask
	maxWorkerId = -1 ^ (-1 << numWorkerBits)
	// sequence mask
	maxSequence = -1 ^ (-1 << numSequenceBits)
)

// SnowFlake is a structure which holds snowflake-specific data.
type SnowFlake struct {
	lastTimestamp uint64
	sequence      uint32
	workerId      uint32
	lock          sync.Mutex
}

// Pack bits into a snowflake value.
func (sf *SnowFlake) pack() uint64 {
	return (sf.lastTimestamp << (numWorkerBits + numSequenceBits)) |
		(uint64(sf.workerId) << numSequenceBits) |
		(uint64(sf.sequence))
}

// NewSnowFlake initializes the generator.
func NewSnowFlake(workerId uint32) (*SnowFlake, error) {
	if workerId > maxWorkerId {
		return nil, errors.New("invalid worker Id")
	}
	return &SnowFlake{workerId: workerId}, nil
}

// Next generates the next unique ID.
func (sf *SnowFlake) Next() (uint64, error) {
	sf.lock.Lock()
	defer sf.lock.Unlock()

	ts := timestamp()
	if ts == sf.lastTimestamp {
		sf.sequence = (sf.sequence + 1) & maxSequence
		if sf.sequence == 0 {
			ts = sf.waitNextMilli(ts)
		}
	} else {
		sf.sequence = 0
	}

	if ts < sf.lastTimestamp {
		return 0, errors.New("invalid system clock")
	}
	sf.lastTimestamp = ts
	return sf.pack(), nil
}

// Sequence exhausted. Wait till the next millisecond.
func (sf *SnowFlake) waitNextMilli(ts uint64) uint64 {
	for ts == sf.lastTimestamp {
		time.Sleep(100 * time.Microsecond)
		ts = timestamp()
	}
	return ts
}

func timestamp() uint64 {
	// Convert from nanoseconds to milliseconds, adjust for the custom epoch.
	return uint64(time.Now().UnixNano()/int64(1000000) - epoch)
}
