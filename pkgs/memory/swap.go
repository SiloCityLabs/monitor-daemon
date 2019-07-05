package memory

import (
	"io/ioutil"
	"os/exec"

	"github.com/capnm/sysinfo"
)

func CheckSwapUsage() {
	si := sysinfo.Get()

	if (si.TotalSwap - si.FreeSwap) >= Settings.Daemon.SwapLimit {
		Service("plexmediaserver", "restart")
		Service("deluged", "restart")
		DropCache()

		si = sysinfo.Get()
		if (si.TotalSwap - si.FreeSwap) >= Settings.Daemon.SwapLimit {
			Telegram("System has High Swap usage over 500MB, Consider Rebooting with command 'reboot', You can ignore this if you are aware that you are running something high memory usage.")
		}
	}
}

func DropCache() {
	exec.Command("sync").Run()
	ioutil.WriteFile("/proc/sys/vm/drop_caches", []byte("3"), 0644)
}
