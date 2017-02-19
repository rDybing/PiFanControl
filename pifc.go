/*****************************
 * piFanControl
 * CCBY Roy Dybing, Feb. 2017
 * github.com/rDybing
 *****************************/
package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type timer_t struct {
	oldTime    int64
	newTime    int64
	intervalMS int64
}

type state_t struct {
	tempC    int
	fanOn    bool
	limitOn  int
	limitOff int
}

func main() {
	var timer timer_t
	var state state_t
	state.limitOn = 65
	state.limitOff = 63
	initTimer(&timer)

	for {
		time.Sleep(time.Millisecond * 100)
		if checkTimer(&timer) {
			state.tempC = getTemp()
			setFan(&state)
			fmt.Printf("temp: %d'C\n", state.tempC)
		}
	}
}

func setFan(s *state_t) {
	if s.fanOn {
		if s.tempC < s.limitOff {
			// turn off fan
			s.fanOn = false
			fmt.Println("fan off")
		}
	} else {
		if s.tempC > s.limitOn {
			// turn on fan
			s.fanOn = true
			fmt.Println("fan on")
		}
	}
}

func getTemp() int {
	cmd := exec.Command("vcgencmd", "measure_temp")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		printError(err)
	}
	tStr := string(cmdOutput.Bytes())
	tStr = cleanString(tStr)
	temp, err := strconv.ParseFloat(tStr, 64)
	if err != nil {
		printError(err)
	}
	return round(temp)
}

func checkTimer(t *timer_t) bool {
	t.newTime = millis()
	if t.newTime > t.oldTime+t.intervalMS {
		t.oldTime = t.newTime
		return true
	}
	return false
}

func initTimer(t *timer_t) {
	t.intervalMS = 2000
	t.oldTime = millis()
	t.newTime = t.oldTime
}

func cleanString(s string) string {
	s = stripString(s, "temp=")
	s = stripString(s, "'C")
	s = stripString(s, "\n")
	return s
}

func stripString(s string, r string) string {
	s = strings.Replace(s, r, "", -1)
	return s
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func millis() int64 {
	return time.Now().Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
