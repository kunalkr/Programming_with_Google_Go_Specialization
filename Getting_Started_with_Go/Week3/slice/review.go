package main
import (
	"fmt"
	"sort"
)
import "strconv"
import "strings"
func main() {
	sli := make([]int,3)
	var s string
	var i int
	var x int = 0
	for {
		fmt.Println("Enter a number or X to exit")
		fmt.Scan(&s)
		strings.ToLower(s)
		if s == "x" {
			break
		}
		i, _ = strconv.Atoi(s)
		if x < len(sli) {
			sli[0] = i
			x++
		} else {
				sli = append(sli, i)
				x++
		}
		sort.Ints(sli)
		fmt.Println(sli)
	}
}
