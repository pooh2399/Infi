package main

import "fmt"
import "time"

type booking struct {
	patient  string
	timeslot time.Time
}

type appointment struct {
	book []booking
}

func (a *appointment) book(bookingTime time.Time, patient string) {
	if a.book == nil {
		a.booking = make([]booking, 0)
	}
	a.booking = append(a.booking, booking{timesot: bookingTime, patient: name})
}

func (a *appointment) cancel(bookingTime time.Time, name string) {
	a.booking = append(a.booking, booking{timesot: bookingTime, patient: name})

}
func main() {
	b := new(booking)
	b := new(appointment)
	var err error
	b.patient = "S.E.S"
	b.timeslot, err = time.Parse(time.RFC3339, "2021-02-18T09:00:00Z")

	if err != nil {
		fmt.Println("Error in parsing:%v".err)
	}
	a.Book(b.timeslot, b.patient)
	fmt.Printf

	for _v := range a.booking {
		fmt.Printf("Timeslot %v", v)
		fmt.Printf("Timeslot %v \n", v.timeslot)
	}
	arr := []int{23, 6, 7, 9, 8}
	l=len(arr)
	arr1:=
	arr=append(arr[:2],arr[3:5])
	//Cancellation
	for i, v := range arr a.booking{
		fmt.Printf("Timeslot%d %d", i, v)
		if v.timeslot == bookingTime{
			a.booking=append(a.booking[:i],a.booking[i+1:]...)
			break
		}
	}

	/*fmt.Printf("\n")
	fmt.Printf("Timeslot %v",t)
	fmt.Println("Enter the Number of Patients you want to register for")
	var x int
	fmt.Scanln(&x)
	var pat[100]  patient
	for k := 0; k < x; k++ {
		var name string
		var phone int
		fmt.Println("Enter name and number")
		fmt.Scanln(&name)
		fmt.Scanln(&phone)
	*/
}
