package oauth

import (
	"fmt"
	"testing"
)

func TestGetUserMeta(t *testing.T) {
	meta := GetUserMeta()
	fmt.Println(meta)
}