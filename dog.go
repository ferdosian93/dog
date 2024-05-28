package dog

import (
	"strings"
)

func WhenGrownUp(s string) string {
	return "when the puppy grows up it says:" + strings.ToUpper(s)
}

func WhenGetHungry(s string) string {
	return "when he get hungry says:" + strings.ToUpper(s)
}
