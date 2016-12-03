package main

import "fmt"
import "github.com/Nore5515/GO_Test_Repo/pointerfuncs"

func main (){

// initial variable = 10
  var toot int;
  toot = 10;

//pointer to a memory of size int
  var fart *int;
//assign the pointer to the beginning of toot's memory
  fart = &toot;

//pointer to a memory of size int
  var gas *int;
//assign the pointer to the beginning of toot's memory
  gas = &toot;

//print init variable
  pointerfuncs.DisplayCool("Initial variable: ", &toot);
//print out the pointer to the beginning of toot's memory
  pointerfuncs.DisplayCool("Memory value of toot: ", fart);
//print out the memory located at the pointer to the beginning of toot's memory
  pointerfuncs.DisplayCool("Using pointer of toot, get value: ", gas);

  fmt.Println("Testing functions");

//give foo init variable
  pointerfuncs.Foo(toot);
//display init variable
  pointerfuncs.DisplayCool("After running foo, got: ",&toot);

//give boo init variable
  pointerfuncs.Boo(&toot);
//display init variable
  pointerfuncs.DisplayCool("After running boo, got: ",&toot);


}
