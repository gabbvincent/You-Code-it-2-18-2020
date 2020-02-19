// Programmer name: Vincent Gabb
// Date completed:  2/18/2020
// Description: SEction 1 - You Code It

package main

import (
    "fmt"
    "math/rand"
    
) //adding the ability to do random numbers

func main() {
  
//create a variable for count
  
  var count int 
  
  count = 1
  
//ask the user to enter a max range for the guessing game and store that value in variable max.
  
  fmt.Println("Please enter any number.")
    
  var max int
    
  fmt.Scanln(&max)
   
//this next line creates a random number from 1 to that guess for the computer to know.  You can test this by printing out the variable computerGuess
  
  var computerGuess = rand.Intn(max)

//ask the user to enter a guess for the computer number

  fmt.Println("Now, The Computer is thinking of a number. Try and guess it!")
    
  var userGuess int
    
  fmt.Scanln(&userGuess)

//create a loop that compares the computerGuess to the userGuess while they are NOT equal go into the loop

  for computerGuess != userGuess{
    
//increase the count by 1
  
  count++
  
//tell the user that the guessed incorrect
//ask the user to enter a new guess for the computer number
      
  fmt.Println("Incorrect!" , "Try again!")
      
  fmt.Scanln(&userGuess)
      
  }
    
//print out that the user got the answer correctly and how many guesses it took (the count)
      
  fmt.Println("Congrats! You got it! It only took you", count, "tries.")
}