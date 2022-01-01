/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var ()

func doReplay() (x bool) {
	fmt.Printf("Replaying %s\n", replaySkits)
	replayWg := &sync.WaitGroup{}
	replayWg.Add(1)
	// this would rightly fail if not running as root.
	go acquireUinputFd(replayWg)

	f, err := os.Open(replaySkits)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open skits file %s: %v\n", replaySkits, err)
		os.Exit(1)
	}
	defer f.Close()
	rc := false

	fscanner := bufio.NewScanner(f)
	for fscanner.Scan() {
		line := fscanner.Text()
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.TrimSpace(line)
		splat := strings.Split(line, "|")
		if len(splat) != 3 {
			continue
		}
		var x, y, z int
		var err error
		x, err = strconv.Atoi(splat[0])
		if err != nil {
			continue
		}
		y, err = strconv.Atoi(splat[1])
		if err != nil {
			continue
		}
		z, err = strconv.Atoi(splat[2])
		if err != nil {
			continue
		}
		rc = relay(x, y, z, replayWg)
		if !rc {
			fmt.Fprintf(os.Stderr, "Could not send %s, exiting\n", line)
			os.Exit(1)
		}
	}

	if rc {
		relay(0, 0, 0, replayWg)
	}
	destroy()
	return
}

func getFd() int {
	fd, err := syscall.Open(inputDevice, syscall.O_WRONLY|syscall.O_NONBLOCK, 0755)
	if err != nil {
		fmt.Fprintf(os.Stderr, "eventLib init err: %v\n", err)
		return -1
	}
	return fd
}

func sendEvents(a, b, c int, fd int) (x bool) {
	fmt.Println("SendEvents called")
	data := make([]byte, 24)
	binary.LittleEndian.PutUint16(data[16:18], uint16(a))
	binary.LittleEndian.PutUint16(data[18:20], uint16(b))
	binary.LittleEndian.PutUint32(data[20:24], uint32(c))
	n, err := syscall.Write(fd, data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "emit err: %v\n", err)
	}
	fmt.Println("Send events bytes", n)
	return err == nil

}
