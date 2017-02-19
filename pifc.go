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
)

func main() {
	var tStr string
	var tInt int

	cmd := exec.Command("vcgencmd", "measure_temp")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	if err := cmd.Run(); err != nil {
		printError(err)
	}
	tStr = string(cmdOutput.Bytes())
	tStr = cleanString(tStr)
	temp, err := strconv.ParseFloat(tStr, 64)
	if err != nil {
		printError(err)
	}
	tInt = round(temp)
	fmt.Printf("string : %s\n", tStr)
	fmt.Printf("length : %d\n", len(tStr))
	fmt.Printf("float  : %f\n", temp)
	fmt.Printf("integer: %d\n", tInt)
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

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
