package time

import (
	"testing"
)

func TestTimeNow(t *testing.T) {
	println(UTCNow())
	println(UTCP8Now())
}
