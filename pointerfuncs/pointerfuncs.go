package pointerfuncs

import "fmt"

//give foo a clone of toot, init variable
func Foo(smell int){
  //change clone of toot, not actually changing toot
  smell++;
  DisplayCool("Within foo, after changing smell got: ", &smell);
  //*smell++;
}

//take pointer of toot, init varaible
func Boo(smell *int){
  //change actual value of toot, the init variable
  *smell++
  DisplayCool("Within boo, after changing smell got: ", smell);
}

//take variable and string, display string and then the value,memory pointer
func DisplayCool(statement string, ack *int){
  //fmt.Println(statement, ack,",",*ack);
  fmt.Printf("%-40s|%12p|%7d|\n",statement,ack,*ack);
}
