package matrixprep

import (
	"reflect"
	"testing"
)

func TestMatrix(t *testing.T) {
	matrix := Matrix(3, 3)

	tester := make([][]int8, 3)
	tester[0] = make([]int8, 3)

	if reflect.DeepEqual(matrix, tester) {
		panic("The matrix is not the same")
	}
}

func TestVisual(t *testing.T) {
	matrix := Visual(Matrix(3, 3))
	tester := "| 0 0 0 |\n| 0 0 0 |\n| 0 0 0 |"

	if matrix != tester {
		panic("The visual of the matrix isn't the same")
	}
}
