package main

import "testing"

var allBowls = make(map[int] (score[2] int))
var scorecard = new(Scorecard)


//Scorecard initializes with score set to 0
func TestInitialize(t *testing.T){
  expected := allBowls
  actual := scorecard.allBowls()
  if actual != expected {
    t.Errorf("Test failed, expected: '%s, got: '%s'", expected, actual)
  }
}
