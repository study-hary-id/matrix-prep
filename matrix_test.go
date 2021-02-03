package matrixprep

import (
	"reflect"
	"testing"
)

func TestMatrix(t *testing.T) {
	matrix := Matrix(7, 7)

	tester := make([][]int8, 7)
	tester[0] = make([]int8, 7)

	if reflect.DeepEqual(matrix, tester) {
		panic("The matrix is not the same")
	}
}
