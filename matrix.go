package matrixprep

/*
Matrix is used to create matrix based on
row and column which has been provided
*/
func Matrix(row int8, col int8) [][]int8 {
	matrix := make([][]int8, row)

	for i := range matrix {
		matrix[i] = make([]int8, col)
	}

	return matrix
}
