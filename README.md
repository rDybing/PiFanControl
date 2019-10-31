# pifc.go

A small app to control a 2 wire CPU fan on a Raspberry Pi running Raspbian Pixel.

As we all know, those little CPU fans that come with the cheap rPi enclosures make quite the racket. At least mine do. The thought of having that thing whirring along all the time is not particularly pleasant. Hence this app that'll only turn it on above a given threshold.

This is a simple on/off control. Chiefly due to me not having a digital potentiometer in stock at the moment. So instead, the CPU fan is hooked up to a simple n-chan mosfet that will turn it on or off depending on temperature reading from the SoC sensor.

To prevent the CPU fan from going at full tilt, I've added a little resistor to the mix.

To prevent the mosfet burning out providing constant current to the fan - when on the mosfet will essentially be switching at a given frequency set in the interval constant.

**Pin assignment:**

- Power, +5V
- Ground, any GND
- Control, GPIO-18

**Temperature Limits:**

- Turns on at > 65'C
- Turns off at < 63'C

## Build

**3rd party libraries:**
- github.com/kidoman/embd
- github.com/kidoman/embd/host/rpi

First make sure you have Go installed and configured correctly, then get 3rd party libraries installed by means of the `go get` command. When in loacal directory of this repo, enter `go build pifc.go`. Finally run using `sudo ./pifc`.

This app should in principle work on any Raspberry Pi, but is only tested on a 3B.

**Contact:**

location   | name/handle |
-----------|-------------|
github:    | rDybing     |
twitter:   | @DybingRoy  |
Linked In: | Roy Dybing  |
MeWe:      | Roy Dybing  |

---

## Releases

- Version format: [major release].[new feature(s)].[bugfix patch-version]
- Date format: yyyy-mm-dd

#### v.1.2.0: 2019-07-06
- Cleaned up from old version released way back in 2017.
- Added switching when fan is on rather than it being constant on.

---

## License: MIT

**Copyright © 2017-2019 Roy Dybing** 

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions: The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

---

ʕ◔ϖ◔ʔ


## More Information:

See https://dybings.blogspot.no for wiring and code walkthrough.