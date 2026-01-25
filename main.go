// go mod init mobilebg2
// go mod tidy

// TODO: print progress

package main

import (
	"fmt"
	"log"
	"mobilebg2/define"
	"mobilebg2/extract_cars"
	"mobilebg2/net"
	"mobilebg2/persistentstorage"
	"os"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"time"

	"github.com/kuche1/channelprofiler"
)

func main() {
	if define.ENABLE_PROFILER_TRACE {
		// record every blocking event
		runtime.SetBlockProfileRate(1)

		// optional: mutex contention too
		runtime.SetMutexProfileFraction(1)

		// TODO: make this a define or config
		f, err := os.Create("trace.out")
		if err != nil {
			log.Fatal(err)
		}

		err = trace.Start(f)
		if err != nil {
			log.Fatal(err)
		}

		defer trace.Stop()

		// TODO: analyze with `go tool trace trace.out`
		// then open the URL printed
		// click on `Goroutine analysis` # `Synchronization blocking profile`
		// click on the suspicious one, then look for `Block time (chan send)` and `Block time (chan receive)`
	}

	if define.ENABLE_PROFILER_CPU {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()

		// TODO: view with `go tool pprof -http=:8080 cpu.prof`

		// TODO: if this is empty run `sh -c 'echo 1 > /proc/sys/kernel/perf_event_paranoid'`
		// TODO: this actually mimght be the solution to the profiler above

		// TODO: I am FUCKING RETARDED BECAUSE I PUT THIS IN THE OTHER MAIN
		// now I need to test that it works and then test it without the kernel parameter
	}

	channelProfiler := channelprofiler.NewChannelProfiler()
	channelProfiler.Start()
	// defer channelProfiler.StopAndPrintResults()

	persistentStorage := persistentstorage.NewPersistentStorage(define.PERSISTENT_STORAGE)
	defer persistentStorage.Close()

	chan_net_req := net.RequesterInit(persistentStorage)
	channelProfiler.AddChannels(channelprofiler.NewChannelData(
		"chan_net_req",
		func() int { return len(chan_net_req) },
		cap(chan_net_req),
	))

	chan_cars := extract_cars.Main(channelProfiler, chan_net_req)

	// allCars := make([]*car.Car, 0)

	for car := range chan_cars {
		fmt.Printf("\n%v", car.Sprintf())
		// allCars = append(allCars, car)
	}

	channelProfiler.StopAndPrintResults()

	// sort.Slice(allCars, func(idxA int, idxB int) bool {
	// 	return allCars[idxA].Horsepower < allCars[idxB].Horsepower
	// })

	// fmt.Print("==================================================\n")
	// fmt.Print("==================================================\n")
	// fmt.Print("==================================================\n")
	// fmt.Print("\n")

	// for _, car := range allCars {
	// 	fmt.Printf("%v\n", car.Sprintf())
	// }

	fmt.Print("==================================================\n")
	fmt.Print("All Done\n")

	for {
		time.Sleep(1 * time.Second)
	}
}
