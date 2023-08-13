package time

import (
	"fmt"
	"testing"
)

func TestTimeNow(t *testing.T) {
	println(UTCNow())
	println(UTCP8Now())

	time, _ := Parse(UTCP8Now())
	fmt.Println(time)
}
