package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS  *ChopS
	rightCS *ChopS
	number  int //
}

var wg sync.WaitGroup

func main() {

	var sem = make(chan int, 2) // semaphore size 2 (punkt 4,5)

	CSticks := make([]*ChopS, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1} //basic order, no lowest-numbered first, philo number added to the struct (punkt 1,3,6 )
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		sem <- 1 // append to semaphore before exec, will lock if the buffer(size 2) is full (punkt 4,5)
		go philos[i].eat()
		<-sem //release one place in a buffer (execution control by host to allow the next cycle to start punkt 4,5)
	}

	wg.Wait()

}

func (p Philo) eat() {
	for i := 0; i < 3; i++ { //eat 3 times punkt 2
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %v \n", p.number)  //punkt 7
		fmt.Printf("finishing eating %v \n", p.number) //punkt 8

		p.rightCS.Unlock()
		p.leftCS.Unlock()

	}
	wg.Done()
}
