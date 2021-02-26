package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello Darkness my little friend")
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(70, items))
	fmt.Println(fibo(5))
	var num, rem, temp int
	var rev int = 0

	fmt.Print("Enter any positive integer : ")
	fmt.Scan(&num)

	temp = num

	//looping for the reverse number declarationm
	for {
		rem = num % 10
		rev = rev*10 + rem
		num /= 10

		if num == 0 {
			break
		}
	}

	if temp == rev {
		fmt.Printf("%d is a Palindrome", temp)
	} else {
		fmt.Printf("%d is not a Palindrome", temp)

	}
}
func binarySearch(x int, arr []int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		med := (low + high) / 2
		if arr[med] < x {
			low = med + 1
		} else {
			high = med - 1
		}
	}

	if low == len(arr) || arr[low] != x {
		return false
	}

	return true
}

func fibo(n int) int {
	var arr [100]int
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-2] + arr[i-1]
	}
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
	return 0
}

/*ints3:=[]int{23,45,46,7}
for key,value :=range ints3{
    fmt.Printf("innt3 array elements %dis %d",key,value)
}
n:=10;
ints3=make([]int, n)
for key, value:=range ints3{
    ints3[i]=i+1
    fmt.Printf("\n Array elements is %d",key value)
}

func divide(x int, y int) float64,error{
    if y==0{
        return nil,errors.new("Dvider is zero")
    }
    return x/y,nil
}
*/
