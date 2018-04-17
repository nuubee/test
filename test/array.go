package main

import (
	"fmt"
)

func main() {
	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j := 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	var a = []string{"10", "a"}

	fmt.Println(len(a))

	var b [1][2]string

	b[0][0] = "tony"
	b[0][1] = "parker"

	fmt.Printf("Element0：[%s] 1：[%s]\n", a[0], a[1])

	fmt.Printf("Element0：[%d] 1：[%d]\n", len(b[0]), len(a))

	//point := "s10"

	_spo7 := "s10"

	fmt.Printf("point address [%d]\n", &_spo7)

	var Book1 Books

	var book2 Books

	Book1.title = "go lang"
	Book1.book_id = 1101

	book2.author = "yuqy"

	fmt.Printf("%s,%d,%s\n", Book1.title, Book1.book_id, book2.author)

	get_point(Book1)

	var s_point *Books

	s_point = &Book1

	s_point.author = "语言结构体指针"

	get_s_point(s_point, Book1)

}

func get_point(b Books) {
	fmt.Printf("struct = %s\n", b.title)
}

type Books struct {
	title   string
	author  string
	book_id int
}

func get_s_point(b *Books, a Books) {
	fmt.Printf("s_point %s -- %s\n", b.author, a.author)
}
