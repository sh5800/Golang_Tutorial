package main

import "testing"

func Test_UpdateMessage(t *testing.T){
	msg = "Hello, World!"
	wg.Add(2)

	go updateMessage("X")
	go updateMessage("Good Bye ! Cruel World !")

	wg.Wait()

	if msg != "Good Bye ! Cruel World !" {
		t.Errorf("Expected 'Good Bye ! Cruel World !' but got %s", msg)
	}

}