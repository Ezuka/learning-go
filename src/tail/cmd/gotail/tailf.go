package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hpcloud/tail"
)

func args2config() (tail.Config, int64) {
	config := tail.Config{Follow: true}
	n := int64(0)
	maxlinesize := int(10)
	flag.Int64Var(&n, "n", 0, "tail from the last Nth location")
	flag.IntVar(&maxlinesize, "max", 0, "max line size")
	flag.BoolVar(&config.ReOpen, "F", false, "follow, and track file rename/rotation")
	flag.BoolVar(&config.Poll, "p", false, "use polling, instead of inotify")
	flag.Parse()

	if config.ReOpen {
		config.Follow = true
	}
	config.MaxLineSize = maxlinesize
	return config, n
}

func main() {
	config, n := args2config()
	if n != 0 {
		config.Location = &tail.SeekInfo{-n, os.SEEK_END}
	}

	done := make(chan bool)
	for _, filename := range flag.Args() {
		go tailFile(filename, config, done)
	}

	for _, _ = range flag.Args() {
		<-done
	}
}

func tailFile(filename string, config tail.Config, done chan bool) {
	defer func() { done <- true }()
	t, err := tail.TailFile(filename, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	for line := range t.Lines {
		fmt.Println(line.Text) // 这里是真正的输出数据
	}
	fmt.Println(len(t.Lines))
	err = t.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
