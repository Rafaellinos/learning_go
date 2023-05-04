package main

/*
working with constants

const (
	Online = 0
	Offline = 1
	Maintenance = 2
	Retired = 3
)

const ( // short-form, log-form use aiot for every const
	Online = iota // 0 auto increment next values
	Offline // 1
	Maintenance // 2
	Retired // 3
)

const (
	i10 = iota + 10 // starts at 10
	i11
	i12
)

use "-" to skip values, e.g.:

const (
	s0 = iota // 0
	-
	-
	s3
)

*/

///////////////

/*
iota Enumeration pattern

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
		case North: // will be replaced by the byte for the compiler
			return "North"
		case South:
			return "South"
		...
		default:
			return "other direction"
	}
}

north := North
fmt.Println(north)
// prints "North", string here acts like __str__ or __repr__

// shorted way to create String representation of an enumeration:

type Direction byte

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return []string{"North", "East", "South", "West"}[d]
}


*/
