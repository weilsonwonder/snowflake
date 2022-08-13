package snowflake

/**
Snowflake ID is based on Twitter's implementation for generation integer ids used for distributed systems.
It is a 64-bit integer, partitioned into different sections, broken down below:

bits used : description
1: not used. to ensure only positive numbers, it will be fixed to 0
42: timestamp. will set to current time in milliseconds, allowing use up to year 2109.
10: system generator id. In a distributed system, each system is a machine. This accommodates up to 1024 machines.
11: sequence number. starts from 0 every new millisecond, up to 2048 unique sequence.

based on this algorithm, the limitation of generating unique sequence is 2048 unique ID per millisecond per machine.
In a distributed system where we max out number of machines or systems (like 1024 nodes),
the max limitation can accommodate up to approximately 2 million unique ID per millisecond.
*/

const (
	timestampBits   = int64(42)
	generatorIdBits = int64(10)
	sequenceBits    = int64(11)

	maxGeneratorId = int64(-1) ^ (int64(-1) << generatorIdBits)
	maxSequence    = int64(-1) ^ (int64(-1) << sequenceBits)

	generatorShiftLeft = sequenceBits
	timestampShiftLeft = generatorIdBits + sequenceBits

	baseepoch = int64(1550836800000) // constant timestamp offset (milliseconds)
)

// default generator
var gen = new(generator)

// Init sets the generatorId. It should be a unique number across distributed systems.
func Init(generatorId int64) {
	gen.Lock()
	gen.GeneratorId = generatorId & maxGeneratorId
	gen.Unlock()
}
