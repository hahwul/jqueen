package printing

import (
	"fmt"
	"os"
)

func banner(){
	fmt.Fprintln(os.Stderr, "   __    ____                                *")
	fmt.Fprintln(os.Stderr, "   \\ \\  /___ \\_   _  ___  ___ _ __         * | *")
	fmt.Fprintln(os.Stderr, "    \\ \\//  / / | | |/ _ \\/ _ \\ '_ \\       * \\|/ *")
	fmt.Fprintln(os.Stderr, " /\\_/ / \\_/ /| |_| |  __/  __/ | | | * * * \\|O|/ * * *")
	fmt.Fprintln(os.Stderr, " \\___/\\___,_\\ \\__,_|\\___|\\___|_| |_|  \\o\\o\\o|O|o/o/o/")
	fmt.Fprintln(os.Stderr, "  "+VERSION+"  by @hahwul                  (<><><>O<><><>) ")
	fmt.Fprintln(os.Stderr, "")
}
