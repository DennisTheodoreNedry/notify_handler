package notify

import (
	"fmt"
	"os"
)

var Verbose_lvl = "0"    // Default level, this will log nothing to the user
var Exit_on_error = true // Set it to false if you don't want the program to exit when 'Error' is called

func Error(msg string, where string) { // Print msg and exit
	fmt.Println("#### Error ####")
	fmt.Println("msg: " + msg)
	fmt.Println("where: " + where)
	if Exit_on_error {
		os.Exit(1)
	}
}

func Inform(msg string) {
	fmt.Println("[*] Inform: " + msg)
}

func Warning(msg string) {
	fmt.Println("[!] Warning: " + msg)
}

func Log(msg string, verbose_lvl string, suggested_verbose_lvl string) {
	if verbose_lvl != "0" && verbose_lvl >= suggested_verbose_lvl {
		fmt.Println("[%] " + msg)
	}
}
