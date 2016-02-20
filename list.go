package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"time"
)


func init() {
}

func renderInterface(b *benchmarkObject) {
	conductor.position = 0
	conductor.arr = b.history
	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	println("lol")
	g.SetLayout(layout)

	fmt.Printf("%v\n", vGl)
	fmt.Fprintln(vGl, b.choose())
	g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if conductor.position > 0 {
			conductor.position--
		}
		vGl.Clear()
		fmt.Fprintln(vGl, b.choose())
		return nil
	})
	g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if conductor.position < len(conductor.arr) - 1 {
			conductor.position++
		}
		vGl.Clear()
		fmt.Fprintln(vGl, b.choose())

		return nil
	})

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}

	for ;; {
		println("хуй")
		time.Sleep(time.Second)
	}


}

var conductor struct {
	position int
	arr      []bench
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
	}
	return nil
}

var vGl *gocui.View

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
