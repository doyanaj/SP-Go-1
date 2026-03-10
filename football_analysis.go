package main

import "sort"

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses != 0 {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	} else {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	}
}

func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	p.calculateRating()
	return p
}

func goalsSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)

	sort.Slice(res, func(i, j int) bool {
		if res[i].Goals != res[j].Goals {
			return res[i].Goals > res[j].Goals
		}
		return res[i].Name < res[j].Name
	})

	return res
}

func ratingSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)

	sort.Slice(res, func(i, j int) bool {
		if res[i].Rating != res[j].Rating {
			return res[i].Rating > res[j].Rating
		}
		return res[i].Name < res[j].Name
	})

	return res
}

func gmSort(players []Player) []Player {
	res := make([]Player, len(players))
	copy(res, players)

	sort.Slice(res, func(i, j int) bool {
		var p1, p2 float64

		if res[i].Misses != 0 {
			p1 = float64(res[i].Goals) / float64(res[i].Misses)
		} else {
			p1 = float64(res[i].Goals)
		}

		if res[j].Misses != 0 {
			p2 = float64(res[j].Goals) / float64(res[j].Misses)
		} else {
			p2 = float64(res[j].Goals)
		}

		if p1 != p2 {
			return p1 > p2
		}
		return res[i].Name < res[j].Name
	})

	return res
}
