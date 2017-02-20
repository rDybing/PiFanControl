# Pi Fan Control

CC-BY Roy Dybing, February 2017

A small app to control a Raspberry Pi - running Raspbian Pixel - CPU fan.

As we all know, those little CPU fans that come with the cheap rPi enclosures make quite the racket. At least mine do. The thought of having that thing whirring along all the time is not particularly pleasant. Hence this app that'll only turn it on above a given threshold.

This is a simple on/off control. Chiefly due to me not having a digital potentiometer in stock at the moment. So instead, the CPU fan is hooked up to a simple NPN transistor that will turn it on or off depending on temperature reading from the SoC sensor.

To prevent the CPU fan from going at full tilt, I've added a little resistor to the mix.

Included in the repo is a rPi3B+ native binary. Since it controls hardware, it must be run as sudo or else it'll do absolutely nothing. That only took me an hour of headscratching to figure out :P

## Pin assignment:

- Power, +5V
- Ground, any GND
- Control, GPIO-18

## Temperature Limits:

- Turns on at > 65'C
- Turns off at < 63'C

## More Information:

See https://dybings.blogspot.no for wiring and code walkthrough.