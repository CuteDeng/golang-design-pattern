package option

import "testing"

//TestType1 test get hiapi with factory
func TestOption(t *testing.T) {
	options := NewOptions(SetIntOption1(100), SetStrOption1("hello"))
	if options.strOption1 != "hello" || options.intOption1 != 100 {
		t.Fatal("newOptions test fail")
	}
}
