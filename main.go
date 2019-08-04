package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	port := flag.Int("p", 7650, "port of a watchers")
	ip := flag.String("h", "0.0.0.0", "ip of a watchers")
	flag.Parse()

	ln, err := net.Listen("tcp4", *ip+":"+strconv.Itoa(*port))
	if err != nil {
		fmt.Errorf("An error occurred: %s \n\n", err.Error())
		os.Exit(-1)
	}
	defer ln.Close()

	wd, err := os.Getwd()
	if err != nil {
		fmt.Errorf("An error occurred: %s \n\n", err.Error())
		os.Exit(-1)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Errorf("An error occurred: %s \n\n", err.Error())
			os.Exit(-1)
		}
		go touch(wd, conn)
	}
}

func touch(wd string, conn net.Conn) {
	defer conn.Close()
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Errorf("An error occurred: %s \n\n", err.Error())
			}
			return
		}
		path := strings.TrimSpace(string(data))
		if path == "" {
			return
		}
		filePath := wd + path
		mTime, err := getMtime(filePath)
		if err != nil {
			fmt.Errorf("An error occurred getMtime: %s \n\n", err.Error())
			continue
		}
		setMtime(filePath, mTime.Add(time.Second*1))
		if err != nil {
			fmt.Errorf("An error occurred setMtime: %s \n\n", err.Error())
			continue
		}
	}
}

func getMtime(path string) (mtime time.Time, err error) {
	fi, err := os.Stat(path)
	if err != nil {
		return
	}
	mtime = fi.ModTime()
	return
}

func setMtime(path string, mtime time.Time) (err error) {
	atime := time.Now()
	err = os.Chtimes(path, atime, mtime)
	return
}
