package myip

import (
	"testing"
)

func TestMyIP(t *testing.T) {
	t.Log(GetMyIP1())
	t.Log(GetMyIP2())
}

func TestIsIP(t *testing.T) {
	t.Log(isIP("175.11.14.161"))
}
