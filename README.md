# pua
Encoder/Decoder using a .ttf font for PUA (Private Use Area) Unicode characters.

Useful if you want to tranfer base64 encoded messages, so that third parties,
which do not have the font, cannot read the messages.

Usage: $ base64 -w 20 < message.txt | pua -f pua.ttf > out.txt  
Decode: $ pua -d -f pua.ttf < out.txt | base64 -d > message.txt
