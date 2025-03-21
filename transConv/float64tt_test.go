package transconv

import (
	"fmt"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	aa := strconv.FormatFloat(3.1415926666, 'G', 3, 64)

	fmt.Println(aa)
}
