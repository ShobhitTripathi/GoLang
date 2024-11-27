/*Just for fun, we will create some basic ping pong game where two users are represented by two goroutines.
We will explore how concurrency works in Go, understand the communication between two routines, and learn about the usage of context.

Implement a ping pong game using channels and goroutines.
Two players will participate in the game, each represented by a goroutine.
The game will continue for a duration of 2 minutes.
Players will take turns holding the ball, waiting for a random time (0–10 seconds), and then serving it back.
Each player serves at random intervals, adding an element of unpredictability to the game.
At the end of the 2-minute period, the player holding the ball is declared the winner.
Context and goroutines will be utilized to manage the game’s duration and execution.*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Main function
func main() {
	// Initialize a buffered channel to simulate the ping pong ball
	ball := make(chan int, 1)
	// Create a context with a 20-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Start the game by sending the ball to player 1
	ball <- 0

	// Player 1 routine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // Exit if context is canceled
			case <-ball:
				fmt.Println("Ball is in player 1's court, sending...")
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // Random delay before serving
				ball <- 1 // Serve the ball to player 2
			}
		}
	}()

	// Player 2 routine
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // Exit if context is canceled
			case <-ball:
				fmt.Println("Ball is in player 2's court, sending...")
				time.Sleep(time.Duration(rand.Intn(10)) * time.Second) // Random delay before serving
				ball <- 2 // Serve the ball back to player 1
			}
		}
	}()

	// Wait for the game to end or timeout
	select {
	case <-ctx.Done():
		fmt.Println("Game Over! Winner Is Player", <-ball)
	}
}


/*
Explanation:
We create a buffered channel ball to simulate the movement of the ball between players.
Using context.WithTimeout, we define a 120-second(2 min) timeout for the game duration.
Two goroutines represent each player in the game. They continuously listen for the ball on the ball channel.
Each player waits for a random time (0–10 seconds) before serving the ball back to the other player.
The game ends after 2 minutes or when the context is canceled due to timeout.
The winner is determined by whoever has the ball at the end of the game.
*/
