package shadows

import (
	"github.com/shadowsocks/go-shadowsocks2/core"
	//"github.com/shadowsocks/go-shadowsocks2/socks"

	"log"

	"net/url"
	//	"os"

	//"os/signal"
	"strings"
	//"syscall"
	"time"
)

var config struct {
	Verbose    bool
	UDPTimeout time.Duration
	TCPCork    bool
}

func Connect(port string) {

	var key []byte
	var cipher, password string
	addr := "ss://AEAD_CHACHA20_POLY1305:password@:" + port
	var err error
	//	config.Verbose = true
	if strings.HasPrefix(addr, "ss://") {
		addr, cipher, password, err = parseURL(addr)
		if err != nil {
			log.Fatal(err)
		}
	}

	ciph, err := core.PickCipher(cipher, key, password)
	if err != nil {
		log.Fatal(err)
	}
	go tcpRemote(addr, ciph.StreamConn)

	//	sigCh := make(chan os.Signal, 1)
	//signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	//<-sigCh
	//killPlugin()

}

func parseURL(s string) (addr, cipher, password string, err error) {
	u, err := url.Parse(s)
	if err != nil {
		return
	}

	addr = u.Host
	if u.User != nil {
		cipher = u.User.Username()
		password, _ = u.User.Password()
	}
	return
}
