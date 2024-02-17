# Welcome

This is a working repo in order to build a proper header parser for CS2 Demo
The goal is to make a contribution to github.com/markus-wa/demoinfocs-golang 

To get the demo I'm using in the main.go you can download them at https://drive.google.com/drive/folders/1BdIbBezunpm4Hh9RbzqAFQCC8yNvjotR?usp=drive_link

Feel free to also edit the main.go if  you wanna use your own demo

# Main Problematic

The `common.DemoHeader` struct of demoinfocs-golang [here](https://github.com/markus-wa/demoinfocs-golang/blob/master/pkg/demoinfocs/common/common.go#L28) is filed by parsing the firsts bytes of a Demo for CS:GO demo
but for CS2 the Header is filled after an event and is incomplet.


```go
type DemoHeader struct {
	Filestamp       string        // aka. File-type, must be HL2DEMO
	Protocol        int           // Should be 4
	NetworkProtocol int           // Not sure what this is for
	ServerName      string        // Server's 'hostname' config value
	ClientName      string        // Usually 'GOTV Demo'
	MapName         string        // E.g. de_cache, de_nuke, cs_office, etc.
	GameDirectory   string        // Usually 'csgo'
	PlaybackTime    time.Duration // Demo duration in seconds (= PlaybackTicks / Server's tickrate)
	PlaybackTicks   int           // Game duration in ticks (= PlaybackTime * Server's tickrate)
	PlaybackFrames  int           // Amount of 'frames' aka demo-ticks recorded (= PlaybackTime * Demo's recording rate)
	SignonLength    int           // Length of the Signon package in bytes
}
```

The code provided by this repo is able to fill all the string variables of this struct by parsing the firsts bytes just like CS:GO.
I'm still looking for the "int" fields because I don't know what they look like.
If anybody can help me.

# The program

It actually display the variables founded on three different demo, and shows the missing parts of the header that i wasn't able to figure out.
![alt text](https://github.com/IMhide/cs2-headerparse/blob/master/screenshot.png?raw=true)

