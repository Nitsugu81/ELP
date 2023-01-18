package main
import "fmt"
import "math/rand"
import (
    "fmt"
    "io/ioutil"
)


/*func main() {
	var m,n,o int
    fmt.Println("Rentrez le nombre de lignes de A ")
	fmt.Scanln(&m)

	fmt.Println("Rentrez le nombre de colonnes de A et ligne de B")
	fmt.Scanln(&n)

	fmt.Println("Rentrez le nombre de colonnes de B ")
	fmt.Scanln(&o)

	var A [m][n] int
	var B [n][o] int


}*/

func main() {
	// Get the size of the matrices from the user
	fmt.Print("Enter the number of rows in the first matrix: ")
	var m int
	fmt.Scan(&m)
	fmt.Print("Enter the number of columns in the first matrix: ")
	var n int
	fmt.Scan(&n)
  
	// Define the first matrix
	var a = make([][]int, m)
	for i := range a {
	  a[i] = make([]int, n)
	}
  
	// Populate the first matrix with values
	for i := 0; i < m; i++ {
	  for j := 0; j < n; j++ {
		var alo int = rand.Intn(100)
		a[i][j] = alo
		fmt.Print(a[i][j])
		fmt.Print(" ")
	  }
	  fmt.Print("\n")
	}
  
	// Get the size of the second matrix from the user
	fmt.Print("\nEnter the number of rows in the second matrix: ")
	var q int
	fmt.Scan(&q)
	fmt.Print("Enter the number of columns in the second matrix: ")
	var p int
	fmt.Scan(&p)
  
	// Define the second matrix
	var b = make([][]int, q)
	for i := range b {
	  b[i] = make([]int, p)
	}
  
	// Populate the second matrix with values
	for i := 0; i < q; i++ {
		for j := 0; j < p; j++ {
		  var alo int = rand.Intn(100)
		  b[i][j] = alo
		  fmt.Print(b[i][j])
		  fmt.Print(" ")
		}
		fmt.Print("\n")
	  }
  
	// Check if the matrices can be multiplied
	if n != q {
	  fmt.Println("Error: the matrices cannot be multiplied")
	  return
	}
  
	// Create a new matrix to store the result
	var result = make([][]int, m)
	for i := range result {
	  result[i] = make([]int, p)
	}
  
	// Perform the multiplication
	for i := 0; i < m; i++ {
	  for j := 0; j < p; j++ {
		for k := 0; k < n; k++ {
		  result[i][j] += a[i][k] * b[k][j]
		}
	  }
	}
  
	// Print

	fmt.Println("\nResult:")
  for i := 0; i < m; i++ {
    for j := 0; j < p; j++ {
      fmt.Print(result[i][j], " ")
    }
    fmt.Println()
  }
}