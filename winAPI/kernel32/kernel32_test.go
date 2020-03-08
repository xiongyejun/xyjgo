package kernel32

import (
	"testing"
)

func Test_shortPath(t *testing.T) {
	t.Log(GetPIDByProcessName("exe32.exe"))
}
