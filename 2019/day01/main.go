package main

import "fmt"

/*
Santa has become stranded at the edge of the Solar System while delivering presents to other planets! To accurately calculate his position in space, safely align his warp drive, and return to Earth in time to save Christmas, he needs you to bring him measurements from fifty stars.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

The Elves quickly load you into a spacecraft and prepare to launch.

At the first Go / No Go poll, every Elf is Go until the Fuel Counter-Upper. They haven't determined the amount of fuel required yet.

Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.

For example:

For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
For a mass of 1969, the fuel required is 654.
For a mass of 100756, the fuel required is 33583.

The Fuel Counter-Upper needs to know the total fuel requirement. To find it, individually calculate the fuel needed for the mass of each module (your puzzle input), then add together all the fuel values.

What is the sum of the fuel requirements for all of the modules on your spacecraft?
*/

func main() {
	fmt.Println("Part 1", totalFuelRequired(modules, basicFuelRequired))
	fmt.Println("Part 1", totalFuelRequired(modules, cumulativeFuelRequired))
}

/* Part 1

   Fuel required to launch a given module is based on its mass. Specifically, to find the fuel required for a module, take its mass, divide by three, round down, and subtract 2.
*/
func basicFuelRequired(mass int64) int64 {
	return mass/3 - 2
}

/* Part 2
   Fuel itself requires fuel just like a module - take its mass, divide by three, round down, and subtract 2. However, that fuel also requires fuel, and that fuel requires fuel, and so on. Any mass that would require negative fuel should instead be treated as if it requires zero fuel; the remaining mass, if any, is instead handled by wishing really hard, which has no mass and is outside the scope of this calculation.

   So, for each module mass, calculate its fuel and add it to the total. Then, treat the fuel amount you just calculated as the input mass and repeat the process, continuing until a fuel requirement is zero or negative. For example:

   A module of mass 14 requires 2 fuel. This fuel requires no further fuel (2 divided by 3 and rounded down is 0, which would call for a negative fuel), so the total fuel required is still just 2.
   At first, a module of mass 1969 requires 654 fuel. Then, this fuel requires 216 more fuel (654 / 3 - 2). 216 then requires 70 more fuel, which requires 21 fuel, which requires 5 fuel, which requires no further fuel. So, the total fuel required for a module of mass 1969 is 654 + 216 + 70 + 21 + 5 = 966.
   The fuel required by a module of mass 100756 and its fuel is: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.

   What is the sum of the fuel requirements for all of the modules on your spacecraft when also taking into account the mass of the added fuel? (Calculate the fuel requirements for each module separately, then add them all up at the end.)
*/
func cumulativeFuelRequired(mass int64) int64 {
	var total int64
	for {
		fuel := basicFuelRequired(mass)
		if fuel <= 0 {
			return total
		}

		total += fuel
		mass = fuel
	}
}

func totalFuelRequired(modules []int64, fuelRequired func(int64) int64) int64 {
	var total int64
	for _, module := range modules {
		total += fuelRequired(module)
		if total < 0 {
			panic("int64 overflow")
		}
	}
	return total
}

var modules = []int64{
	50951,
	69212,
	119076,
	124303,
	95335,
	65069,
	109778,
	113786,
	124821,
	103423,
	128775,
	111918,
	138158,
	141455,
	92800,
	50908,
	107279,
	77352,
	129442,
	60097,
	84670,
	143682,
	104335,
	105729,
	87948,
	59542,
	81481,
	147508,
	62687,
	64212,
	66794,
	99506,
	137804,
	135065,
	135748,
	110879,
	114412,
	120414,
	72723,
	50412,
	124079,
	57885,
	95601,
	74974,
	69000,
	66567,
	118274,
	136432,
	110395,
	88893,
	124962,
	74296,
	106148,
	59764,
	123059,
	106473,
	50725,
	116256,
	80314,
	60965,
	134002,
	53389,
	82528,
	144323,
	87791,
	128288,
	109929,
	64373,
	114510,
	116897,
	84697,
	75358,
	109246,
	110681,
	94543,
	92590,
	69865,
	83912,
	124275,
	94276,
	98210,
	69752,
	100315,
	142879,
	94783,
	111939,
	64170,
	83629,
	138743,
	141238,
	77068,
	119299,
	81095,
	96515,
	126853,
	87563,
	101299,
	130240,
	62693,
	139018,
}
