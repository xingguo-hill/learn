package mocktest

import (
	xtestmock "learn/x/mocktest/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestMyThing(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockObjB := xtestmock.NewMockMyInterfaceB(mockCtrl)
	mockObjB.EXPECT().SomeMethodB("hello").Return("hello world!")
	b := NewBObj(A{"hello"})
	assert.Equal(t, mockObjB.SomeMethodB("hello"), b.SomeMethodB(" world!"))
}
