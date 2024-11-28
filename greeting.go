package main

import "fmt"

// 構造体の定義
type Student struct {
	name          string
	math, english float64
}

// 平均値を計算するメソッド
func (s Student) avg() float64 {
	return (s.math + s.english) / 2
}

func main() {
	var name string
	var math, english float64

	// ユーザーに入力を促す
	fmt.Print("名前を入力してください")
	fmt.Scanln(&name)

	fmt.Print("数学の点数を入力してください")
	fmt.Scanln(&math)

	fmt.Print("英語の点数を入力してください")
	fmt.Scanln(&english)

	// 構造体を初期化
	student := Student{name, math, english}

	fmt.Println("%sさんの平均点数は%.2fです\n", student.name, student.avg())

}
