package myip

import (
	"testing"
)

func TestMyIP(t *testing.T) {
	t.Log(getMyIP1())
	t.Log(getMyIP2())
}

func TestIsIP(t *testing.T) {
	t.Log(isIP("175.11.14.161"))
	t.Log(isIP("240e:381:e04:2a00:1e69:7aff:fea3:f75c"))
}
