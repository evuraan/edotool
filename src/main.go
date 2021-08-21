/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

/*
#cgo CFLAGS: -Wall -Werror
#include <stdio.h>
#include <sys/types.h>		// open
#include <sys/stat.h>		// open
#include <fcntl.h>		// open
#include <unistd.h>		// read
#include <sys/epoll.h>		// epoll
#include <stdlib.h>		// EXIT_FAILURE
#include <time.h> 		// time stuff
#include <stdint.h> // need int32
#include <stdarg.h> // vaprintf
#include <string.h> // strncpy

#include <linux/input.h> // /dev/input stuff
#include <linux/uinput.h> // for uinput fu

#define ARRAY_MAX 8192 // arrayLen in go
#define BUF 8192
#define END 65535

int fd = 0;
unsigned debug = 0;

int supportedKeys[] = {KEY_RESERVED, KEY_ESC, KEY_1, KEY_2, KEY_3, KEY_4, KEY_5, KEY_6, KEY_7, KEY_8, KEY_9, KEY_0, KEY_MINUS, KEY_EQUAL, KEY_BACKSPACE, KEY_TAB, KEY_Q, KEY_W, KEY_E, KEY_R, KEY_T, KEY_Y, KEY_U, KEY_I, KEY_O, KEY_P, KEY_LEFTBRACE, KEY_RIGHTBRACE, KEY_ENTER, KEY_LEFTCTRL, KEY_A, KEY_S, KEY_D, KEY_F, KEY_G, KEY_H, KEY_J, KEY_K, KEY_L, KEY_SEMICOLON, KEY_APOSTROPHE, KEY_GRAVE, KEY_LEFTSHIFT, KEY_BACKSLASH, KEY_Z, KEY_X, KEY_C, KEY_V, KEY_B, KEY_N, KEY_M, KEY_COMMA, KEY_DOT, KEY_SLASH, KEY_RIGHTSHIFT, KEY_KPASTERISK, KEY_LEFTALT, KEY_SPACE, KEY_CAPSLOCK, KEY_F1, KEY_F2, KEY_F3, KEY_F4, KEY_F5, KEY_F6, KEY_F7, KEY_F8, KEY_F9, KEY_F10, KEY_NUMLOCK, KEY_SCROLLLOCK, KEY_KP7, KEY_KP8, KEY_KP9, KEY_KPMINUS, KEY_KP4, KEY_KP5, KEY_KP6, KEY_KPPLUS, KEY_KP1, KEY_KP2, KEY_KP3, KEY_KP0, KEY_KPDOT, KEY_ZENKAKUHANKAKU, KEY_102ND, KEY_F11, KEY_F12, KEY_RO, KEY_KATAKANA, KEY_HIRAGANA, KEY_HENKAN, KEY_KATAKANAHIRAGANA, KEY_MUHENKAN, KEY_KPJPCOMMA, KEY_KPENTER, KEY_RIGHTCTRL, KEY_KPSLASH, KEY_SYSRQ, KEY_RIGHTALT, KEY_LINEFEED, KEY_HOME, KEY_UP, KEY_PAGEUP, KEY_LEFT, KEY_RIGHT, KEY_END, KEY_DOWN, KEY_PAGEDOWN, KEY_INSERT, KEY_DELETE,  KEY_MUTE, KEY_VOLUMEDOWN, KEY_VOLUMEUP, KEY_POWER, KEY_KPEQUAL, KEY_KPPLUSMINUS, KEY_PAUSE, KEY_SCALE, KEY_KPCOMMA, KEY_HANGEUL, KEY_HANJA, KEY_YEN, KEY_LEFTMETA, KEY_RIGHTMETA, KEY_COMPOSE, KEY_STOP, KEY_AGAIN, KEY_PROPS, KEY_UNDO, KEY_FRONT, KEY_COPY, KEY_OPEN, KEY_PASTE, KEY_FIND, KEY_CUT, KEY_HELP, KEY_MENU, KEY_CALC, KEY_SETUP, KEY_SLEEP, KEY_WAKEUP, KEY_FILE, KEY_SENDFILE, KEY_DELETEFILE, KEY_XFER, KEY_PROG1, KEY_PROG2, KEY_WWW, KEY_MSDOS, KEY_COFFEE, KEY_ROTATE_DISPLAY, KEY_CYCLEWINDOWS, KEY_MAIL, KEY_BOOKMARKS, KEY_COMPUTER, KEY_BACK, KEY_FORWARD, KEY_CLOSECD, KEY_EJECTCD, KEY_EJECTCLOSECD, KEY_NEXTSONG, KEY_PLAYPAUSE, KEY_PREVIOUSSONG, KEY_STOPCD, KEY_RECORD, KEY_REWIND, KEY_PHONE, KEY_ISO, KEY_CONFIG, KEY_HOMEPAGE, KEY_REFRESH, KEY_EXIT, KEY_MOVE, KEY_EDIT, KEY_SCROLLUP, KEY_SCROLLDOWN, KEY_KPLEFTPAREN, KEY_KPRIGHTPAREN, KEY_NEW, KEY_REDO, KEY_F13, KEY_F14, KEY_F15, KEY_F16, KEY_F17, KEY_F18, KEY_F19, KEY_F20, KEY_F21, KEY_F22, KEY_F23, KEY_F24, KEY_PLAYCD, KEY_PAUSECD, KEY_PROG3, KEY_PROG4, KEY_DASHBOARD, KEY_SUSPEND, KEY_CLOSE, KEY_PLAY, KEY_FASTFORWARD, KEY_BASSBOOST, KEY_PRINT, KEY_HP, KEY_CAMERA, KEY_SOUND, KEY_QUESTION, KEY_EMAIL, KEY_CHAT, KEY_SEARCH, KEY_CONNECT, KEY_FINANCE, KEY_SPORT, KEY_SHOP, KEY_ALTERASE, KEY_CANCEL, KEY_BRIGHTNESSDOWN, KEY_BRIGHTNESSUP, KEY_MEDIA, KEY_SWITCHVIDEOMODE, KEY_KBDILLUMTOGGLE, KEY_KBDILLUMDOWN, KEY_KBDILLUMUP, KEY_SEND, KEY_REPLY, KEY_FORWARDMAIL, KEY_SAVE, KEY_DOCUMENTS, KEY_BATTERY, KEY_BLUETOOTH, KEY_WLAN, KEY_UWB, KEY_UNKNOWN, KEY_VIDEO_NEXT, KEY_VIDEO_PREV, KEY_BRIGHTNESS_CYCLE, KEY_BRIGHTNESS_AUTO, KEY_DISPLAY_OFF, KEY_WWAN, KEY_RFKILL, KEY_MICMUTE, KEY_OK, KEY_SELECT, KEY_GOTO, KEY_CLEAR, KEY_POWER2, KEY_OPTION, KEY_INFO, KEY_TIME, KEY_VENDOR, KEY_ARCHIVE, KEY_PROGRAM, KEY_CHANNEL, KEY_FAVORITES, KEY_EPG, KEY_PVR, KEY_MHP, KEY_LANGUAGE, KEY_TITLE, KEY_SUBTITLE, KEY_ANGLE, KEY_FULL_SCREEN, KEY_MODE, KEY_KEYBOARD, KEY_ASPECT_RATIO, KEY_PC, KEY_TV, KEY_TV2, KEY_VCR, KEY_VCR2, KEY_SAT, KEY_SAT2, KEY_CD, KEY_TAPE, KEY_RADIO, KEY_TUNER, KEY_PLAYER, KEY_TEXT, KEY_DVD, KEY_AUX, KEY_MP3, KEY_AUDIO, KEY_VIDEO, KEY_DIRECTORY, KEY_LIST, KEY_MEMO, KEY_CALENDAR, KEY_RED, KEY_GREEN, KEY_YELLOW, KEY_BLUE, KEY_CHANNELUP, KEY_CHANNELDOWN, KEY_FIRST, KEY_LAST, KEY_AB, KEY_NEXT, KEY_RESTART, KEY_SLOW, KEY_SHUFFLE, KEY_BREAK, KEY_PREVIOUS, KEY_DIGITS, KEY_TEEN, KEY_TWEN, KEY_VIDEOPHONE, KEY_GAMES, KEY_ZOOMIN, KEY_ZOOMOUT, KEY_ZOOMRESET, KEY_WORDPROCESSOR, KEY_EDITOR, KEY_SPREADSHEET, KEY_GRAPHICSEDITOR, KEY_PRESENTATION, KEY_DATABASE, KEY_NEWS, KEY_VOICEMAIL, KEY_ADDRESSBOOK, KEY_MESSENGER, KEY_DISPLAYTOGGLE, KEY_SPELLCHECK, KEY_LOGOFF, KEY_DOLLAR, KEY_EURO, KEY_FRAMEBACK, KEY_FRAMEFORWARD, KEY_CONTEXT_MENU, KEY_MEDIA_REPEAT, KEY_10CHANNELSUP, KEY_10CHANNELSDOWN, KEY_IMAGES, KEY_DEL_EOL, KEY_DEL_EOS, KEY_INS_LINE, KEY_DEL_LINE, KEY_FN, KEY_FN_ESC, KEY_FN_F1, KEY_FN_F2, KEY_FN_F3, KEY_FN_F4, KEY_FN_F5, KEY_FN_F6, KEY_FN_F7, KEY_FN_F8, KEY_FN_F9, KEY_FN_F10, KEY_FN_F11, KEY_FN_F12, KEY_FN_1, KEY_FN_2, KEY_FN_D, KEY_FN_E, KEY_FN_F, KEY_FN_S, KEY_FN_B,  KEY_BRL_DOT1, KEY_BRL_DOT2, KEY_BRL_DOT3, KEY_BRL_DOT4, KEY_BRL_DOT5, KEY_BRL_DOT6, KEY_BRL_DOT7, KEY_BRL_DOT8, KEY_BRL_DOT9, KEY_BRL_DOT10, KEY_NUMERIC_0, KEY_NUMERIC_1, KEY_NUMERIC_2, KEY_NUMERIC_3, KEY_NUMERIC_4, KEY_NUMERIC_5, KEY_NUMERIC_6, KEY_NUMERIC_7, KEY_NUMERIC_8, KEY_NUMERIC_9, KEY_NUMERIC_STAR, KEY_NUMERIC_POUND, KEY_NUMERIC_A, KEY_NUMERIC_B, KEY_NUMERIC_C, KEY_NUMERIC_D, KEY_CAMERA_FOCUS, KEY_WPS_BUTTON, KEY_TOUCHPAD_TOGGLE, KEY_TOUCHPAD_ON, KEY_TOUCHPAD_OFF, KEY_CAMERA_ZOOMIN, KEY_CAMERA_ZOOMOUT, KEY_CAMERA_UP, KEY_CAMERA_DOWN, KEY_CAMERA_LEFT, KEY_CAMERA_RIGHT, KEY_ATTENDANT_ON, KEY_ATTENDANT_OFF, KEY_ATTENDANT_TOGGLE, KEY_LIGHTS_TOGGLE, KEY_ALS_TOGGLE, KEY_ROTATE_LOCK_TOGGLE, KEY_BUTTONCONFIG, KEY_TASKMANAGER, KEY_JOURNAL, KEY_CONTROLPANEL, KEY_APPSELECT, KEY_SCREENSAVER, KEY_VOICECOMMAND, KEY_ASSISTANT, KEY_KBD_LAYOUT_NEXT, KEY_BRIGHTNESS_MIN, KEY_BRIGHTNESS_MAX, KEY_KBDINPUTASSIST_PREV, KEY_KBDINPUTASSIST_NEXT, KEY_KBDINPUTASSIST_PREVGROUP, KEY_KBDINPUTASSIST_NEXTGROUP, KEY_KBDINPUTASSIST_ACCEPT, KEY_KBDINPUTASSIST_CANCEL, KEY_RIGHT_UP, KEY_RIGHT_DOWN, KEY_LEFT_UP, KEY_LEFT_DOWN, KEY_ROOT_MENU, KEY_MEDIA_TOP_MENU, KEY_NUMERIC_11, KEY_NUMERIC_12, KEY_AUDIO_DESC, KEY_3D_MODE, KEY_NEXT_FAVORITE, KEY_STOP_RECORD, KEY_PAUSE_RECORD, KEY_VOD, KEY_UNMUTE, KEY_FASTREVERSE, KEY_SLOWREVERSE, KEY_DATA, KEY_ONSCREEN_KEYBOARD, END};

unsigned enableDebug(){
	debug = 1;
	return debug;
}


char *get_currentTime() {
	static char currentTime[2048] = {0};
	time_t t;
	time(&t);
	strncpy(currentTime, ctime(&t), 1000);
	currentTime[strcspn(currentTime, "\r\n")] = 0;	// works for LF, CR, CRLF, LFCR, ...
	return currentTime;
}

void print(char *format, ...) {
	if (!debug){
		return;
	}
	va_list arguments;
	va_start(arguments, format);
	char temp[BUF] = {0};
	snprintf(temp, BUF, "%s [C] %s", get_currentTime(), format);
	vfprintf(stdout, temp, arguments);
	va_end(arguments);
}

unsigned emit(int fd, int type, int code, int val) {
	struct input_event ie = { 0 };
	ie.type = type;
	ie.code = code;
	ie.value = val;
	size_t wrote = write(fd, &ie, sizeof(ie));
	if (wrote > 0) {
		print("[%s] emitted %ld bytes type %d code %d\n", __func__, wrote, type , code);
		return wrote;
	} else {
		return 0;
	}
}

int getFd(char *device) {
	if (fd) {
		return fd;
	}
	int fda = open(device, O_WRONLY | O_NONBLOCK);
	if (fda < 1) {
		fprintf(stderr, "Error opening %s: ", device);
		perror("");
		exit(1);
		return -1;
	}
	fd = fda;
	print("[%s] fd opened: %d\n", __func__, fd);
	return fd;
}

int getUinputFd(){
	if (fd){
		return fd;
	}

	char *device = "/dev/uinput";
	int uinFd = open("/dev/uinput", O_WRONLY | O_NONBLOCK);
	if (uinFd < 1){
		fprintf(stderr, "Error opening %s: ", device);
		perror("");
		exit(1);
		return -1;
	}
	fd = uinFd;

	if (ioctl(uinFd, UI_SET_EVBIT, EV_KEY) < 0){
		perror("UI_SET_EVBIT: ");
		exit(1);
	}

	int i = 0;
	while(1){
		if (supportedKeys[i] == END){
			break;
		}
    		if (ioctl(uinFd, UI_SET_KEYBIT, supportedKeys[i]) < 0){
			fprintf(stderr, "UI_SET_KEYBIT err on %d\n", supportedKeys[i] );
			perror("UI_SET_KEYBIT: ");
		}
		i++;
	}


	struct uinput_setup usetup = {0};
	usetup.id.bustype = BUS_USB;
	usetup.id.vendor = 0x1234;
	usetup.id.product = 0x5678;
	snprintf(usetup.name, UINPUT_MAX_NAME_SIZE, "Evuraan's virtual Keyboard");
	if (ioctl(fd, UI_DEV_SETUP, &usetup) <0){
		perror("UI_DEV_SETUP: ");
		exit(1);
		return -1;
	}

	if (ioctl(uinFd, UI_DEV_CREATE) < 0){
		perror("UI_DEV_CREATE: ");
		exit(1);
		return -1;
	}
	sleep(1);
	return 0;
}


int closeFd() {
	if (!fd) {
		close(fd);
		fd = -1;
	}
	return fd;
}

void printInt(int someInt){
	print("[%s] int: %d\n", __func__, someInt);
}

void handleEvents(int32_t *events){
	unsigned x = 0;
	// Agenda: Key press, report the event, send key release, and report again

	// do key press
	for (int i = 0; i < ARRAY_MAX; i++){
		if (events[i] == END){
			break;
		}
		x++;
		emit(fd, EV_KEY, events[i], 1);
	}

	if (!x){
		// we got nothing.
		return;
	}
	// report the keys
	emit(fd, EV_SYN, SYN_REPORT, 0);
	// release keys
	for (int i = 0; i < ARRAY_MAX; i++){
		if (events[i] == END){
			break;
		}
		emit(fd, EV_KEY, events[i], 0);
	}
	// report the release
	emit(fd, EV_SYN, SYN_REPORT, 0);

	print("[%s] Handled %d events\n", __func__, x);
}


void handleComboEvents(int32_t *events){
	unsigned x = 0;
	// Agenda: Key press, report the event, send key release, and report again

	// do key press
	for (int i = 0; i < ARRAY_MAX; i++){
		if (events[i] == END){
			break;
		}
		x++;
		emit(fd, EV_KEY, events[i], 1);
		emit(fd, EV_SYN, SYN_REPORT, 0);
		emit(fd, EV_KEY, events[i], 0);
		emit(fd, EV_SYN, SYN_REPORT, 0);
	}

	if (!x){
		// we got nothing.
		return;
	}

	print("[%s] Handled %d events\n", __func__, x);
}


*/
import "C"
import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
	"unsafe"
)

