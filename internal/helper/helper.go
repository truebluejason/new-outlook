package helper

import (
	"fmt"
)

func DescribeEmotion(intent string, symptoms []string) string {
	var concat string
	for i,symptom := range symptoms {
		if i == len(symptoms) - 1 {
			concat = concat + "and " + symptom
		}
		concat = concat + symptom + ", "
	}
	formatted := fmt.Sprintf("The sensation of %v is characterized by the presence of %v. ", intent, symptoms)
	return formatted
}
