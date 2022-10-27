// epoll experiment
//

package main

import (
	"fmt"
	"os"
	sys "golang.org/x/sys/unix"
)

func main() {
	var event sys.EpollEvent
	var events [10]sys.EpollEvent
	var buf [128]byte

	flag := sys.EPOLL_CLOEXEC
//	flag = 0
	epfd, err := sys.EpollCreate1(flag)
	if err != nil {
		fmt.Printf("epollCreate1 error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("epoll file desc: %d\n", epfd)

//	event:
//  type EpollEvent struct {
//		Events uint32
//		Fd     int32
//		Pad    int32
//  }

	fd := sys.Stdin
	op := sys.EPOLL_CTL_ADD
	event.Events = sys.EPOLLIN
	event.Fd = 0


	err = sys.EpollCtl(epfd,op,fd, &event)
	if err != nil {
		fmt.Printf("epollCtl error: %v\n", err)
		os.Exit(-1)
	}

// replace with endless loop
	timInt := 5000 //millisec
	for i:=0; i<5; i++ {
		fmt.Printf("polling %d\n",i)
		eventCount, err := sys.EpollWait(epfd, events[:], timInt )
		if err != nil {
			fmt.Printf("epollWait error: %v\n", err)
			os.Exit(-1)
		}

		fmt.Printf("event counts: %d \n", eventCount)

		for ev:= 0; ev< eventCount; ev++ {

			n, err := sys.Read(fd, buf[:])
			if err != nil {
				fmt.Printf("Read error: %v\n", err)
				os.Exit(-1)
			}
			fmt.Printf("read %d: %s\n", n, string(buf[:n]))
		} //ev
	}

	err = sys.Close(epfd)
	if err != nil {
		fmt.Printf("epoll Close error: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("*** success ***\n")
}
