package main
import "fmt"
const NMAX int = 100
func main(){
	var A[NMAX]int
	var n, i, j, jenis, makan, k int

	fmt.Scan(&n)
	for i = 0; i < n; i++{
		fmt.Scan(&A[i])
	}
	for i = 0; i < n; i++{
		for j = i+1; j < n; j++{
			if A[i] != A[j]{
				jenis++
			}
		}
		k++
	}
	if jenis == 0{
		jenis = 1
	}
	makan = n/2
	if jenis < makan{
		fmt.Print(jenis)
	}else{
		fmt.Print(makan)
	}
}