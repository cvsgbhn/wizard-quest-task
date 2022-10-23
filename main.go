package main

import (
	"encoding/json"
	"fmt"
)

type Room struct {
	Visited  bool
	PastPath []string
	NextMaze interface{}
	NextPath map[string]interface{}
}

func isRoom(node interface{}) (Room, bool) {
	res, ok := node.(map[string]interface{})
	if !ok {
		return Room{}, ok
	}

	return Room{
		Visited:  false,
		PastPath: []string{},
		NextMaze: res,
	}, ok
}

func isExit(node interface{}) bool {
	fmt.Println("isExit")
	a, ok := node.(string)
	if ok && len(a) > 0 && a == "exit" {
		return true
	}

	return false
}

func visitNode(r *Room, pathPart []string) {
	r.PastPath = append(r.PastPath, pathPart...)
}

func BFS(queue []Room) []string {
	newQ := make([]Room, 0)
	for len(queue) > 0 {
		fmt.Println("newQ before new checks: ", newQ)
		newQ = newQ[:0]

		fmt.Println("queue: ", queue)
		for i := range queue {
			if queue[i].Visited == true {
				continue
			}

			queue[i].Visited = true

			queue[i].NextPath, _ = queue[i].NextMaze.(map[string]interface{})

			for k, v := range queue[i].NextPath {
				path := append(queue[i].PastPath, k)
				r, ok := isRoom(v)
				if !ok && isExit(v) {
					return path
				}

				if r.NextMaze != nil {
					visitNode(&r, path)
					newQ = append(newQ, r)
				}
			}
		}

		queue = append(queue, newQ...)
		if queue[0].Visited {
			queue = queue[1:]
		}

		fmt.Println(queue)
	}

	return nil
}

//func MazeSolver(maze map[string]interface{}) string {
//
//}

func main() {

	example := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`

	maze := Room{}

	json.Unmarshal([]byte(example), &maze.NextMaze)

	fmt.Println(maze)

	rooms := []Room{maze}

	res := BFS(rooms)
	fmt.Println("Path to exit: ", res)
}
