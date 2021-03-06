package autotestingchp

import (
	"strings"
)

// Library to Join Multiple string values ( ex: if 2 : apples and oranges . if multipule : apples, oranges and pear)
func JoinWithCommas(phrases []string) string {
	/*
		The code makes use of the strings.Join function,which takes a slice of strings and a string to join
		them all together with. Join returns a single string with all the items from the slice combined,
		with the joining string separating each entry.
	*/
	if len(phrases) == 0 {
		return "" // to avoid the panic err of trying to access elements from an empty slice
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		res := strings.Join(phrases[:len(phrases)-1], ",") // phrases[:len(phrases)-1] : gives all but the last one
		res += ", and "
		res += phrases[len(phrases)-1] // gives the last one
		return res
	}

}
