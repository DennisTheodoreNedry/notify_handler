# Error Handler
A simple error handler written for c++

Written for my own projects, but felt that it might have a home here.

There are two functions which are shown in the example below

Example
```
errorHandler _eH;

_eH("This is an error message", "functionName");
_eH("This is an error message");
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
