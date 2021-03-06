package test

//import "fmt"

func add(x, y int) int{
  return x + y
}

func subtract(x, y int) int{
  return x - y
}

func divide(x, y int) float64{
  x1 := float64(x)
  y1 := float64(y)
  return x1/y1
}

func multiply(x, y int) int{
  return x*y
}

func power(x, y int) int{
  x1 := x
  for i := 1; i < y; i++{
      x = x * x1
  }
  return x
}

/*
func main() {
  a := 5
  b := 3
  fmt.Println(add(a,b))
  fmt.Println(subtract(a, b))
  fmt.Println(divide(a,b))
  fmt.Println(multiply(a,b))
  fmt.Println(power(a,b))
  var input string
  fmt.Scanln(&input)
  fmt.Print(input)


}
*/
