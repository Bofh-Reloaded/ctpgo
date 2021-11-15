package common

import "log"

func Assert(prefix string, err error) {
	if err != nil {
		log.Fatalf("%v Fatal error in assertion %v", prefix, err)
	}
}
