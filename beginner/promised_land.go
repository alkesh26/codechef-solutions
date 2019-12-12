package main
import "fmt"

func main(){
	var numberOfTests, n, m, x, y, s, num int
	fmt.Scanf("%d", &numberOfTests)

	for t := 0; t < numberOfTests; t++ {
		var sumX, sumY, prevX, prevY int
		fmt.Scanf("%d %d", &n, &m)
		fmt.Scanf("%d %d %d", &x, &y, &s)

		for i := 0; i < n; i++ {
			fmt.Scanf("%d", &num)
			sumX += (num-prevX-1)/s
			prevX = num
		}

		sumX += (n - prevX)/s

		for i := 0; i < m; i++ {
			fmt.Scanf("%d", &num)
			sumY += (num-prevY-1)/s
			prevY = num
		}

		sumY += (m - prevY)/s

		fmt.Println(sumX*sumY)
	}
}