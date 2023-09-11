package time

import (
	"fmt"
	"testing"
)

func TestTimeNow(t *testing.T) {
	println(UTCNow())
	println(UTCP8Now())
	println(UTCNowAddMinutes(60))
	println(UTCNowAddMinutes(-60))

	time, _ := Parse(UTCP8Now())
	fmt.Println(time)
}