var (
	deBug         = false
	notifyBool    = false
	workChan      chan string
	kbd           = ""
	inputEvents   = ""
	eventLibStuff = &eventLib{eventCodes: map[string]int{"KEY_RESERVED": 0, "KEY_ESC": 1, "KEY_1": 2, "KEY_2": 3, "KEY_3": 4, "KEY_4": 5, "KEY_5": 6, "KEY_6": 7, "KEY_7": 8, "KEY_8": 9, "KEY_9": 10, "KEY_0": 11, "KEY_MINUS": 12, "KEY_EQUAL": 13, "KEY_BACKSPACE": 14, "KEY_TAB": 15, "KEY_Q": 16, "KEY_W": 17, "KEY_E": 18, "KEY_R": 19, "KEY_T": 20, "KEY_Y": 21, "KEY_U": 22, "KEY_I": 23, "KEY_O": 24, "KEY_P": 25, "KEY_LEFTBRACE": 26, "KEY_RIGHTBRACE": 27, "KEY_ENTER": 28, "KEY_LEFTCTRL": 29, "KEY_A": 30, "KEY_S": 31, "KEY_D": 32, "KEY_F": 33, "KEY_G": 34, "KEY_H": 35, "KEY_J": 36, "KEY_K": 37, "KEY_L": 38, "KEY_SEMICOLON": 39, "KEY_APOSTROPHE": 40, "KEY_GRAVE": 41, "KEY_LEFTSHIFT": 42, "KEY_BACKSLASH": 43, "KEY_Z": 44, "KEY_X": 45, "KEY_C": 46, "KEY_V": 47, "KEY_B": 48, "KEY_N": 49, "KEY_M": 50, "KEY_COMMA": 51, "KEY_DOT": 52, "KEY_SLASH": 53, "KEY_RIGHTSHIFT": 54, "KEY_KPASTERISK": 55, "KEY_LEFTALT": 56, "KEY_SPACE": 57, "KEY_CAPSLOCK": 58, "KEY_F1": 59, "KEY_F2": 60, "KEY_F3": 61, "KEY_F4": 62, "KEY_F5": 63, "KEY_F6": 64, "KEY_F7": 65, "KEY_F8": 66, "KEY_F9": 67, "KEY_F10": 68, "KEY_NUMLOCK": 69, "KEY_SCROLLLOCK": 70, "KEY_KP7": 71, "KEY_KP8": 72, "KEY_KP9": 73, "KEY_KPMINUS": 74, "KEY_KP4": 75, "KEY_KP5": 76, "KEY_KP6": 77, "KEY_KPPLUS": 78, "KEY_KP1": 79, "KEY_KP2": 80, "KEY_KP3": 81, "KEY_KP0": 82, "KEY_KPDOT": 83, "KEY_ZENKAKUHANKAKU": 85, "KEY_102ND": 86, "KEY_F11": 87, "KEY_F12": 88, "KEY_RO": 89, "KEY_KATAKANA": 90, "KEY_HIRAGANA": 91, "KEY_HENKAN": 92, "KEY_KATAKANAHIRAGANA": 93, "KEY_MUHENKAN": 94, "KEY_KPJPCOMMA": 95, "KEY_KPENTER": 96, "KEY_RIGHTCTRL": 97, "KEY_KPSLASH": 98, "KEY_SYSRQ": 99, "KEY_RIGHTALT": 100, "KEY_LINEFEED": 101, "KEY_HOME": 102, "KEY_UP": 103, "KEY_PAGEUP": 104, "KEY_LEFT": 105, "KEY_RIGHT": 106, "KEY_END": 107, "KEY_DOWN": 108, "KEY_PAGEDOWN": 109, "KEY_INSERT": 110, "KEY_DELETE": 111, "KEY_MACRO": 112, "KEY_MUTE": 113, "KEY_VOLUMEDOWN": 114, "KEY_VOLUMEUP": 115, "KEY_POWER": 116, "KEY_KPEQUAL": 117, "KEY_KPPLUSMINUS": 118, "KEY_PAUSE": 119, "KEY_SCALE": 120, "KEY_KPCOMMA": 121, "KEY_HANGEUL": 122, "KEY_HANJA": 123, "KEY_YEN": 124, "KEY_LEFTMETA": 125, "KEY_RIGHTMETA": 126, "KEY_COMPOSE": 127, "KEY_STOP": 128, "KEY_AGAIN": 129, "KEY_PROPS": 130, "KEY_UNDO": 131, "KEY_FRONT": 132, "KEY_COPY": 133, "KEY_OPEN": 134, "KEY_PASTE": 135, "KEY_FIND": 136, "KEY_CUT": 137, "KEY_HELP": 138, "KEY_MENU": 139, "KEY_CALC": 140, "KEY_SETUP": 141, "KEY_SLEEP": 142, "KEY_WAKEUP": 143, "KEY_FILE": 144, "KEY_SENDFILE": 145, "KEY_DELETEFILE": 146, "KEY_XFER": 147, "KEY_PROG1": 148, "KEY_PROG2": 149, "KEY_WWW": 150, "KEY_MSDOS": 151, "KEY_COFFEE": 152, "KEY_ROTATE_DISPLAY": 153, "KEY_CYCLEWINDOWS": 154, "KEY_MAIL": 155, "KEY_BOOKMARKS": 156, "KEY_COMPUTER": 157, "KEY_BACK": 158, "KEY_FORWARD": 159, "KEY_CLOSECD": 160, "KEY_EJECTCD": 161, "KEY_EJECTCLOSECD": 162, "KEY_NEXTSONG": 163, "KEY_PLAYPAUSE": 164, "KEY_PREVIOUSSONG": 165, "KEY_STOPCD": 166, "KEY_RECORD": 167, "KEY_REWIND": 168, "KEY_PHONE": 169, "KEY_ISO": 170, "KEY_CONFIG": 171, "KEY_HOMEPAGE": 172, "KEY_REFRESH": 173, "KEY_EXIT": 174, "KEY_MOVE": 175, "KEY_EDIT": 176, "KEY_SCROLLUP": 177, "KEY_SCROLLDOWN": 178, "KEY_KPLEFTPAREN": 179, "KEY_KPRIGHTPAREN": 180, "KEY_NEW": 181, "KEY_REDO": 182, "KEY_F13": 183, "KEY_F14": 184, "KEY_F15": 185, "KEY_F16": 186, "KEY_F17": 187, "KEY_F18": 188, "KEY_F19": 189, "KEY_F20": 190, "KEY_F21": 191, "KEY_F22": 192, "KEY_F23": 193, "KEY_F24": 194, "KEY_PLAYCD": 200, "KEY_PAUSECD": 201, "KEY_PROG3": 202, "KEY_PROG4": 203, "KEY_DASHBOARD": 204, "KEY_SUSPEND": 205, "KEY_CLOSE": 206, "KEY_PLAY": 207, "KEY_FASTFORWARD": 208, "KEY_BASSBOOST": 209, "KEY_PRINT": 210, "KEY_HP": 211, "KEY_CAMERA": 212, "KEY_SOUND": 213, "KEY_QUESTION": 214, "KEY_EMAIL": 215, "KEY_CHAT": 216, "KEY_SEARCH": 217, "KEY_CONNECT": 218, "KEY_FINANCE": 219, "KEY_SPORT": 220, "KEY_SHOP": 221, "KEY_ALTERASE": 222, "KEY_CANCEL": 223, "KEY_BRIGHTNESSDOWN": 224, "KEY_BRIGHTNESSUP": 225, "KEY_MEDIA": 226, "KEY_SWITCHVIDEOMODE": 227, "KEY_KBDILLUMTOGGLE": 228, "KEY_KBDILLUMDOWN": 229, "KEY_KBDILLUMUP": 230, "KEY_SEND": 231, "KEY_REPLY": 232, "KEY_FORWARDMAIL": 233, "KEY_SAVE": 234, "KEY_DOCUMENTS": 235, "KEY_BATTERY": 236, "KEY_BLUETOOTH": 237, "KEY_WLAN": 238, "KEY_UWB": 239, "KEY_UNKNOWN": 240, "KEY_VIDEO_NEXT": 241, "KEY_VIDEO_PREV": 242, "KEY_BRIGHTNESS_CYCLE": 243, "KEY_BRIGHTNESS_AUTO": 244, "KEY_DISPLAY_OFF": 245, "KEY_WWAN": 246, "KEY_RFKILL": 247, "KEY_MICMUTE": 248, "KEY_OK": 0x160, "KEY_SELECT": 0x161, "KEY_GOTO": 0x162, "KEY_CLEAR": 0x163, "KEY_POWER2": 0x164, "KEY_OPTION": 0x165, "KEY_INFO": 0x166, "KEY_TIME": 0x167, "KEY_VENDOR": 0x168, "KEY_ARCHIVE": 0x169, "KEY_PROGRAM": 0x16a, "KEY_CHANNEL": 0x16b, "KEY_FAVORITES": 0x16c, "KEY_EPG": 0x16d, "KEY_PVR": 0x16e, "KEY_MHP": 0x16f, "KEY_LANGUAGE": 0x170, "KEY_TITLE": 0x171, "KEY_SUBTITLE": 0x172, "KEY_ANGLE": 0x173, "KEY_FULL_SCREEN": 0x174, "KEY_MODE": 0x175, "KEY_KEYBOARD": 0x176, "KEY_ASPECT_RATIO": 0x177, "KEY_PC": 0x178, "KEY_TV": 0x179, "KEY_TV2": 0x17a, "KEY_VCR": 0x17b, "KEY_VCR2": 0x17c, "KEY_SAT": 0x17d, "KEY_SAT2": 0x17e, "KEY_CD": 0x17f, "KEY_TAPE": 0x180, "KEY_RADIO": 0x181, "KEY_TUNER": 0x182, "KEY_PLAYER": 0x183, "KEY_TEXT": 0x184, "KEY_DVD": 0x185, "KEY_AUX": 0x186, "KEY_MP3": 0x187, "KEY_AUDIO": 0x188, "KEY_VIDEO": 0x189, "KEY_DIRECTORY": 0x18a, "KEY_LIST": 0x18b, "KEY_MEMO": 0x18c, "KEY_CALENDAR": 0x18d, "KEY_RED": 0x18e, "KEY_GREEN": 0x18f, "KEY_YELLOW": 0x190, "KEY_BLUE": 0x191, "KEY_CHANNELUP": 0x192, "KEY_CHANNELDOWN": 0x193, "KEY_FIRST": 0x194, "KEY_LAST": 0x195, "KEY_AB": 0x196, "KEY_NEXT": 0x197, "KEY_RESTART": 0x198, "KEY_SLOW": 0x199, "KEY_SHUFFLE": 0x19a, "KEY_BREAK": 0x19b, "KEY_PREVIOUS": 0x19c, "KEY_DIGITS": 0x19d, "KEY_TEEN": 0x19e, "KEY_TWEN": 0x19f, "KEY_VIDEOPHONE": 0x1a0, "KEY_GAMES": 0x1a1, "KEY_ZOOMIN": 0x1a2, "KEY_ZOOMOUT": 0x1a3, "KEY_ZOOMRESET": 0x1a4, "KEY_WORDPROCESSOR": 0x1a5, "KEY_EDITOR": 0x1a6, "KEY_SPREADSHEET": 0x1a7, "KEY_GRAPHICSEDITOR": 0x1a8, "KEY_PRESENTATION": 0x1a9, "KEY_DATABASE": 0x1aa, "KEY_NEWS": 0x1ab, "KEY_VOICEMAIL": 0x1ac, "KEY_ADDRESSBOOK": 0x1ad, "KEY_MESSENGER": 0x1ae, "KEY_DISPLAYTOGGLE": 0x1af, "KEY_SPELLCHECK": 0x1b0, "KEY_LOGOFF": 0x1b1, "KEY_DOLLAR": 0x1b2, "KEY_EURO": 0x1b3, "KEY_FRAMEBACK": 0x1b4, "KEY_FRAMEFORWARD": 0x1b5, "KEY_CONTEXT_MENU": 0x1b6, "KEY_MEDIA_REPEAT": 0x1b7, "KEY_10CHANNELSUP": 0x1b8, "KEY_10CHANNELSDOWN": 0x1b9, "KEY_IMAGES": 0x1ba, "KEY_DEL_EOL": 0x1c0, "KEY_DEL_EOS": 0x1c1, "KEY_INS_LINE": 0x1c2, "KEY_DEL_LINE": 0x1c3, "KEY_FN": 0x1d0, "KEY_FN_ESC": 0x1d1, "KEY_FN_F1": 0x1d2, "KEY_FN_F2": 0x1d3, "KEY_FN_F3": 0x1d4, "KEY_FN_F4": 0x1d5, "KEY_FN_F5": 0x1d6, "KEY_FN_F6": 0x1d7, "KEY_FN_F7": 0x1d8, "KEY_FN_F8": 0x1d9, "KEY_FN_F9": 0x1da, "KEY_FN_F10": 0x1db, "KEY_FN_F11": 0x1dc, "KEY_FN_F12": 0x1dd, "KEY_FN_1": 0x1de, "KEY_FN_2": 0x1df, "KEY_FN_D": 0x1e0, "KEY_FN_E": 0x1e1, "KEY_FN_F": 0x1e2, "KEY_FN_S": 0x1e3, "KEY_FN_B": 0x1e4, "KEY_BRL_DOT1": 0x1f1, "KEY_BRL_DOT2": 0x1f2, "KEY_BRL_DOT3": 0x1f3, "KEY_BRL_DOT4": 0x1f4, "KEY_BRL_DOT5": 0x1f5, "KEY_BRL_DOT6": 0x1f6, "KEY_BRL_DOT7": 0x1f7, "KEY_BRL_DOT8": 0x1f8, "KEY_BRL_DOT9": 0x1f9, "KEY_BRL_DOT10": 0x1fa, "KEY_NUMERIC_0": 0x200, "KEY_NUMERIC_1": 0x201, "KEY_NUMERIC_2": 0x202, "KEY_NUMERIC_3": 0x203, "KEY_NUMERIC_4": 0x204, "KEY_NUMERIC_5": 0x205, "KEY_NUMERIC_6": 0x206, "KEY_NUMERIC_7": 0x207, "KEY_NUMERIC_8": 0x208, "KEY_NUMERIC_9": 0x209, "KEY_NUMERIC_STAR": 0x20a, "KEY_NUMERIC_POUND": 0x20b, "KEY_NUMERIC_A": 0x20c, "KEY_NUMERIC_B": 0x20d, "KEY_NUMERIC_C": 0x20e, "KEY_NUMERIC_D": 0x20f, "KEY_CAMERA_FOCUS": 0x210, "KEY_WPS_BUTTON": 0x211, "KEY_TOUCHPAD_TOGGLE": 0x212, "KEY_TOUCHPAD_ON": 0x213, "KEY_TOUCHPAD_OFF": 0x214, "KEY_CAMERA_ZOOMIN": 0x215, "KEY_CAMERA_ZOOMOUT": 0x216, "KEY_CAMERA_UP": 0x217, "KEY_CAMERA_DOWN": 0x218, "KEY_CAMERA_LEFT": 0x219, "KEY_CAMERA_RIGHT": 0x21a, "KEY_ATTENDANT_ON": 0x21b, "KEY_ATTENDANT_OFF": 0x21c, "KEY_ATTENDANT_TOGGLE": 0x21d, "KEY_LIGHTS_TOGGLE": 0x21e, "KEY_ALS_TOGGLE": 0x230, "KEY_ROTATE_LOCK_TOGGLE": 0x231, "KEY_BUTTONCONFIG": 0x240, "KEY_TASKMANAGER": 0x241, "KEY_JOURNAL": 0x242, "KEY_CONTROLPANEL": 0x243, "KEY_APPSELECT": 0x244, "KEY_SCREENSAVER": 0x245, "KEY_VOICECOMMAND": 0x246, "KEY_ASSISTANT": 0x247, "KEY_KBD_LAYOUT_NEXT": 0x248, "KEY_BRIGHTNESS_MIN": 0x250, "KEY_BRIGHTNESS_MAX": 0x251, "KEY_KBDINPUTASSIST_PREV": 0x260, "KEY_KBDINPUTASSIST_NEXT": 0x261, "KEY_KBDINPUTASSIST_PREVGROUP": 0x262, "KEY_KBDINPUTASSIST_NEXTGROUP": 0x263, "KEY_KBDINPUTASSIST_ACCEPT": 0x264, "KEY_KBDINPUTASSIST_CANCEL": 0x265, "KEY_RIGHT_UP": 0x266, "KEY_RIGHT_DOWN": 0x267, "KEY_LEFT_UP": 0x268, "KEY_LEFT_DOWN": 0x269, "KEY_ROOT_MENU": 0x26a, "KEY_MEDIA_TOP_MENU": 0x26b, "KEY_NUMERIC_11": 0x26c, "KEY_NUMERIC_12": 0x26d, "KEY_AUDIO_DESC": 0x26e, "KEY_3D_MODE": 0x26f, "KEY_NEXT_FAVORITE": 0x270, "KEY_STOP_RECORD": 0x271, "KEY_PAUSE_RECORD": 0x272, "KEY_VOD": 0x273, "KEY_UNMUTE": 0x274, "KEY_FASTREVERSE": 0x275, "KEY_SLOWREVERSE": 0x276, "KEY_DATA": 0x277, "KEY_ONSCREEN_KEYBOARD": 0x278, "KEY_PRIVACY_SCREEN_TOGGLE": 0x279, "KEY_SELECTIVE_SCREENSHOT": 0x27a, "KEY_MACRO1": 0x290, "KEY_MACRO2": 0x291, "KEY_MACRO3": 0x292, "KEY_MACRO4": 0x293, "KEY_MACRO5": 0x294, "KEY_MACRO6": 0x295, "KEY_MACRO7": 0x296, "KEY_MACRO8": 0x297, "KEY_MACRO9": 0x298, "KEY_MACRO10": 0x299, "KEY_MACRO11": 0x29a, "KEY_MACRO12": 0x29b, "KEY_MACRO13": 0x29c, "KEY_MACRO14": 0x29d, "KEY_MACRO15": 0x29e, "KEY_MACRO16": 0x29f, "KEY_MACRO17": 0x2a0, "KEY_MACRO18": 0x2a1, "KEY_MACRO19": 0x2a2, "KEY_MACRO20": 0x2a3, "KEY_MACRO21": 0x2a4, "KEY_MACRO22": 0x2a5, "KEY_MACRO23": 0x2a6, "KEY_MACRO24": 0x2a7, "KEY_MACRO25": 0x2a8, "KEY_MACRO26": 0x2a9, "KEY_MACRO27": 0x2aa, "KEY_MACRO28": 0x2ab, "KEY_MACRO29": 0x2ac, "KEY_MACRO30": 0x2ad, "KEY_MACRO_RECORD_START": 0x2b0, "KEY_MACRO_RECORD_STOP": 0x2b1, "KEY_MACRO_PRESET_CYCLE": 0x2b2, "KEY_MACRO_PRESET1": 0x2b3, "KEY_MACRO_PRESET2": 0x2b4, "KEY_MACRO_PRESET3": 0x2b5, "KEY_KBD_LCD_MENU1": 0x2b8, "KEY_KBD_LCD_MENU2": 0x2b9, "KEY_KBD_LCD_MENU3": 0x2ba, "KEY_KBD_LCD_MENU4": 0x2bb, "KEY_KBD_LCD_MENU5": 0x2bc, "KEY_MAX": 0x2ff}}
	otherMap      = map[string]string{"=": "KEY_EQUAL", "\n": "KEY_ENTER", ",": "KEY_COMMA", "/": "KEY_SLASH", " ": "KEY_SPACE", "\t": "KEY_TAB", ".": "KEY_DOT", "[": "KEY_LEFTBRACE", "]": "KEY_RIGHTBRACE", "?": "KEY_QUESTION", "\\": "KEY_BACKSLASH", "//": "KEY_SLASH", "+": "KEY_KPPLUS", "*": "KEY_KPASTERISK"}
)

