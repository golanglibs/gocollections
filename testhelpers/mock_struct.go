package testhelpers

type MockStruct struct {
	Prop int
}

func NewMockStruct(prop int) *MockStruct {
	return &MockStruct{
		Prop: prop,
	}
}
