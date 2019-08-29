package main

import (
  "fmt" 
  "bufio"
  "strconv"
  "os"
)

func GenDisplaceFn(acc, velocity, displacement float64) func (float64) float64 {
  fn := func (time float64) float64 {
    return acc * time * time / 2 + velocity * time + displacement
  }

  return fn
}


func main() {
	fmt.Printf("This is the displacement program.\n")

	fmt.Printf("Please type accelaration: ")
	scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan() 
  acc, _ := strconv.ParseFloat(scanner.Text(), 64)
  fmt.Println("You typed: ", acc)

  fmt.Printf("Please type initial velocity: ")
  scanner.Scan() 
  velocity, _:= strconv.ParseFloat(scanner.Text(), 64)
  fmt.Println("You typed: ", velocity)

  fmt.Printf("Please type initial displacement: ")
  scanner.Scan() 
  displacement, _ := strconv.ParseFloat(scanner.Text(), 64)
  fmt.Println("You typed: ", displacement)

  fmt.Printf("Please type time: ")
  scanner.Scan() 
  time, _ := strconv.ParseFloat(scanner.Text(), 64)
  fmt.Println("You typed: ", time)

  fn := GenDisplaceFn(acc, velocity, displacement)
  fmt.Println("The final displacement is: ", fn(time))

}