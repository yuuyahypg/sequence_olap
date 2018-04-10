package main

import (
	"fmt"
	//mmr "github.com/yuuyahypg/sequence_olap/multi_match_recognize"
)

type Test struct {
	run interface{}
}

func main() {
	a := Test{}

	if a.run == nil {
		fmt.Println(a.run)
	}
}
