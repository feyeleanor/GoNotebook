package main
import "bufio"
import . "fmt"
import "net"

var CRLF = ([]byte)("\n")

func main() {
  if address, e := net.ResolveUDPAddr("udp", ":1024"); e == nil {
    if server, e := net.DialUDP("udp", nil, address); e == nil {
      defer server.Close()
      if _, e = server.Write(CRLF); e == nil {
        if text, e := bufio.NewReader(server).ReadString('\n'); e == nil {
          Printf("%v", text)
        }
      }
    }
  }
}
