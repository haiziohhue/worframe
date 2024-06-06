package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func readMenStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	log.Printf("=>MEM 状态: Alloc = %v MiB | HeapIdle: %d(Bytes) | HeapRelaeased: %d(Bytes)", ms.Alloc/1024, ms.HeapIdle/1024/1024, ms.HeapReleased/1024/1024)
}
func test() {
	container := make([]int, 8)
	log.Println("===> loop begin")
	for i := 0; i < 32*1000*1000; i++ {
		container = append(container, i)
		if i == 16*1000*1000 {
			readMenStats()
		}
	}
	log.Println("<=== loop stop")
}
func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	log.Println("start")
	readMenStats()

	test()
	readMenStats()

	log.Println("force gc")
	runtime.GC()
	readMenStats()

	log.Println("done")
	go func() {
		for {
			readMenStats()
			time.Sleep(10 * time.Second)
		}
	}()
	time.Sleep(3600 * time.Second)
}
