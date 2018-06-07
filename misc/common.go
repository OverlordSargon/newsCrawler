package misc

import "errors"

func SplitLinks(links []string, workersNumber int) ([][]string, error) {
	if workersNumber < 1 {
		return nil, errors.New("workers number must be more than zero")
	}
	result := [][]string{}
	for i := 0; i < workersNumber; i++ {
		result = append(result, []string{})
	}
	counter := 0
	for _, link := range links {
		result[counter] = append(result[counter], link)
		counter++
		if counter == workersNumber {
			counter = 0
		}
	}
	return result, nil
}
