package main

/*
 * EXERCISE_2.1
 * Add types, constants and functions to tempconv for processing temperatures
 * in the kelvin scale, where kelvine is -273.15°C and a difference of 1K has
 * the same magnitude as 1°C
*/

import "fmt"

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CtoK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtoK(f Fahrenheit) Kelvin {
	return Kelvin((f + 459.67) * 5/9)
}

func KtoC (k Kelvin) Celsius {
	return Celsius(k + Kelvin(ZeroKelvin))
}

func KtoF (k Kelvin) Fahrenheit {
	return Fahrenheit(k * 9/5 - 459.67)
}

func main() {
	fmt.Printf("50°F = %g°C\n", FToC(Fahrenheit(50)))
	fmt.Printf("50°C = %g°F\n", CToF(Celsius(50)))

	exampleNumber := 35
	fmt.Printf("" +
		"Nesting the two conversions should result in\n" +
		"the same output as what was input as the \n" +
		"operations are being reversed. Inputting \n" +
		"%d° should result in the same number - %g°\n", exampleNumber,
			FToC(CToF(Celsius(exampleNumber))))

	fmt.Printf("50°C = %gK\n", CtoK(50))
	fmt.Printf("50°F = %gK\n", FtoK(50))
	fmt.Printf("50°K = %gC\n", KtoC(50))
	fmt.Printf("50°K = %gF\n", KtoF(50))
}


/*
====================================================================
						EXAMPLE OUTPUT
====================================================================

50°F = 10°C
50°C = 122°F
Nesting the two conversions should result in
the same output as what was input as the
operations are being reversed. Inputting
35° should result in the same number - 35°

----------

++Kelvin exercise results++

50°C = 323.15K
50°F = 283.15K
50°K = -223.14999999999998C
50°K = -369.67F


 */