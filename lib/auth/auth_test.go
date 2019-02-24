package auth

import (
	"testing"
)

func TestAuth(t *testing.T) {

	tok := HashEncode("angryMonkey")
	if tok != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Fatalf("Incorrect password")
	}
	t.Log(tok)
}
