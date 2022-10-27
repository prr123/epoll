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

	err = sys.Close(epfd)
	if err != nil {
		fmt.Printf("epoll Close error: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("*** success ***\n")
}
