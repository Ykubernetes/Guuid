package utils

import (
	"crypto/md5"
	"os"
	"strconv"
	"strings"
)

func getPid() int {
	return int(os.Getpid())
}

func getHostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "localhost"
	}
	return hostname
}

func machineHash() (machHash []byte) {
	pid := getPid()
	hostname := getHostName()
	randId := myrand()
	machine := strings.Join([]string{hostname, strconv.Itoa(pid), uint32ToHexString(randId)}, "")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(machine))
	machHash = md5Ctx.Sum(nil)
	return
}

var MachineId []byte = machineHash()
