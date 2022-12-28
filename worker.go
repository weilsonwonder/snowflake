package snowflake

import (
	"sync"
	"time"
)

type generator struct {
	sync.Mutex
	LastTimestamp int64 // timestamp of the last ID
	GeneratorId   int64 // system's id, to be unique across distributed systems
	Sequence      int64 // sequence numbers generated in the current millisecond
}

// Id safely generates a unique positive number.
func (z *generator) Id() int64 {
	z.Lock()
	defer z.Unlock()

	ts := time.Now().UnixMilli()
	if ts < z.LastTimestamp {
		// time is moving backwards!
		panic("time is moving backwards!")
	}

	// update sequence
	if ts == z.LastTimestamp {
		z.Sequence = (z.Sequence + 1) & maxSequence

		// check overflow
		if z.Sequence == 0 {
			// wait until next millisecond
			for ts <= z.LastTimestamp {
				ts = time.Now().UnixMilli()
			}
		}
	} else {
		z.Sequence = 0
	}

	// update timestamp
	z.LastTimestamp = ts

	// generate and return id
	return (ts-baseepoch)<<timestampShiftLeft |
		z.GeneratorId<<generatorShiftLeft |
		z.Sequence
}
