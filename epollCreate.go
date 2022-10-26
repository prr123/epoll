// epoll experiment
//

package main

import (
	"fmt"
	"os"
	syscall "golang.org/x/sys/unix"
)

func main() {

	flag := syscall.EPOLL_CLOEXEC
//	flag = 0
	epfd, err := syscall.EpollCreate1(flag)
	if err != nil {
		fmt.Printf("epollCreate1 error: %v\n", err)
		os.Exit(-1)
	}

	fmt.Printf("epoll file desc: %d\n", epfd)

	err = syscall.Close(epfd)
	if err != nil {
		fmt.Printf("epoll Close error: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("*** success ***\n")
}
