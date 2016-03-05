package main

import (
	"fmt"
	"log"
	"bytes"

	"github.com/jroimartin/gocui"
)


func init() {
}

var trigger = false


func renderInterface() {
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	g.SetLayout(layout)

	g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if benchObject.listPosition < len(benchObject.history) - 1 {
			benchObject.listPosition++
		}
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())
		return nil
	})
	g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		// benchObject.lastBenchmarkPosition was init like -1
		if benchObject.listPosition - 1 > benchObject.lastBenchmarkPosition {
			benchObject.listPosition--
		}
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())

		return nil
	})

	g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if trigger {
			return quit(g, v)
		} else {
			// first argument don't be the current result
			if benchObject.listPosition != len(benchObject.history) - 1 {
				benchObject.lastBenchmark = bytes.NewBufferString(benchObject.history[benchObject.listPosition].result)
				// first argument is "previous current". Second doesn't have sense
				if benchObject.listPosition == len(benchObject.history) - 2 {
					//trigger to get current result
					benchObject.listPosition = len(benchObject.history) - 1
					return quit(g, v)
				}
				benchObject.lastBenchmarkPosition = benchObject.listPosition
				benchObject.listPosition++
				trigger = true
				vGl.Clear()
				fmt.Fprintln(vGl, benchObject.choose())
			}
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
	for i := len(b.history) - 1; i >= 0; i-- {
		a := b.history[i]
		stdString := fmt.Sprint(a.Date) + "\n" + a.Message + "\n"
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
		} else {
			str += "[]" + stdString
		}
	}
	/*
	for i, a := range b.history {
		if i == b.listPosition {
			str += "[*]" + a.hash + "\n"
		} else if i == b.lastBenchmarkPosition {
			str += "[#]" + a.hash + "\n"
		} else {
			str += "[]" + a.hash + "\n"
		}
	}
	*/
	return
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("hello", 0, 0, maxX, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		vGl = v
		fmt.Fprintln(vGl, benchObject.choose())
	}
	return nil
}

var vGl *gocui.View

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
