package main

import notify "github.com/DennisTheodoreNedry/notify_handler"

func main() {
	handler := notify.Constructor("1")
	handler.SetLvl("3")
	handler.Log("You can't see this", 4)
	handler.Log("You can see this", 2)

	notify.Inform("Information information!!")
	notify.Warning("Warning warning!!")
	notify.Error("Abort mission!!", "main.main()", 4)
}
