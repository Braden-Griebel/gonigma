# Welcome to Gonigma!

This is an implementation of a basic version of the Enigma machine in go, 
an electromechanical cryptographic device used by the Germans during WWII. 
It was ultimately cracked by the British at Bletchley park, whose code 
breaking efforts significantly shortened the war and saved many lives. 
See https://en.wikipedia.org/wiki/Enigma_machine, 
https://www.iwm.org.uk/history/how-alan-turing-cracked-the-enigma-code, and 
https://bletchleypark.org.uk/our-story/enigma/ for more information.

# Background

The Enigma machine used a plug board, rotors, and a reflecting wheel to encode a 
string of characters. In order to send encoded messages, the settings of the 
machine would need to be agreed upon by both the sender and the receiver in advance. 
Normally this would happen with shared code books which would describe the settings 
for a particular date. These settings included which of the 8 rotors to use, which
reflector to use, the plug board setting, and the initial rotation of each of the
rotors, and reflector.

## Plug board
The plug board used wires to connect two characters together, which would 
then be swapped during the encoding. So if a wire connected a and e, then each 
time 'a' was encoded it would be first swapped to an 'e', and 'e' would similarly
be swapped to an 'a'. Note that due to the reflector, this swapping would occur 
twice during each character encoding.

## Rotors
Each rotor translated an incoming character into another, and could also "step", 
changing which positions would be translated. The rotors also has a setting, 
which describes their starting position. The early Enigma machine used 3 rotors
(which is represented here, later naval versions used 4 rotors and a thin reflector),
arranged in sequence left to right. Each time a key was pressed current would flow 
through each of the rotors, then the reflector wheel, and back through the rotors
in the other direction. Following the key press the rightmost rotor would step, 
then if it stepped past a notch, its left neighbor would similarly step (then if 
that neighbor also stepped past a notch, the next rotor would step as well). This 
turned the simple substitution cipher of each rotor into a cipher that changed 
after each keypress.

## Reflector
In order to easily encode and decode messages, it is useful if you can simply 
feed encoded text back through a machine (with the same initial settings) and
retrieve the decoded text. To allow for this, the Enigma machine used a reflector 
which passed the current back through the rotors and plug board before it encoded 
a letter. This looping back through the machine meant that (as long as both machines
were set up in the same way) encoded text could simply be fed back into the machine
to decode it.

# Usage
This program is designed as a CLI, which will take in settings as flags and then translate either
a given file, or a string read from STDIN. For a more interactive interface, see 
[enigmars](https://github.com/Braden-Griebel/enigmars).

## Installation
Currently, the only installation method is building from source (no binary releases are currently available).
To do this, you must first install the go toolchain (see official instructions [here](https://go.dev/doc/install), 
or use your preferred package manager).   

  
Then clone this repository, and build the binary:
```shell
git clone https://github.com/Braden-Griebel/gonigma.git
cd gonigma
go build
```

This will produce an executable file in the directory. You can then run it using 
```shell
./gonigma [<options>] <filename>
```
or add it to your PATH environment variable and run it with 
```shell
gonigma [<options>] <filename>
```

## Running Gonigma
When running the CLI, you will (optionally) specify the settings for the Enigma machine, as well as 
a file to translate. If no filename is provided, then the STDIN will be read and translated. 
  
The settings for the enigma machine are:
* --rotors, -r: Choose the rotors to use, should be 3 comma seperated roman numerals (I-VII)
* --rotor-settings, -s: Choose the rotor setting, should be 3 comma seperated characters
* --reflector, -f: Choose the reflector, can be A, B or C (single character)
* --reflector-settings, -fs: Choose the reflector setting, should be a single character
* --plugboard-setting, -p: Choose the plugboard settings, should be a comma seperated list of 'wires', 
where each wire is two characters seperated by a '-'

### Example Usage:
To translate a file:
```shell
./gonigma -r I,VII,VII -s r,d,w -f C -fs i -p r-g,j-k,l-w,x-c myfile.txt > translatedfile.txt
```
To translate STDIN:
```shell
echo "Hello Gonigma!" | ./gonigma -r I,VII,VII -s r,d,w -f C -fs i -p r-g,j-k,l-w,x-c
```


