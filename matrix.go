package matrixprep

import "strconv"

/*
Matrix is used to create matrix based on row and column
which has been provided. The matrix is a multidimensional
slice with int8 type.
*/
func Matrix(row int8, col int8) [][]int8 {
	matrix := make([][]int8, row)

	for i := range matrix {
		matrix[i] = make([]int8, col)
	}

	return matrix
}

/*
Visual is used to visualize matrix with return type of string.
*/
func Visual(matrix [][]int8) string {
	var visualMatrix string

	for i, v := range matrix {
		visualMatrix += "| "
		for _, subv := range v {
			visualMatrix += strconv.Itoa(int(subv))
			visualMatrix += " "
		}

		visualMatrix += "|"
		if i < len(matrix)-1 {
			visualMatrix += "\n"
		}
	}

	return visualMatrix
}
