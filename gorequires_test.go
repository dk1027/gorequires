package gorequires

import (
	"os"
	"testing"
	"github.com/dk1027/gotesttools"
)

func TestEnv(t *testing.T) {
	os.Setenv("myvar", "123")
	val := Env("myvar")
	gotesttools.AssertTrue(t, val == "123")
}

func TestEnv2(t *testing.T) {
	doAssert := gotesttools.AssertPanic(t)
	defer doAssert()
	val := Env("myvar2")
	gotesttools.AssertTrue(t, val == "123")
}

func TestFile(t *testing.T) {
	filename := "gorequires_test.go"
	fi := File(filename)
	gotesttools.AssertTrue(t, filename == fi.Name())
}

func TestFile2(t *testing.T) {
	filename := "../gorequires/gorequires_test.go"
	fi := File(filename)
	gotesttools.AssertTrue(t, "gorequires_test.go" == fi.Name())
}

func TestFile3(t *testing.T) {
	doAssert := gotesttools.AssertPanic(t)
	defer doAssert()
	filename := "does-not-exist"
	fi := File(filename)
	gotesttools.AssertTrue(t, filename == fi.Name())
}
