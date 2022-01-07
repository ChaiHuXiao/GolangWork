package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type MyType struct {
	TCName string `json:"name"`
}

type ServiceType string

const (
	ServiceTypeCluster  = "cluster"
	ServiceTypePort     = "port"
	ServiceTypeLoad     = "load"
	ServiceTypeExternal = "external"
)

func main33() {
	// test_str := "测试string类型是否可变"
	// fmt.Printf("字符串生命时在内存中分配的地址：%p \n", &test_str)

	// f := func(str string) {
	// 	fmt.Printf("字符串作为参数传递后在内存中的地址：%p \n", &str)
	// }
	// f(test_str)

	// // to_test_selecter()
	// httpservice()
}

func to_test_selecter() {

	drawchan := make(chan string)
	calculateChan := make(chan int)
	persistenceChan := make(chan bool)
	dput := func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 1000 * 2)
			str := "stu" + strconv.Itoa(i)
			drawchan <- str
		}
	}

	cput := func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 1000 * 2)
			calculateChan <- i * i
		}
	}

	pput := func() {
		for i := 0; i < 20; i++ {
			time.Sleep(time.Millisecond * 1000 * 2)
			if rand.Intn(100)%2 > 0 {
				persistenceChan <- true
			} else {
				persistenceChan <- false
			}
		}
	}

	go dput()
	go cput()
	go pput()

	timer := time.NewTimer(time.Millisecond * 1000 * 5)
	for {
		select {
		case v := <-drawchan:
			println("the draw content is：%s", v)
		case v := <-calculateChan:
			fmt.Printf("the calculate result is：%d \n", v)
		case v := <-persistenceChan:
			if v {
				println("Data persistence succeeded")
			} else {
				println("the Data is lost")
			}
		case <-timer.C:
			println("timeout waiting from channel")
		default:
			println("No message was received")
		}
	}
	time.Sleep(time.Minute * 5)
}

func to_test_channel_covert() {

	producers := func(c chan<- int) {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Millisecond * 1000)
			fmt.Printf("api is pushing,the no is：%d \n", i)
			c <- i
		}
		defer close(c)
	}

	consumers := func(c <-chan int) {
		for v := range c {
			fmt.Printf("the api no is ：%d \n", v)
		}
		println("the channel is closed!")
	}

	ServiceApi := make(chan int)

	go producers(ServiceApi)
	go consumers(ServiceApi)

	time.Sleep(time.Millisecond * 1000 * 10)
}

func to_test_enum() {
	st := new(ServiceType)
	println(st)
}

func to_test_reflection() {
	mt := MyType{TCName: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	fmt.Printf("%T \n", name.Name)
	tag := name.Tag.Get("json")
	println(tag)
	// reflect.ValueOf()

}

func to_test_array() {
	// var times [5][0]int
	// print(times)
	// for range times {
	// 	fmt.Println("hello")
	// }

	a := []int{}
	b := []int{1, 2, 3}
	c := a
	fmt.Printf("a address is %p \n", &a)
	fmt.Printf("b address is %p \n", &b)
	fmt.Printf("c address is %p \n", &c)
	a = append(b, 1)
	fmt.Printf("c address is %p \n", &c)
}

func to_test_string() {
	s := "Hellow Word!"
	c := []byte(s)
	c[0] = 'P'
	s2 := string(c)
	fmt.Printf("s is %s\n", s)
	fmt.Printf("c is %s\n", c)
	fmt.Printf("s2 is %s\n", s2)

}

func to_test_array_replace() {
	my_arry := [5]string{"I", "am", "stupid", "and", "weak"}
	fmt.Printf("%v \n", my_arry)
	for i, v := range my_arry {
		switch v {
		case "stupid":
			my_arry[i] = "smart"
		case "weak":
			my_arry[i] = "strong"
		}
	}
	fmt.Printf("%v", my_arry)
}

func to_test_goroutine() {
	input := make(chan<- int, 10)
	output := make(<-chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("Prepared in the %d \n", i)
			input <- i
		}
	}()

	// time_out := make(chan bool)

	// go func() {
	// 	for tmp := range input {

	// 		select{
	// 		case:
	// 		}
	// 	}
	// }()

	go func() {
		res := <-output
		fmt.Printf("get %d", res)
	}()

	time.Sleep(10000 * time.Millisecond * 15)
}
