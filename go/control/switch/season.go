package main

import "fmt"

func main() {
	seasonName, err := season(12)
	if err != nil {
		fmt.Printf("发生错误：%s", err)
	} else {
		fmt.Println(seasonName)
	}
}

func season(season int) (seasonName string, err error) {
	switch {
	case season >= 1 && season <= 3:
		seasonName = "春天"
	case season >= 4 && season <= 6:
		seasonName = "夏天"
	case season >= 7 && season <= 9:
		seasonName = "秋天"
	case season == 10 || season == 11 || season == 12:
		seasonName = "冬天"
	default:
		err = fmt.Errorf("非法的季节: %d", season)
	}
	return seasonName, err
}