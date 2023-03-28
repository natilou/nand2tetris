// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

(INIT)
@8192     // 8k memory words of 16bits --> 8 x 1024 = 8192
D=A 
@n        // index start at 8192
M=D

(LOOP)
@n
M=M-1
D=M
@INIT
D;JLT

@KBD     // get keyboard (selected data memory registrer)
D=M

@WHITE   // if RAM[KBD] == 0 goto white
D;JEQ

@BLACK   // else goto BLACK
D;JMP

// WHITE SCREEN
(WHITE)
@SCREEN   // get screen address register
D=A
@n       // word address = sum screen address + index 
A=D+M
M=0      // word calue = 0
@LOOP
0;JMP

// BLACK SCREEN
(BLACK)
@SCREEN  // get screen address register
D=A
@n       // word address = sum screen address + index
A=D+M
M=-1     // word value = -1
@LOOP
0;JMP
