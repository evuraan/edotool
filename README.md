# edotool

- Simulate keystrokes. 
- Like `xdotool`, with support for both `X11` and `Wayland`. 

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
$ ./edotool -i /dev/input/event9 -e "KEY_SPACE + KEY_M + KEY_A + KEY_D + KEY_SPACE" -d
Thu Aug 19 22:41:09 2021 edotool/1.03a Incoming events: KEY_SPACE + KEY_M + KEY_A + KEY_D + KEY_SPACE
Thu Aug 19 22:41:09 2021edotool/1.03aCopyright © 2021 Evuraan <evuraan@gmail.com>. All rights reserved.
This program comes with ABSOLUTELY NO WARRANTY.
Thu Aug 19 22:41:09 2021 edotool/1.03a Howdy!
Thu Aug 19 22:41:09 2021 edotool/1.03a keyboard device: /dev/input/event9
Thu Aug 19 22:41:09 2021 [C] [getFd] fd opened: 3
Thu Aug 19 22:41:09 2021 edotool/1.03a key: KEY_SPACE val: 57
Thu Aug 19 22:41:09 2021 edotool/1.03a key: KEY_M val: 50
Thu Aug 19 22:41:09 2021 edotool/1.03a key: KEY_A val: 30
Thu Aug 19 22:41:09 2021 edotool/1.03a key: KEY_D val: 32
Thu Aug 19 22:41:09 2021 edotool/1.03a key: KEY_SPACE val: 57
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 57
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 50
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 30
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 32
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 57
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 0 code 0
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 57
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 50
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 30
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 32
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 1 code 57
Thu Aug 19 22:41:09 2021 [C] [emit] emitted 24 bytes type 0 code 0
Thu Aug 19 22:41:09 2021 [C] [handleEvents] Handled 5 events
Thu Aug 19 22:41:09 2021 edotool/1.03a Bye bye!
 mad
```
## Optional: Build 
If you prefer to build yourself, you will need the [Go Programming Language](https://golang.org/dl/) installed on your `Linux` System. 

Go into the `src` folder and build as: 
``` 
go build
```
## Related
- [Swipe Gestures on Linux](https://evuraan.info/Swipe/?ref=odotool)

