package exercises

import (
	"log"
)

/////////////////////
// go routines///////
////////////////////

func helper(c chan int) {
	log.Printf("test start:")
	var i int
	for i = 0; i < 10; i++ {
		c <- i
		log.Printf("test chan: %v", i)
	}
	log.Println("close ()) ")
	close(c)
	log.Println("close ()) ")
	//time.Sleep(time.Second * 2)

}

func Exercise8() {
	channel := make(chan int, 2)
	log.Println("exercise start ")
	go helper(channel)
	data := <-channel

	log.Println("test2 ", data)
	//fmt.Sprintf()
}
