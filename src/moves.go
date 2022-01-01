/* Copyright (C) 2021  Evuraan, <evuraan@gmail.com> */

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func moveIt(relOrAbs string) (ok bool) {
	inputMoves := absMove
	if relOrAbs == "rel" {
		inputMoves = relMove
	}
	incoming := strings.TrimSpace(inputMoves)
	splat := strings.Split(incoming, "x")
	success := false
	var x, y int
	var err error
	if len(splat) != 2 {
		goto badAbs
	}

	x, err = strconv.Atoi(splat[0])
	if err != nil {
		goto badAbs
	}

	y, err = strconv.Atoi(splat[1])
	if err != nil {
		goto badAbs
	}
	success = true // Good to go..
	defer destroy()

	if relOrAbs == "abs" {
		return relaySendAbs(x, y)
	} else {
		return relaySendRel(x, y)
	}

badAbs:
	if !success {
		fmt.Fprintf(os.Stderr, "Invalid move coordinates %s\n", absMove)
		os.Exit(1)
	}
	return true
}
