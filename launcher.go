package main

import (
	"bytes"
	"fmt"
	"github.com/andlabs/ui"
	"os"
	"os/exec"
	"time"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: launcher <program>")
		os.Exit(1)
	}

	program := args[1]
	ret := exists(program)
	if !ret {
		fmt.Printf("Error: %s not found\n", program)
		os.Exit(1)
	}

	err := ui.Main(func() {

		message := fmt.Sprintf("Press to cancel %s starting ....", program)
		button := ui.NewButton(message)

		box := ui.NewVerticalBox()
		message = fmt.Sprintf("About to start %s ....", program)
		box.Append(ui.NewLabel(message), false)
		ip := ui.NewProgressBar()
		ip.SetValue(0)
		box.Append(ip,false)
		box.Append(button, true)

		window := ui.NewWindow("Launcher", 600, 300, false)
		window.SetMargined(true)
		window.SetChild(box)

		button.OnClicked(func(*ui.Button) {
			os.Exit(0)
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
		go counter(ip, program)
	})
	if err != nil {
		panic(err)
	}
}

func counter(progress *ui.ProgressBar, prg string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		ui.QueueMain(func() {
			progress.SetValue(i * 10)
		})
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(prg)
	stdout.Reset()
	stderr.Reset()
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		//Ignore errors
	}
	os.Exit(0)
}

// exists reports whether the named file or directory exists.
func exists(name string) bool {
	if _ , err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

