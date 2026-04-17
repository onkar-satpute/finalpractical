// Write a program in GO language to accept two matrices and display its multiplication.

package main

import "fmt"

func main() {
	var r1, c1, r2, c2 int

	fmt.Print("Enter rows and columns of first matrix: ")
	fmt.Scan(&r1, &c1)

	fmt.Print("Enter rows and columns of second matrix: ")
	fmt.Scan(&r2, &c2)

	// Condition for multiplication
	if c1 != r2 {
		fmt.Println("Matrix multiplication not possible (columns of first must equal rows of second).")
		return
	}

	// Declare matrices
	A := make([][]int, r1)
	B := make([][]int, r2)
	C := make([][]int, r1)

	for i := 0; i < r1; i++ {
		A[i] = make([]int, c1)
		C[i] = make([]int, c2)
	}

	for i := 0; i < r2; i++ {
		B[i] = make([]int, c2)
	}

	// Accept first matrix
	fmt.Println("\nEnter elements of first matrix:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c1; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	// Accept second matrix
	fmt.Println("\nEnter elements of second matrix:")
	for i := 0; i < r2; i++ {
		for j := 0; j < c2; j++ {
			fmt.Scan(&B[i][j])
		}
	}

	// Matrix multiplication
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			C[i][j] = 0
			for k := 0; k < c1; k++ {
				C[i][j] = C[i][j] + (A[i][k] * B[k][j])
			}
		}
	}

	// Display result matrix
	fmt.Println("\nMultiplication of two matrices is:")
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			fmt.Print(C[i][j], "\t")
		}
		fmt.Println()
	}
}

//Write a program in GO language to copy all elements of one array into another using a method.

package main

import "fmt"

type Array []int

func (a Array) copyTo(b Array) {
	for i, v := range a {
		b[i] = v
	}
}

func main() {
	var n int
	fmt.Println("Enter size:")
	fmt.Scan(&n)

	a := make(Array, n)

	fmt.Println("Enter elements:")
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	b := make(Array, len(a))

	a.copyTo(b)

	fmt.Println("Array a :", a)
	fmt.Println("Array b :", b)
}
