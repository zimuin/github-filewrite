package fileWrite

import "testing"

func TestFileWrite(t *testing.T) {
    result := Hello("zimuin")
    want := "Hi, zimuin. Welcome!"
    if result != want {
        t.Errorf("fileWrite.Hello() = %q want %q", result, want)
    }
}