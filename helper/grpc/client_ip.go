package grpc

import (
	"context"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc/peer"
)

// ClientIP extract client ip from context
func ClientIP(ctx context.Context) (string, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("could not get peer from context")
	}

	addr := p.Addr.String()
	ip := addr

	if host, _, err := net.SplitHostPort(addr); err == nil {
		ip = host
	}

	ip = strings.TrimSpace(ip)
	return ip, nil
}
