package testdata

type Struct1 struct {
	one   int
	two   string
	three bool
}

type ParentStruct struct {
	ChildStruct
}

type ChildStruct struct {
	name string
}
