# Markivian LM in go
This repository contains the function (markov) to create a text using a rudimental LM model.
The program takes in a txt file. It asks for the k, the lenght of characters previous to the current one to take as reference, and the desired lenght of the outputted text.

# Using the function in a program
The function can be called in go just like any other function.
An example of program execution would look like this:
'./markov-lm 13 100 <file.txt' 
Where 13 is k and 100 is the lenght. 