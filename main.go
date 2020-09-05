package lg

import (
	"fmt"
	"os"
	"os/exec"
	"time"

)

func main() {
	board := NewLiveGameBoard(22, 150,15)
	for {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}

		board.Show()
		board.NextState()
		time.Sleep(100 * time.Millisecond)
	}
}
