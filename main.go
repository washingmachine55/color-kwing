package main

import (
	"color-kwing/api"
)

func main() {
	// var hexCodes []string
	// for _, row := range ArrOffColors {
	// 	// Check if the row actually has a second element to prevent a panic
	// 	if len(row) > 1 {
	// 		hexCodes = append(hexCodes, row[1])
	// 	}
	// }
	// var reader io.Reader
	var data string = "#FFBF01"

	// _, err := fmt.Scanln(&data)
	// if err != nil {
	// 	fmt.Errorf("%s", err)
	// }

	if data != "" {
		val, err := api.NewRgbFromHex(data)
		if err != nil {
			api.Logg.Error().Err(err).Msg("[Error Occurred]")
		}

		// colorName, err := api.GetColorName(data)
		// if err != nil {
		// 	err.Error()
		// }
		//
		colorName, err := api.ClosestMatch(data)
		if err != nil {
			err.Error()
		}

		api.Logg.Info().Any("val", val).Msg("[FINAL VALUE]")
		api.Logg.Info().Any("Color Name:", colorName).Msg("[FINAL VALUE]")
	}
}
