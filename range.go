package main
import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		fmt.Print(sum, "+", num, "=")
		sum += num
		fmt.Println(sum)
	}

	fmt.Println("---------")
	fmt.Println("total", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "Apple", "b": "Banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}


}