const (
	progName  = "edotool"
	ver       = "1.03d"
	tag       = progName + "/" + ver
	layout    = "Mon Jan 02 15:04:05 2006"
	notifyCmd = "notify-send " + progName
	arrayLen  = 8192 // max key+key+key events: 8
)

type eventLib struct {
	sync.RWMutex
	eventCodes map[string]int
	otherMaps  map[string]string
}

func main() {

	parseArgs()
	if len(inputEvents) < 1 {
		fmt.Fprintf(os.Stderr, "Err 4.1 - inputEvents too short? (Hint: -e option)\n")
		os.Exit(1)
	}

	print("Copyright © 2021 Evuraan <evuraan@gmail.com>. All rights reserved.")
	print("This program comes with ABSOLUTELY NO WARRANTY.")
	print("Howdy!")
	print("Incoming: |%s|", inputEvents)
	eventLibStuff.otherMaps = otherMap
	workChan = make(chan string, 2)
	go func() {
		for cmdString := range workChan {
			if len(cmdString) > 0 {
				go doRun(cmdString)
			}
		}
	}()
	if len(kbd) < 1 {
		kbd = getDeviceForPattern(keyboard)
	}
	if len(kbd) < 1 {
		fmt.Fprintf(os.Stderr, "Error: Could not find %s device\nPlease use the '-i' option to specify a suitable device.\n", keyboard)
		os.Exit(1)
	}
	print("%s device: %s", keyboard, kbd)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if strings.Contains(kbd, "/dev/input/event") {
			cKbd := C.CString(kbd)
			C.getFd(cKbd)
			defer wg.Done()
		} else {
			C.getUinputFd()
			defer wg.Done()
		}
	}()
	wg.Wait()

	eventLibStuff.handleEvent()

	C.closeFd()
	print("Bye bye!")

}

