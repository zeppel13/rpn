# rpn

RPN is a small reverse polish (programmable) calculator. It has 4 Stack levels and was written by myself in 2015. 
I use it every day to crunch some numbers on the commandline. You can invoke RPN from the terminal and let it parse its 
arguments or start the program and use it interactively.

This cute little piece of written go code is free software by any means. Do what you what you want with. Enjoy it. Change it and distribute it 
as long as you mention my name somewhere. Feel free to contact me. 

https://sebastiankind.de/post/terminalrechner

and 

https://github.com/zeppel13/rpn/blob/master/howto.org

The Link to my blog offers a German description.

- Install Go any version should do fine
- compile rpn with

```
cd ./rpn
go build 
```
- test it 

```
rpn 3 3 a 5 d
```

should give you 1.20

If you wish to compile RPN for older computers e.g. with a Pentium II or anything else which has 387 compatible 
floatingpoint processor, you need to set this in order to compile and run it correctly. Tested for a Pentium II on Windows XP and 
Debian Linux. Somewhere on my old computers HDD should exist a small C Port for a 286 Personal Computer. But it shouldn't be too
difficult to reprogram this software in any language by yourself.

```
GO386=387 go build
```
