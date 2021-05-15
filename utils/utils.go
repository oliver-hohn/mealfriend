package utils

import (
	"fmt"
	"regexp"
)

func MustMapGroupToIndex(r *regexp.Regexp) map[string]int {
	ret := map[string]int{}
	for i, grpName := range r.SubexpNames() {
		if grpName == "" {
			// Groups without a name do not need a mapping to an index
			continue
		}

		if ret[grpName] != 0 {
			panic(fmt.Sprintf("more than one group with the same name: %s, regex: %s", grpName, r.String()))
		}

		ret[grpName] = i
	}

	return ret
}