func parseArgs() {
	argc := len(os.Args)
	if argc > 1 {
		for i := range os.Args {
			arg := os.Args[i]
			if strings.Contains(arg, "help") || arg == "h" || arg == "--h" || arg == "-h" || arg == "?" {
				showhelp()
				os.Exit(0)
			}
			if strings.Contains(arg, "keys") || arg == "k" || arg == "--k" || arg == "-k" {
				fmt.Println("Showing keys, arg is", arg)
				showKeys()
				os.Exit(0)
			}

			if strings.Contains(arg, "version") || arg == "v" || arg == "--v" || arg == "-v" {
				fmt.Println("Version:", tag)
				os.Exit(0)
			}
			if strings.Contains(arg, "debug") || arg == "d" || arg == "--d" || arg == "-d" {
				deBug = true
				C.enableDebug()
			}
			if strings.Contains(arg, "events") || arg == "e" || arg == "--e" || arg == "-e" {
				nextArg := i + 1
				if argc > nextArg {
					inputEvents = os.Args[nextArg]
				}
			}

			if arg == "-i" {
				nextArg := i + 1
				if argc > nextArg {
					kbd = os.Args[nextArg]
					if strings.Contains(kbd, "/dev/input/event") || strings.Contains(kbd, "/dev/uinput") {
					} else {
						fmt.Fprintf(os.Stderr, "Invalid input device %s\n", kbd)
						os.Exit(1)
					}
					_, err := os.Stat(kbd)
					if err != nil {
						fmt.Fprintf(os.Stderr, "Could not access %s: %v\n", kbd, err)
						os.Exit(1)
					}
				}
			}
		}
	}

}

