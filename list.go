package main

import (
	"fmt"
	"log"
	"bytes"
	"strings"

	"github.com/jroimartin/gocui"
)


func init() {
}

var trigger = false


func renderInterface() {
	g := gocui.NewGui()
	g.Mouse = true
	g.Cursor = true
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetLayout(layout)

	g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		//remember about revers list index
		if benchObject.listPosition == 1 && !trigger {
			return nil
		}
		_, y := vGl.Cursor()
		//if target wasn't chosen. you can't choose "current" item
		if y == 1 && !trigger {
			return nil
		}
		if benchObject.listPosition - 1 >= 0 {
			benchObject.listPosition--
		} else {
			return nil
		}
		//get offset before moving of index
		nextLine := benchObject.lineOffset("previous")
		vGl.MoveCursor(0, -nextLine, false)
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())
		return nil
	})


	g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		//remember about revers list index
		if benchObject.listPosition == len(benchObject.history) - 1 {
			return nil
		}
		//get offset before moving of index
		nextLine := benchObject.lineOffset("next")
		// benchObject.lastBenchmarkPosition was init like -1
		if benchObject.listPosition + 1 < benchObject.lastBenchmarkPosition {
			benchObject.listPosition++
		} else {
			return nil
		}
		vGl.MoveCursor(0, nextLine, false)
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())

		return nil
	})

	g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if trigger {
			return quit(g, v)
		} else {
			benchObject.lastBenchmark = bytes.NewBufferString(benchObject.history[benchObject.listPosition].result)
			// first argument is "previous current". Second doesn't have sense
			if benchObject.listPosition == 1 {
				//trigger to get current result
				benchObject.listPosition = 0
				return quit(g, v)
			}
			benchObject.lastBenchmarkPosition = benchObject.listPosition
			benchObject.listPosition--
			//get offset before moving of index
			nextLine := benchObject.lineOffset("previous")
			vGl.MoveCursor(0, -nextLine, false)
			trigger = true
			vGl.Clear()
			fmt.Fprintln(vGl, benchObject.choose())
		}
		return nil
	})

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}


}


func (b *benchmarkObject) choose() (str string) {
	str = ""
	for i := 0; i < len(b.history); i++ {
		a := b.history[i]
		stdString := a.StandardStr
		if a.hash == "current" {
			stdString = "current\n"
		}
		if a.hash == "previous current" {
			stdString = "previous current\n"
		}
		if i == b.listPosition {
			str += "[*]" + stdString
		} else if i == b.lastBenchmarkPosition && trigger {
			str += "[#]" + stdString
			break
		} else {
			str += "[]" + stdString
		}
	}
	return
}
func (b *benchmarkObject) lineOffset(course string) int {
	var str string
	for i := 0; i < len(b.history); i++ {
		if i == b.listPosition {
			if course == "next" {
				if b.history[i].hash == "current" || b.history[i].hash == "previous current" {
					return 1
				}
				str = b.history[i].StandardStr
			} else if course == "previous" {
				if b.history[i].hash == "current" || b.history[i].hash == "previous current" {
					return 1
				}
				str = b.history[i].StandardStr
			}
			return strings.Count(str, "\n")
		}
	}
	return 0
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("list", 0, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		vGl = v

		vGl.SelBgColor = gocui.ColorBlue
		vGl.SelFgColor = gocui.ColorCyan
		vGl.Highlight = true
		// set on second item; first item always "current"
		vGl.SetCursor(0,1)
		fmt.Fprintln(vGl, benchObject.choose())
	}
	if v, err := g.SetView("helpLayout", maxX / 2, 0, maxX - maxX/4, maxY/4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = true
		fmt.Fprintln(v, "lol")
	}
	return nil
}

var vGl *gocui.View

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
