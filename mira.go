package main

import "fmt"

// 構造体
type Student struct {
	name          string
	math, english int
}

func (s Student) avg() float64 {
	return float64(s.math+s.english) / 2
}

func main() {
	a001 := Student{"miya", 80, 70}
	average := a001.avg()
	fmt.Printf("%sさんの平均点は%.2fです\n", a001.name, average)
}
