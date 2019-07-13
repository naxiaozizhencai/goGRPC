package util

import (
	"fmt"
	"github/leel0330/grpcdemo/model"
	"google.golang.org/grpc/resolver"

	"strings"
)

// BuildPrefix has last "/"
func BuildPrefix(info model.ServerNodeInfo) string {
	return fmt.Sprintf("/%s/%s/", info.Name, info.Version)
}
func BuildRegPath(info model.ServerNodeInfo) string {
	return fmt.Sprintf("%s%s",
		BuildPrefix(info), info.Addr)
}

// split addr from reg full path ip:port
func SplitPath(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

// Exist helper function
func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}

// Remove helper function
func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
