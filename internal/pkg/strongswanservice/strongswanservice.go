package strongswanservice

import (
	"os"
	"os/exec"
)

func Restart() error {
	cmd := exec.Command("systemctl", "restart", "strongswan-starter.service")
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