func showKeys() {
	fmt.Println("Available keys:")
	eventLibStuff.showKeys()
}

func showhelp() {
	fmt.Printf("Usage: %s\n", os.Args[0])
	fmt.Println("  -h  --help             print this usage and exit")
	fmt.Println("  -v  --version          print version information and exit")
	fmt.Println("  -d  --debug            show verbose output")
	fmt.Println("  -k  --keys             show available keys")
	fmt.Println("  -i  /dev/input/event1  kbd device to use")
	fmt.Println("  -e  --events           events to relay")
}

func (eventLibPtr *eventLib) showKeys() {
	self := eventLibPtr
	self.RLock()
	defer self.RUnlock()
	for i := range self.eventCodes {
		fmt.Println("key --> ", i)
	}
	fmt.Printf("%d keys available\n", len(self.eventCodes))
}

func (eventLibPtr *eventLib) handleEvent() (handleStatBool bool) {
	self := eventLibPtr

	switch {
	case strings.Contains(inputEvents, "KEY_"):
		return self.process(inputEvents, false)
	default:
		longString := ""
		blip := ""
		for i := range inputEvents {
			key := inputEvents[i]
			keyString := string(key)
			upperCase := strings.ToUpper(keyString)
			self.RLock()
			hit, ok := self.otherMaps[keyString]
			self.RUnlock()
			switch {
			case ok:
				blip = hit
			default:
				keyu := strings.ToUpper(keyString)
				if keyString == upperCase {
					blip = fmt.Sprintf("KEY_CAPSLOCK + KEY_%s + KEY_CAPSLOCK", keyu)
				} else {
					blip = fmt.Sprintf("KEY_%s", keyu)
				}
			}
			if len(blip) < 1 {
				print("44.1 Cannot handle: %s", keyString)
				continue
			}
			longString = fmt.Sprintf("%s+%s", longString, blip)
		}

		if strings.HasPrefix(longString, "+") {
			longString = longString[1:]
		}
		print("translated: %s", longString)
		if len(longString) < 1 {
			print("translated string is too short")
			return false
		}
		return self.process(longString, true)
	}
}

