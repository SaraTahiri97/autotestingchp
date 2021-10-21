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
	if len(phrases) > 1 {
		res := strings.Join(phrases[:len(phrases)-1], ",") // phrases[:len(phrases)-1] : gives all but the last one
		if len(phrases) == 2 {
			res += " and"
		} else {
			res += ", and "
		} // add the and

		res += phrases[len(phrases)-1] // gives the last one
		return res
	} else {
		return phrases[0]
	}
}
