package notify

import (
	"fmt"
	"os"
	"strconv"

	"github.com/s9rA16Bf4/notify_handler/colors"
)

type verbose_t struct {
	verbose_lvl int
}

// Constructor, creates and returns a verbose object
func Constructor(default_verbose_lvl string) verbose_t {
	var new_verbose verbose_t
	new_verbose.SetLvl(default_verbose_lvl)

	return new_verbose
}

// Sets the internal verbose level for this object
// Will call `object.Error` if the new verbose lvl
// is invalid
func (handler *verbose_t) SetLvl(new_verbose_lvl string) {

	result, err := strconv.Atoi(new_verbose_lvl) // Tries to convert

	if err != nil {
		Error(fmt.Sprintf("Failed to convert %s to an integer", new_verbose_lvl), "notify.SetLvl()", 1)
	}

	handler.verbose_lvl = result
}

// Prints the error message and exits
func Error(msg string, where string, return_code int) {
	fmt.Printf("%s#### Error ####%s\n", colors.Red, colors.Reset)
	fmt.Printf("%smsg: %s %s\n", colors.Red, colors.Reset, msg)
	fmt.Printf("%swhere: %s %s\n", colors.Red, colors.Reset, where)
	os.Exit(return_code)
}

// Prints a message in green
func Inform(msg string) {
	fmt.Printf("%s[*] Inform:%s %s\n", colors.Blue, colors.Reset, msg)
}

// Prints a message in yellow
func Warning(msg string) {
	fmt.Printf("%s[!] Warning:%s %s\n", colors.Cyan, colors.Reset, msg)
}

// Prints a message depending if the configured verbose lvl
// is high enough
func (handler *verbose_t) Log(msg string, suggested_verbose_lvl int) {
	if handler.verbose_lvl >= suggested_verbose_lvl {
		fmt.Println("[%] " + msg)
	}
}
