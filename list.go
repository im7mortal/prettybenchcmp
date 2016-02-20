package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"bytes"
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
		if benchObject.listPosition > 0 {
			benchObject.listPosition--
		}
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())
		return nil
	})
	g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if benchObject.listPosition < len(benchObject.history) - 1 {
			benchObject.listPosition++
		}
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())

		return nil
	})

	g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if trigger {
			benchObject.currentBenchmark = bytes.NewBufferString(benchObject.history[benchObject.listPosition].result)
			return quit(g, v)
		} else {
			benchObject.lastBenchmark = bytes.NewBufferString(benchObject.history[benchObject.listPosition].result)
			trigger = true
		}
		vGl.Clear()
		fmt.Fprintln(vGl, benchObject.choose())
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
	for i, a := range b.history {
		if i == b.listPosition {
			str += "[*]" + a.hash + "\n"
		} else {
			str += "[]" + a.hash + "\n"
		}
	}
	return
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("hello", 0, 0, maxX / 2 + 7, maxY / 2 + 2); err != nil {
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