func (eventLibPtr *eventLib) process(someString string, combo bool) (x bool) {
	self := eventLibPtr
	splat := strings.Split(someString, "+")
	if len(splat) < 1 {
		return
	}
	k := 0
	var eventArray [arrayLen]int32
	self.RLock()
	defer self.RUnlock()
	for i := range splat {
		keyEvent := splat[i]
		keyEvent = strings.TrimSpace(keyEvent)
		keyIntVal, ok := self.eventCodes[keyEvent]
		if !ok {
			print("Missing keycode for %s", keyEvent)
			continue
		}
		print("key: %s val: %d", keyEvent, keyIntVal)
		eventArray[k] = int32(keyIntVal)
		k++
	}
	if k < 1 {
		return false
	}
	print("combo: %v", combo)
	eventArray[k] = 65535
	if combo {
		C.handleComboEvents((*C.int)(unsafe.Pointer(&eventArray[0])))
	} else {
		C.handleEvents((*C.int)(unsafe.Pointer(&eventArray[0])))
	}

	return true
}

func print(strings string, args ...interface{}) {
	if !deBug {
		return
	}
	a := time.Now()
	msg := fmt.Sprintf(strings, args...)
	if msg[len(msg)-1] == '\n' {
		fmt.Print(a.Format(layout), tag, msg)
	} else {
		fmt.Println(a.Format(layout), tag, msg)
	}
}

func doRun(cmdIn string) error {

	if len(cmdIn) < 1 {
		err := errors.New("cmdIn len 0")
		return err
	}

	safetySplat := strings.Split(cmdIn, " ")
	cmdSplat := []string{}
	x := 0
	for i := range safetySplat {
		block := safetySplat[i]
		if block != "" {
			cmdSplat = append(cmdSplat, block)
			x++
		}
	}
	if x < 1 {
		err := errors.New("splat x 0")
		return err
	}

	cmd := exec.Command(cmdSplat[0], cmdSplat[1:]...)
	err := cmd.Run()
	return err
}
