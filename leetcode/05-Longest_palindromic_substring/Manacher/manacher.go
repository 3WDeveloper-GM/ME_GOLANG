// This is my implementation of the longest palindromic substring algorithm 
// I learned about the Manacher Algortithm, and decided it was a good idea to
// implement it.

package manacher

import (
	"fmt"
	lon "longest/stringmodif"

	"golang.org/x/exp/slices"
)

func Manacher(chain string) int {
	newchain := lon.Modifystring(chain)
	runes := []rune(newchain)
	Palindromecenterarray := make([]int, len(runes))
	Center := 0
	Radii := Center

	for Center < len(runes) {
		for Center-(Radii+1) >= 0 && Center+(Radii+1) < len(runes) && runes[Center-(Radii+1)] == runes[Center+(Radii+1)] {
			Radii++
		}
		Palindromecenterarray[Center] = Radii

		OldCenter := Center
		OldRadii := Radii
		Center++
		Radii = 0

		for Center <= OldCenter+OldRadii {
			MirrorCentre := 2*OldCenter - Center
			MaxMirrorRadii := OldCenter + OldRadii - Center
			if Palindromecenterarray[MirrorCentre] < MaxMirrorRadii {
				Palindromecenterarray[Center] = Palindromecenterarray[MirrorCentre]
				Center++
			} else if Palindromecenterarray[MirrorCentre] > MaxMirrorRadii {
				Palindromecenterarray[Center] = MaxMirrorRadii
				Center++
			} else {
				Radii = MaxMirrorRadii
				break
			}
		}
	}
	Long := slices.Max(Palindromecenterarray)
	fmt.Println(Palindromecenterarray)
	return Long
}
