package profiler

import (
	"fmt"
	"log"
	"time"
)

//////////
////////// channel profiler
//////////

type ChannelProfiler struct {
	running  bool
	channels []*ChannelData
}

func NewChannelProfiler(
	channels ...*ChannelData,
) *ChannelProfiler {
	self := &ChannelProfiler{
		running:  false,
		channels: channels,
	}

	go self.checkBlockings()

	return self
}

func (self *ChannelProfiler) checkBlockings() {
	if self.running {
		panic("already started")
	}

	self.running = true

	for self.running {
		for _, channel := range self.channels {
			channel.samples += 1

			length := channel.getLength()

			if length == 0 {
				channel.countEmpty += 1
			}

			if length == channel.capacity {
				channel.capacity += 1
			}
		}

		// TODO: this needs to become a constant
		time.Sleep(time.Millisecond * 100)
	}
}

func (self *ChannelProfiler) ShowResults() {
	self.running = false

	time.Sleep(time.Millisecond * 100 * 3) // TODO: this is shit

	fmt.Printf("Channels:\n")

	for _, channel := range self.channels {
		fmt.Printf("    %v:\n", channel.name)
		fmt.Printf("        Empty: %6.2f%% | %3v / %3v\n", 100*float32(channel.countEmpty)/float32(channel.samples), channel.countEmpty, channel.samples)
		fmt.Printf("        Full : %6.2f%% | %3v / %3v\n", 100*float32(channel.countFull)/float32(channel.samples), channel.countFull, channel.samples)
	}
}

//////////
////////// channel data
//////////

type ChannelData struct {
	name string

	getLength func() int
	capacity  int

	samples    int
	countEmpty int
	countFull  int
}

func NewChannelData(name string, getLength func() int, capacity int) *ChannelData {
	if capacity <= 0 {
		log.Printf("Channel profiler will not work correctly with capacity <=0 (in this case %v)", capacity)
	}

	return &ChannelData{
		name: name,

		getLength: getLength,
		capacity:  capacity,

		samples:    0,
		countEmpty: 0,
		countFull:  0,
	}
}
