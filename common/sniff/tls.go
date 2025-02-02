package sniff

import (
	"context"
	"crypto/tls"
	"io"

	"github.com/sagernet/sing-box/adapter"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing/common/bufio"
)

func TLSClientHello(ctx context.Context, reader io.Reader, sniffdata chan SniffData) {
	data := SniffData{
		metadata: nil,
		err:      nil,
	}
	defer func() {
		sniffdata <- data
	}()
	var clientHello *tls.ClientHelloInfo
	err := tls.Server(bufio.NewReadOnlyConn(reader), &tls.Config{
		GetConfigForClient: func(argHello *tls.ClientHelloInfo) (*tls.Config, error) {
			clientHello = argHello
			return nil, nil
		},
	}).HandshakeContext(ctx)
	if clientHello != nil {
		data.metadata = &adapter.InboundContext{Protocol: C.ProtocolTLS, Domain: clientHello.ServerName}
		return
	}
	data.err = err
}
