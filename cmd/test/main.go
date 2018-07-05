package main

import "time"

func main() {
	println("Mulai fungsi tutup channel")

	test1 := make(chan bool)
	test2 := make(chan bool)

	closedTest1 := make(chan bool)
	closedTest2 := make(chan bool)

	go func() {
		println("fungsi test 1 berjalan")
		for {
			select {
			case <-test1:
				test2 <- true
				closedTest1 <- true
				break
			}
		}
	}()

	go func() {
		println("fungsi test 2 berjalan")
		for {
			select {
			case <-test2:
				test1 <- true
				closedTest2 <- true
				break
			}
		}
	}()

	time.Sleep(1 * time.Second)
	println("close test 1 dimulai")
	test1 <- true
	println("close test 1 selesai")

	<-closedTest1
	println("fungsi test 1 telah ditutup")

	<-closedTest2
	println("fungsi test 2 telah ditutup")

	println("Berhasil")
}
