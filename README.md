# edotool

- Simulate keystrokes. 
- Like `xdotool` (well, kind of), with support for both `X11` and `Wayland`. 

<img src="https://user-images.githubusercontent.com/39205936/130298707-8c845a13-4438-4bdc-a815-89451f549db6.png" width="50%">

[edotool](https://evuraan.info/evuraan/stuff/edotool.png)

## Screengrab
Using `edotool` keystrokes can be sent remotely to the active window. 
```bash
$ sleep 10; for i in `seq 1 5`; do  ./edotool -i /dev/input/event9 -e "$i: She sells seasells, at the seashore. ${RANDOM}";   ./edotool -i /dev/input/event9 -e "KEY_ENTER"; echo "Sent $i"; done
```

https://user-images.githubusercontent.com/39205936/130187383-776d230f-2dbd-480b-9446-feb444831ed9.mp4

[mp4](https://evuraan.info/evuraan/stuff/edotool.mp4) | [gif](./edotool.gif)

## Usage
```bash
$ ./edotool -h
Usage: ./edotool
  -h  --help             print this usage and exit
  -v  --version          print version information and exit
  -d  --debug            show verbose output
  -k  --keys             show available keys
  -i  /dev/input/event1  kbd device to use
  -e  --events           events to relay
```
## Setup
Either download the [release](https://github.com/evuraan/edotool/releases), or clone this repo and [build](./README.md#optional-build) it yourself. 

### Using /dev/input/eventXX 
If you are  a member of `input` group, `edotool` can be run without root permissions. 
```bash  
sudo gpasswd -a $USER input
newgrp input
```
In this mode, `edotool` will try to locate an appropriate `input` device it can use. 
* If an appropriate kbd device cannot be found, `edotool` will ask you to specify a suitable device using the `-i` option.
```bash
$ ./edotool -i /dev/input/event9 -e "She sells seasells, at the seashore."
She sells seasells, at the seashore.
```
### Using /dev/uinput 
`/dev/uinput` requires root permission:
```bash
$ ./edotool -i /dev/uinput -e "She sells seasells, at the seashore."
Error opening /dev/uinput: Permission denied
```
Retrying, with `sudo` access:
```bash
$ sudo ./edotool -i /dev/uinput -e "She sells seasells, at the seashore." 
She sells seasells, at the seashore.
```
### Sending key codes
Specific keystrokes can be sent, either chained together with a `+` or individually:
```bash
$ ./edotool -i /dev/input/event9 -e "KEY_SPACE + KEY_S + KEY_A + KEY_D + KEY_SPACE"
 sad
```
#### keys supported:
Use `-k` option to get a list of supported keys:
```bash
$ ./edotool -k | head -10
Available keys:
key -->  KEY_LEFT_UP
key -->  KEY_CHANNELUP
key -->  KEY_FN_1
key -->  KEY_BATTERY
key -->  KEY_UWB
key -->  KEY_WWAN
key -->  KEY_BOOKMARKS
key -->  KEY_F24
./snip/.
```
### Debug mode
Send `-d` to turn on debugging. This would produce a <b>lot</b> of debug output:
```bash
$ ./edotool -d -e " Fun. "
Sat Aug 21 12:20:25 2021 edotool/1.03d Copyright © 2021 Evuraan <evuraan@gmail.com>. All rights reserved.
Sat Aug 21 12:20:25 2021 edotool/1.03d This program comes with ABSOLUTELY NO WARRANTY.
Sat Aug 21 12:20:25 2021 edotool/1.03d Howdy!
Sat Aug 21 12:20:25 2021 edotool/1.03d Incoming: | Fun. |
Sat Aug 21 12:20:25 2021 edotool/1.03d keyboard device: /dev/input/event3
Sat Aug 21 12:20:25 2021 [C] [getFd] fd opened: 3
Sat Aug 21 12:20:25 2021 edotool/1.03d translated: KEY_SPACE+KEY_CAPSLOCK + KEY_F + KEY_CAPSLOCK+KEY_U+KEY_N+KEY_DOT+KEY_SPACE
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_SPACE val: 57
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_CAPSLOCK val: 58
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_F val: 33
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_CAPSLOCK val: 58
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_U val: 22
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_N val: 49
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_DOT val: 52
Sat Aug 21 12:20:25 2021 edotool/1.03d key: KEY_SPACE val: 57
Sat Aug 21 12:20:25 2021 edotool/1.03d combo: true
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 57
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 57
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 58
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 58
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 33
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 33
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 58
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 58
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 22
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 22
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 49
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 49
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 52
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 52
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 57
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 1 code 57
Sat Aug 21 12:20:25 2021 [C] [emit] emitted 24 bytes type 0 code 0
Sat Aug 21 12:20:25 2021 [C] [handleComboEvents] Handled 8 events
Sat Aug 21 12:20:25 2021 edotool/1.03d Bye bye!
 Fun.
```
## Optional: Build 
If you prefer to build yourself, you will need the [Go Programming Language](https://golang.org/dl/) installed on your `Linux` System. 

Go into the `src` folder and build as: 
``` 
go build
```
## Related
- [Swipe Gestures on Linux](https://evuraan.info/Swipe/?ref=odotool)

