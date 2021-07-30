package conn

import (
	"bufio"
	"fmt"
	"github.com/cebilon123/ElytraGo/packet"
	"net"
)

// HandleConnection handles connection client->server and vice versa
func HandleConnection(c net.Conn, wd *WorkerDispatcher) {
	defer c.Close()
	fmt.Printf("\nServing: %s\n", c.RemoteAddr().String())

	for {
		length := make([]byte, 1)
		reader := bufio.NewReader(c)
		_, err := reader.Read(length)
		if err != nil {
			fmt.Println(err)
			return
		}

		pctBytes := make([]byte, length[0])
		_, err = reader.Read(pctBytes)
		if err != nil {
			fmt.Println(err)
			return
		}

		//Create packet based on bytes
		pct := packet.NewPacket(pctBytes, true)
		wd.ClientPackets <- pct
	}
}
