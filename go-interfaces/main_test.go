package main

import "testing"

func TestGreeting(t *testing.T) {
	// arrange
	englishBot := englishBot{}
	spanishBot := spanishBot{}
	// act
	actualEnglishGreet := englishBot.getGreeting()
	actualSpanishGreet := spanishBot.getGreeting()

	expectedEnglishGreeting := "Hi There!"
	expectedSpanishGreeting := "Ola!"
	// assert
	if expectedEnglishGreeting != actualEnglishGreet {
		t.Errorf("Expected to be: '%s', but got '%s' instead.", expectedEnglishGreeting, actualEnglishGreet)
	}

	if expectedSpanishGreeting != actualSpanishGreet {
		t.Errorf("Expected to be '%s', but got '%s' instead.", expectedSpanishGreeting, actualSpanishGreet)
	}

}
