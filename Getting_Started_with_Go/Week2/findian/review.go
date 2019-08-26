package main
import "fmt"
import "strings"
func main() {
  var s string
  var pre,last,contain bool
	fmt.Scanf("%s",&s)
  s=strings.ToLower(s)
  pre=strings.HasPrefix(s,"i")
  last=strings.HasSuffix(s,"n")
  contain=strings.Contains(s,"a")
  if pre && last && contain{
    fmt.Print("Found!")
  } else{
    fmt.Print("Not Found!")
  }

}