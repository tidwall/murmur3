package murmur3

import (
	"fmt"
	"testing"
)

func TestMurmur3(t *testing.T) {
	testHash32x86(1234, "Hello, world!", "faf6cdb3")
	testHash32x86(4321, "Hello, world!", "bf505788")
	testHash32x86(1234, "xxxxxxxxxxxxxxxxxxxxxxxxxxxx", "8905ac28")
	testHash32x86(1234, "", "0f2cc00b")
}

func testHash32x86(seed uint32, key string, expect string) {
	v := fmt.Sprintf("%08x", Sum32Seed(key, seed))
	if v != expect {
		panic(fmt.Sprintf("expected '%v', got '%v'", expect, v))
	}
}
