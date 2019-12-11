package main
import "fmt"
import "math"

// Solution is not accepted on Codechef
func main(){
	var numberOfTests, n, a, b, c, d, p, q, y int
	for t := 0; t < numberOfTests; t++ {
		fmt.Scanf("%d %d %d %d %d %d %d %d", &n, &a, &b, &c, &d, &p, &q, &y)

		arr := make([]int, n+1)
		for i := 1; i <= n; i++ {
			fmt.Scanf("%d", &arr[i])
		}

		walk := int(math.Abs(float64(arr[b] - arr[a])))*p
		train := int(math.MaxInt64)

		if int(math.Abs(float64(arr[c] - arr[a])))*p <= y {
			train = y + int(math.Abs(float64(arr[d] - arr[b])))*p + int(math.Abs(float64(arr[d] - arr[c])))*q
		}

		fmt.Println(int(math.Min(float64(walk), float64(train))))
	}

}
