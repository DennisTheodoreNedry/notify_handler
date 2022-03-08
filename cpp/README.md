# Error Handler
A simple error handler written for c++

There are two functions which are shown in the example below

Example
```
errorHandler _eH;

_eH.error("This is an error message", "functionName");
_eH.error("This is an error message");
```

Result
```
######## Error ########
* Msg: This is an error message
* Func: functionName

######## Error ########
* Msg: This is an error message
* Func: Not provided
```
The function name is supposed to be used when a user reports a bug and/or an error to the developer
