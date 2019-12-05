package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ttacon/libphonenumber"
)

var phones = []string{
	"1",
	"49",
	"0044",
	"+1-11111111",
	"+498963648018",
	"+4919636480",
	"+1 (242) 397-7400", // bahamas
	"+86 10 6552 9988",  // china
	"+35988741388",
	"+359-988-723-023",
	"+359-988-954-650",
}

func main() {
	for _, p := range phones {
		ok := validateNumber(p)
		//fmt.Printf("valid: %v phone %v\n", ok, p)
		if !ok {
			fmt.Printf("invalid: %v\n", p)
		}
	}
	// not good
	//for _, p := range phones {
	//	ok := validateNumber2(p)
	//	//fmt.Printf("valid: %v phone %v\n", ok, p)
	//	if !ok {
	//		fmt.Printf("invalid2: %v\n", p)
	//	}
	//}
}

func validateNumber2(number string) bool {
	_, err := libphonenumber.Parse(number, "US")
	return err != nil
}

//validateNumber ensures that the given provided phone number is following
//the rules of E.164. We assume that:
// A) The number starts with a `+``,
// B) There are no more than 15 digits, and
// C) There is a country code.
//First, we ensure the first character is a +. We then attempt to match the
//country code
// uses libphonenumber
//
func validateNumber(number string) bool {
	reg, err := regexp.Compile("[^0-9+]+")
	if err != nil {
		return false
	}
	number = reg.ReplaceAllString(number, "")

	phoneNumber := []rune(number)

	//Empty check
	if len(phoneNumber) == 0 {
		return false
	} else if len(phoneNumber) > 16 {
		//At least 15 digits with a +
		return false
	} else if phoneNumber[0] != '+' {
		//Ensure the first character is a '+'
		return false
	}

	//Isolate the country code
	for i := 1; i <= 3; i++ {
		countryCode, err := strconv.Atoi(string(phoneNumber[1 : 1+i]))
		//If we failed to convert the country code for some reason to an int,
		//move on with the next batch
		if err != nil {
			continue
		}
		//fmt.Printf("countryCode %v\n", countryCode)
		regions, ok := libphonenumber.CountryCodeToRegion[countryCode]
		//If the region wasn't found, continue
		if !ok || regions == nil || len(regions) == 0 {
			continue
		}
		//fmt.Printf("regions %+v\n", regions)

		//Now we validate the number against that region
		//Note that we only need the first region, as libphonenumber
		//checks against all possible regions
		parsedNumber, err := libphonenumber.Parse(number, regions[0])
		//again, if err, move on
		if err != nil {
			continue
		}

		if libphonenumber.IsValidNumber(parsedNumber) {
			return true
		}
	}

	//If we failed to validate by now, it's not a phone number
	return false
}
