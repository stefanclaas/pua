# pua
Encoder/Decoder using a .ttf font for PUA (Private Use Area) Unicode characters.

Useful if you want to tranfer base64 encoded messages, so that mail filters
have probably a hard time to filter messages.

Usage: $ base64 -w 20 < message.txt | pua -f pua.ttf > out.txt  
Decode: $ pua -d -f pua.ttf < out.txt | base64 -d > message.txt
