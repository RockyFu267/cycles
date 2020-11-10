package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Actions struct {
	Name  string
	Value string
}

func main() {
	fmt.Println("shit")
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		switch input.Text() {
		case "level1":
			fmt.Println("fuck")
			var playwin, aiwin, winwin int
			playwin = 0
			aiwin = 0
			winwin = 0
			for i := 0; i < 10; i++ {
				input.Scan()
				switch input.Text() {
				case "1":
					var playActions Actions
					playActions.Name = "you"
					playActions.Value = "1"
					playres := TypeNameRes(playActions.Value)
					fmt.Println(playActions.Name + "选择出" + playres)
					airestmp := AIAction()
					var AIActions Actions
					AIActions.Name = "AI"
					AIActions.Value = airestmp
					aires := TypeNameRes(AIActions.Value)
					fmt.Println(AIActions.Name + "选择出" + aires)
					res := JudgeRes(playActions, AIActions)
					if len(res) > 1 {
						fmt.Println("双方平局")
						winwin = winwin + 1
					} else {
						fmt.Println("winner is " + res[0].Name)
						if res[0].Name == "you" {
							playwin = playwin + 1
						} else {
							aiwin = aiwin + 1
						}
					}

				case "2":
					var playActions Actions
					playActions.Name = "you"
					playActions.Value = "2"
					playres := TypeNameRes(playActions.Value)
					fmt.Println(playActions.Name + "选择出" + playres)
					airestmp := AIAction()
					var AIActions Actions
					AIActions.Name = "AI"
					AIActions.Value = airestmp
					aires := TypeNameRes(AIActions.Value)
					fmt.Println(AIActions.Name + "选择出" + aires)
					res := JudgeRes(playActions, AIActions)
					if len(res) > 1 {
						fmt.Println("双方平局")
						winwin = winwin + 1
					} else {
						fmt.Println("winner is " + res[0].Name)
						if res[0].Name == "you" {
							playwin = playwin + 1
						} else {
							aiwin = aiwin + 1
						}
					}
				case "3":
					var playActions Actions
					playActions.Name = "you"
					playActions.Value = "3"
					playres := TypeNameRes(playActions.Value)
					fmt.Println(playActions.Name + "选择出" + playres)
					airestmp := AIAction()
					var AIActions Actions
					AIActions.Name = "AI"
					AIActions.Value = airestmp
					aires := TypeNameRes(AIActions.Value)
					fmt.Println(AIActions.Name + "选择出" + aires)
					res := JudgeRes(playActions, AIActions)
					if len(res) > 1 {
						fmt.Println("双方平局")
						winwin = winwin + 1
					} else {
						fmt.Println("winner is " + res[0].Name)
						if res[0].Name == "you" {
							playwin = playwin + 1
						} else {
							aiwin = aiwin + 1
						}
					}
				default:
					fmt.Println("无效指令")
					i = i - 1
				}
			}
			if playwin > aiwin {
				fmt.Println("最终的赢家是你")
				fmt.Println("比分:\n", playwin, ":", aiwin)
				fmt.Println("平局:", winwin)
			}
			if playwin < aiwin {
				fmt.Println("最终的赢家是AI")
				fmt.Println("比分:\n", playwin, ":", aiwin)
				fmt.Println("平局:", winwin)
			}
			if playwin == aiwin {
				fmt.Println("双方平局")
				fmt.Println("比分:\n", playwin, ":", aiwin)
				fmt.Println("平局:", winwin)
			}
		default:
			fmt.Println("retry")
		}
	}
}

func Game(input *bufio.Scanner) {

}

func AIAction() (output string) {
	rand.Seed(time.Now().UnixNano())
	var intlist []string
	intlist = append(intlist, "1", "2", "3")
	x := rand.Intn(3)
	return intlist[x]
}

func TypeNameRes(Input string) (OutPut string) {
	switch Input {
	case "1":
		return "石头/rock"
	case "2":
		return "剪刀/scissors"
	case "3":
		return "布/paper"
	default:
		return ""
	}
}

func JudgeRes(play1 Actions, play2 Actions) (winner []Actions) {
	switch play1.Value {
	case "1":
		switch play2.Value {
		case "1":
			winner = append(winner, play1, play2)
			return winner
		case "2":
			winner = append(winner, play1)
			return winner
		case "3":
			winner = append(winner, play2)
			return winner
		}
	case "2":
		switch play2.Value {
		case "1":
			winner = append(winner, play2)
			return winner
		case "2":
			winner = append(winner, play1, play2)
			return winner
		case "3":
			winner = append(winner, play1)
			return winner
		}
	case "3":
		switch play2.Value {
		case "1":
			winner = append(winner, play1)
			return winner
		case "2":
			winner = append(winner, play2)
			return winner
		case "3":
			winner = append(winner, play1, play2)
			return winner
		}
	default:
		return winner
	}

	return winner
}
