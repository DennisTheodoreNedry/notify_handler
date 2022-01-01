# Notify handler
There are 4 functions in this project.

`notify.Error(msg, function)`<br/>
`notify.Inform(msg)`<br/>
`notify.Warning(msg)`<br/>
`notify.Log(msg, log_level, verbose_level)`<br/>


Example
```
package main

import (
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	notify.Inform("Hi")
	notify.Warning("Hi")
	notify.Log("Hi", "3", "1")           // The user will see this as the verbose level >= log_level
	notify.Log("severe error", "1", "3") // The user will not see this as the verbose level < log_level
	notify.Error("Hi", "main.main()")
}

```

Result
```
[*] Inform: Hi
[!] Warning: Hi
[%] Hi
#### Error ####
msg: Hi
where: main.main()
```
