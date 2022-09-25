package snowflake

// Id generates a unique positive number.
func Id() int64 {
	return gen.Id()
}

// UId wraps Id() as uint64.
func UId() uint64 {
	return uint64(Id())
}
