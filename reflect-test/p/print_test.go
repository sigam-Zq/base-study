package p

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	a := 3.14

	t.Logf("%v", reflect.TypeOf(a))
	t.Logf("%v", reflect.ValueOf(a))

}
