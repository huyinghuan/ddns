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
}
