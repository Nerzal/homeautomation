package main

import "encoding/json"

type ColorChange struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	c := ColorChange{
		Red:   255,
		Green: 0,
		Blue:  0,
	}

	bytes, err := json.Marshal(c)
	if err != nil {
		println(err)
		return
	}

	var result ColorChange

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		println(err)
		return
	}

	if result.Red != 255 {
		panic("FML")
	}
}
