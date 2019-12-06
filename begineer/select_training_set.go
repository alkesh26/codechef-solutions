package main
import "fmt"

type counts struct {
	zeroCount int
	oneCount int
}

func main(){
	var numberOfTests, size, count int
	var str string
	fmt.Scanf("%d", &numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		fmt.Scanf("%d", &size)

		m := make(map[string]counts, size)
		for j := 0; j < size; j++ {
			fmt.Scanf("%s %d", &str, &count)
			if count == 1 {
				x := m[str]
				x.oneCount++
				m[str] = x
			} else {
				x := m[str]
				x.zeroCount++
				m[str] = x
			}
		}

		maxCount := 0
		for _, v := range m {
			if v.zeroCount >= v.oneCount {
				maxCount += v.zeroCount
			} else {
				maxCount += v.oneCount
			}
		}

		fmt.Println(maxCount)
	}
}
