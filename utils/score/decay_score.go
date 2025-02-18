package score

import (
	"math"
	"time"
)

type ScoreStore struct {
	Num   float64 // for performance i didn't use math/big rationale arithmetic
	Denom float64
	Time  time.Time
}

func NewScoreStore(num float64, denom float64, inpTime time.Time) ScoreStore {
	return ScoreStore{Num: num, Denom: denom, Time: inpTime}
}

// CalculateTimeDecayFunctionUpdate calculates the time decayed score update between two ScoreStore entries.
// It uses a decay function with a half life of halfLife to factor in the time elapsed since the oldScore was recorded.
// Both the numerator and the denominator of the newScore are decayed by this function.
// Additionally, the newScore is factored by a weight of updateWeight.
// The function returns a new ScoreStore entry with the updated numerator, denominator, and current time.
//
// The mathematical equation used to calculate the update is:
//
//	updatedNum = oldScore.Num*exp(-(now-oldScore.Time)/halfLife) + newScore.Num*exp(-(now-newScore.Time)/halfLife)*updateWeight
//	updatedDenom = oldScore.Denom*exp(-(now-oldScore.Time)/halfLife) + newScore.Denom*exp(-(now-newScore.Time)/halfLife)*updateWeight
//
// where now is the current time.
//
// Note that the returned ScoreStore has a new Time field set to the current time.
func CalculateTimeDecayFunctionUpdate(oldScore ScoreStore, newScore ScoreStore, halfLife time.Duration, updateWeight float64) ScoreStore {
	oldDecayExponent := math.Ln2 * time.Since(oldScore.Time).Seconds() / halfLife.Seconds()
	oldDecayFactor := math.Exp(-oldDecayExponent)
	newDecayExponent := math.Ln2 * time.Since(newScore.Time).Seconds() / halfLife.Seconds()
	newDecayFactor := math.Exp(-newDecayExponent)
	updatedNum := oldScore.Num*oldDecayFactor + newScore.Num*newDecayFactor*updateWeight
	updatedDenom := oldScore.Denom*oldDecayFactor + newScore.Denom*newDecayFactor*updateWeight
	return NewScoreStore(updatedNum, updatedDenom, time.Now())
}
