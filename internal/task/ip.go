package task

import (
	"bufio"
	"io"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func InitRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

func isIPv4(ip string) bool {
	return strings.Contains(ip, ".")
}

func loadIPRanges(fileName string) (ips []*net.IPAddr) {
	file, err := os.Open("./data/" + fileName + ".txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		ip, err := reader.ReadString('\n')
		ips = append(ips, &net.IPAddr{IP: net.ParseIP(strings.TrimSpace(ip))})
		if err == io.EOF {
			return
		}
	}
	return
}
