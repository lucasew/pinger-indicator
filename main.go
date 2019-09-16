package main

import (
	"bytes"
	"github.com/getlantern/systray"
	_golog "github.com/lucasew/golog"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"
)

const checkInterval = time.Second

var log = _golog.Default.NewLogger("main")

var pingaddr = "google.com"

func main() {
	var err error
	matchRegexp, err = regexp.Compile("time=\\d+")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		pingaddr = os.Args[1]
	}
	log.Info("Inicializando icone de tray")
	systray.Run(onStart, nil)
}

var matchRegexp *regexp.Regexp

func parseLatency(output string) (ms int) {
	i, err := strconv.Atoi(matchRegexp.FindAllString(output, -1)[0][5:])
	if err != nil {
		return 9999
	}
	return i
}

var latency = 9999

func onStart() {
	systray.SetIcon(icoNoNet)
	for {
		time.Sleep(checkInterval)
		cmd := exec.Command("ping", "-c 1", pingaddr)
		var output bytes.Buffer
		cmd.Stdout = &output
		err := cmd.Start()
		if err != nil {
			log.Error(err.Error())
			continue
		}
		err = cmd.Wait()
		if err != nil {
			log.Error(err.Error())
		}
		log.Verbose(0, "%s: ecode: %d latencia: %d\n", cmd.Args, cmd.ProcessState.ExitCode(), latency)
		if cmd.ProcessState.ExitCode() == 0 {
			latency = parseLatency(output.String())
			if latency > 1000 {
				systray.SetIcon(icoBadNet)
				continue
			}
			systray.SetIcon(icoGoodNet)
			continue
        } else {
            systray.SetIcon(icoNoNet)
        }
	}
}
