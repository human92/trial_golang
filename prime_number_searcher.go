package main

import(
	"fmt"
	"time"
)

func main(){
	max := 100
	fmt.Printf("%v 以下の素数:", max)

	start := time.Now() //start time
	for n := 2; n <= max; n++ {
		flag := true
		for m := 2; m < n; m++ {
			if (n % m) == 0 { //nがmで割り切れる⇒素数でない
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf(" %v", n)
		}
	}
	goal := time.Now() //end time
	fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}