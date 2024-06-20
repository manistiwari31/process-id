package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
)

func Pids() ([]int, error) {
	f, err := os.Open(`/proc`)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	pids := make([]int, 0, len(names))
	for _, name := range names {
		if pid, err := strconv.ParseInt(name, 10, 0); err == nil {
			pids = append(pids, int(pid))
		}
	}
	return pids, nil
}

func ProcPidStat(pid int) ([]byte, error) {
	filename := `/proc/` + strconv.FormatInt(int64(pid), 10) + `/stat`
	return ioutil.ReadFile(filename)
}

func pick(stat []byte) (string, string) {
	count := 0
	var copystr = ""
	var a string
	var b string
	for _, str := range string(stat) {
		if str == ' ' {
			count++
			if count == 1 {
				a = copystr
				copystr = ""
			}
			if count == 2 {
				b = copystr
				return a, b
			}

		} else if str == '(' || str == ')' {

		} else {
			copystr += string(str)
		}
	}
	return "", ""
}

func main() {
	var proccess string
	var killcheck string
	var allowquery bool
	var kill bool
	var still bool
	if len(os.Args) >= 2 {
		proccess = os.Args[1]
		allowquery = true
		if proccess == "show" || proccess == "Show" {
			still = true
			allowquery = false
		}
	}
	if len(os.Args) >= 3 {
		killcheck = os.Args[2]
		if killcheck == "kill" || killcheck == "Kill" {
			kill = true
		}
	}

	pidCollect := []string{}
	pids, err := Pids()
	if err != nil {
		fmt.Println("pids:-", err)
		return
	}
	mapProcess := make(map[string]int)
	if !allowquery {
		for i, pid := range pids {
			stat, err := ProcPidStat(pid)
			if err != nil {
				fmt.Println("pid:-", pid, err)
				return
			}
			a, b := pick(stat)
			mapProcess[b]++
			if !still {
				fmt.Println(i, a, b)
			}

		}
		if still {

			keys := make([]string, 0, len(mapProcess))
			for k := range mapProcess {
				keys = append(keys, k)
			}

			sort.Slice(keys, func(i, j int) bool {
				return mapProcess[keys[i]] > mapProcess[keys[j]]
			})

			for k, v := range keys {
				if k <= 8 {
					fmt.Println(v, mapProcess[v])
				}
			}

		}
		return
	}
	for _, pid := range pids {
		stat, err := ProcPidStat(pid)
		if err != nil {
			fmt.Println("pid:-", pid, err)
			return
		}
		a, b := pick(stat)

		if strings.Contains(b, proccess) {
			pidCollect = append(pidCollect, a)
		}

	}
	fmt.Println("looking for proccess", proccess)
	if len(pidCollect) > 0 {
		fmt.Println(pidCollect)
		if kill {
			status := false
			for _, pid := range pidCollect {
				status = killProcess(pid)
			}
			if status {
				fmt.Println("Process Killed")
			} else {
				fmt.Println("Cannot kill")
			}
		}
	} else {
		fmt.Println("Found nothing")
	}

}

func killProcess(pid string) bool {

	fmt.Println("Killing", pid)

	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Printf("Invalid PID: %s\n", pid)
		return false
	}

	// Find the process by its PID
	process, err := os.FindProcess(pidInt)
	if err != nil {
		fmt.Printf("Process not found: %s\n", pid)
		return false
	}

	// Attempt to kill the process using SIGKILL
	err = process.Signal(syscall.SIGKILL)
	if err != nil {
		fmt.Printf("Failed to kill process %s: %v\n", pid, err)
		return false
	}

	return true
}
