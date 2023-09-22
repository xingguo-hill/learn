package x

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
)

func cmdNoOutput(ip string) error {
	cmd := exec.Command("ping", ip, "-c", "1", "-t", "1")
	return cmd.Run()
}
func cmdOutput() bool {
	cmd := exec.Command("go", "version")
	out, _ := cmd.Output()
	return strings.Contains(cast.ToString(out), "version")
}
func TestCmd(t *testing.T) {
	assert.Equal(t, nil, cmdNoOutput("180.76.76.76"))
	assert.Equal(t, true, cmdOutput())
}
