package ping

import (
	"fmt"
	"io"
	"net"
	"time"
)

func Ping(host, port string, out io.Writer) {
	done := make(chan struct{})
	ticks := 0
	tick := time.Tick(750 * time.Millisecond)
	timeout := 500 * time.Millisecond
	conn, err := net.DialTimeout("tcp", host+":"+port, timeout)

	for {
		select {
		case <-done:
			return
		case <-tick:
			ticks++
			if ticks == 5 {
				close(done)
			}

			if err != nil {
				_, err := fmt.Fprintf(out, "%s is not available on port %s\n", host, port)
				if err != nil {
					return
				}
			} else {
				_, err := fmt.Fprintf(out, "connected to %s at %s\n", host, conn.RemoteAddr().String())
				if err != nil {
					return
				}
			}

			conn, err = net.DialTimeout("tcp", host+":"+port, timeout)
		}
	}
}
