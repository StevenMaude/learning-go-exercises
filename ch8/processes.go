package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func getProcessScanner() (*bufio.Scanner, error) {
	cmd := exec.Command("ps", "-e", "-opid,ppid,comm")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(&out), err
}

func collectPIDs(scanner *bufio.Scanner) (map[int][]int, error) {
	pids := make(map[int][]int)

	// Skip header line.
	scanner.Scan()
	for scanner.Scan() {
		l := scanner.Text()
		fields := strings.Fields(l)
		pid, err := strconv.Atoi(fields[0])
		if err != nil {
			return nil, err
		}
		ppid, err := strconv.Atoi(fields[1])
		if err != nil {
			return nil, err
		}
		pids[ppid] = append(pids[ppid], pid)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return pids, nil
}

func printPIDs(pids map[int][]int) {
	ppids := make([]int, len(pids))
	i := 0
	for k := range pids {
		ppids[i] = k
		i++
	}

	sort.Ints(ppids)

	for _, ppid := range ppids {
		children := pids[ppid]
		var childrenEnd string
		if len(children) > 1 {
			childrenEnd = "ren"
		}
		fmt.Printf("Pid %v has %v child%v: %v\n",
			ppid, len(children), childrenEnd, children)
	}
}

func main() {
	scanner, err := getProcessScanner()
	if err != nil {
		log.Fatal(err)
	}

	pids, err := collectPIDs(scanner)
	if err != nil {
		log.Fatal(err)
	}

	printPIDs(pids)
}
