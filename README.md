# Snowflake ID Generator

For distributed systems, based on the idea from [Twitter's Snowflake announcement](https://blog.twitter.com/engineering/en_us/a/2010/announcing-snowflake). 

## TLDR;

- Thread safe
- Generate up to 2048 unique IDs per millisecond, per system 
- Allow up to 1024 systems
- As a distributed system, able to generate up to approximately 2 million unique IDs
- Uniqueness guaranteed until year 2150

## Installation

`go get github.com/weilsonwonder/snowflake`

## Usage

```go
func main() {
	// initialize snowflake to system id.
	// must be unique in a distributed system architecture.
	snowflake.Init(101)
	
	// generate a unique id everytime, throughout your system.
	newId := snowflake.Id()
}
```

## References

- https://blog.twitter.com/engineering/en_us/a/2010/announcing-snowflake
- https://en.wikipedia.org/wiki/Snowflake_ID
