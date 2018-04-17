package main

import (
	"fmt"
)

func jc(n uint64) (result uint64) {
	if n > 0 {
		result = n * jc(n-1)
		fmt.Printf("n=%d --- value=%d \n", n, result)
		return result
	}
	return 1
}

func main() {

	jcValue := 15

	fmt.Printf("%d 的阶乘 %d \n", jcValue, jc(uint64(jcValue)))

	fmt.Println("------------ map test ------------- \n")

	var mmap map[string]string
	mmap = make(map[string]string)

	mmap["id"] = "10"
	mmap["name"] = "yuqy"
	mmap["name"] = "hrf"

	for m := range mmap {
		fmt.Printf("name=%s \n", mmap["id"])
		fmt.Printf("key=%s \n", m)
	}

	delete(mmap, "id")

	for m2 := range mmap {
		fmt.Printf("name=%s \n", mmap["id"])
		fmt.Printf("key=%s \n", m2)
	}
	//fmt.Printf(len(mmap), ...)
	//fmt.Println(cap(maap, ...)

	k1, k2 := 1, 2

	fmt.Printf("%d -- %d \n", k1, k2)

	k1, k2 = k2, k1

	fmt.Printf("%d -- %d \n", k1, k2)

	arr := []int{1, 2, 3, 4, 5}

	for k := range arr {
		fmt.Printf("arr[%d]=%d \n", k, arr[k])
	}
	fmt.Printf("len(arr)=%d \n", len(arr))

	//arr[0], arr[1], arr[3], arr[4] = arr[4], arr[3], arr[1], arr[0]

	j := 0
	var arr2 [5]int
	for i := len(arr) - 1; i >= 0; i-- {
		arr2[j] = arr[i]
		fmt.Printf("i=%dj=%d\n", i, j)
		j++
	}

	for k := range arr2 {
		fmt.Printf("arr2[%d]=%d \n", k, arr2[k])
	}

}
