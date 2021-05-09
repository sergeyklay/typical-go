package typgen_test

type sampleInterface interface {
	sampleMethod()
}

// sampleStruct
// @tag1
// @tag2 (key1:"", key2: "", key3:"value3")
type sampleStruct struct {
	sampleInt    int    `default:"value1"`
	sampleString string `default:"value2"`
}

func sampleFunction(param1 int, param2 int) {
	// intentionally blank
}

// GetWriter to get writer to greet the world
// @ctor
func sampleFunction2() {
	// intentionally blank
}

type (
	// @tag3
	sampleInterface2 interface {
		sampleMethod2()
	}

	// sampleStruct2 asdf
	// @tag4
	sampleStruct2 struct {
	}
)

type sampleStruct3 struct {
	Name    string
	Address string
}

func (s *sampleStruct3) someMethod() {
}

func (s sampleStruct3) someMethod2() {
}
