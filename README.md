# Bessel Tone Generator

Bessel-tone-generator or `btg` is a utility for using the Bessel Null technique to set the FM deviation for digital modes such as AFSK for packet.

This program takes as an input the desired deviation and calculates the frequency of an audio tone to emit based on an MI of 2.40466.  It can also optionally toggle PTT using either RTS and/or DTR.  

`btg` is written in the Go language so it's cross platform.

## To use `btg`

1. Run the program
1. Select the audio device depending on the operating system.  Unfortunately, the audio library used for `btg` does not have a provision to select the audio device to open, so this has to be done manually.
1. Set the desired deviation.
1. Set the optional PTT type and the serial port if PTT is something other than NONE.
1. Toggle the PTT either using RTS/DTR or manually on the transmitter.
1. Set the audio drive level as appropriate for your setup.
1. Toggle the PTT off when done.
