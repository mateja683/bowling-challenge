package main

import ("fmt"
        "math/rand"
        "time"
      )

  var bowl1 int
  var bowl2 int
  var bowl3 int
  var scorecard []int
  var finalFrame []int
  var framesTotals []int
  var bonusTotals []int
  var totalScore []int


func random(min, max int) int {
     rand.Seed(int64(time.Now().Nanosecond()))
     return rand.Intn(max - min) + min
  }

func bowl() {
  for i :=1; i <=9; i++ {
    bowl1 = random(0,11)
  if bowl1 == 10 {
    bowl2 = 0
  } else {
    bowl2 = random(0, 11 - bowl1)
  }
  pushScores(bowl1, bowl2)
  }
}


func bowlFinalFrame() {
  bowl1 = random(0,11)
  if bowl1 == 10 {
    bowl2 = random(0,11)
    if bowl2 == 10 {
      bowl3 = random(0,11)
    } else {
      bowl3 = random(0, 11-bowl2)
    }
  } else {
    bowl2 = random(0, 11-bowl1)
    if (bowl1 + bowl2 == 10) {
      bowl3 = random(0,11)
    } else {
      bowl3 = 0
    }
  pushFinalFrameScore(bowl1, bowl2, bowl3)
  }
}


func pushScores(bowl1, bowl2 int) {
  scorecard = append(scorecard, bowl1, bowl2)
}

func pushFinalFrameScore(bowl1, bowl2, bowl3 int) {
  finalFrame = append(finalFrame, bowl1, bowl2, bowl3)
  scorecard = append(scorecard, bowl1, bowl2, bowl3)
}

func addScores(scorecard []int) int {
  sum := 0
  for _, value := range scorecard {
    sum += value
  }
  return sum
}

func createFrames(scorecard []int) [][]int {
  frames := make([][]int, 9)
  for i, idx := 0, 0; i < 9; i, idx = i + 1, idx + 2 {
    frames[i] = scorecard[idx:idx+2]
  }
  frames = append(frames, finalFrame)
  return frames
}


func calcFramesTotal(frames [][]int) []int {
  for i := 0; i < len(frames); i++ {
    frame := frames[i]
    sum := 0
      for _, value := range frame {
        sum += value
      }
      framesTotals = append(framesTotals, sum)
  }
  return framesTotals
}


// REFACTOR THIS TO USE CASES, NOT NESTED IFS
func calcBonusScore(framesTotals, scorecard []int) []int {
  for i := 0; i < len(framesTotals)-2; i++ {
    var bonus int
    if framesTotals[i] == 10 {
      if scorecard[i*2] == 10 {
        if scorecard[((i*2)+2)] == 10 {
          // case two strikes are bowled consecutively
          bonus = framesTotals[i+1] + scorecard[((i*2)+4)]
        } else {
          // case one strike is bowled
          bonus = framesTotals[i+1]
        }
      } else {
        // case a spare is bowled
        bonus = scorecard[(i*2)+2]
      }
    }
    bonusTotals = append(bonusTotals, bonus)
  }
  return bonusTotals
}

func calcFinalFrameBonusScore(framesTotals, scorecard []int) []int {
  var bonus int
  if framesTotals[8] == 10 {
    if scorecard[16] == 10 {
      bonus = scorecard[18] + scorecard[19]
    } else {
      bonus = scorecard[18]
    }
  }
  bonusTotals = append(bonusTotals, bonus)
  bonusTotals = append(bonusTotals, 0)
  return bonusTotals
}


func calcTotalScore(framesTotals, bonusTotals []int) []int {
  var framePlusBonus int
  for i := 0; i < len(framesTotals); i++ {
    framePlusBonus = framesTotals[i] + bonusTotals[i]
    totalScore = append(totalScore, framePlusBonus)
  }
  return totalScore
}

func addTotalScores(totalScore []int) int {
  sum := 0
  for _, value := range totalScore {
    sum += value
  }
  return (sum / 2) // WHY IS THIS BEING DOUBLED? SOMETHING TO DO WITH CALLING THE FUNCTION, NOT THE ARGUMENT
}


func main () {
  bowl()
  bowlFinalFrame()

  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(0, 0)
  // pushScores(10, 0)
  // pushFinalFrameScore(10, 8, 0)



  fmt.Println("Scorecard:",scorecard)

  fmt.Println("Frames:",createFrames(scorecard))
  fmt.Println("Score without bonus", calcFramesTotal(createFrames(scorecard)))
  calcBonusScore(framesTotals, scorecard)
  calcFinalFrameBonusScore(framesTotals, scorecard)
  fmt.Println("Bonus Scores:", bonusTotals)
  fmt.Println("Frame Total:", calcTotalScore(framesTotals, bonusTotals))
  fmt.Println("Total Score:", addTotalScores(calcTotalScore(framesTotals, bonusTotals)))
}
