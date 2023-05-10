package tests

import (
	"learning_go/basics"
	"testing"
)

/*
Fail() = mark fail
Errorf(msg) = Fail and print msg
FailNow() = mark fail, abort
Fatalf(msg) = Fail, abort print msg
Logf() = Log msg when fail
*/

func TestCalcUtilsBinaryToDecimal(t *testing.T) {
	expectedResult := int64(13) //arrange

	res := basics.BinaryToDecimal("1101") // act

	if res != expectedResult { // assert
		t.Errorf("Binary 1101 %v != %v", res, expectedResult)
	}
}
