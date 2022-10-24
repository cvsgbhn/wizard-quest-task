package maze

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
	a, ok := node.(string)
	if ok && len(a) > 0 && a == "exit" {
		return true
	}

	return false
}

func addPath(r *Room, pathPart []string) {
	r.PastPath = append(r.PastPath, pathPart...)
}

func BFS(queue []Room) []string {
	newQ := make([]Room, 0)
	for len(queue) > 0 {
		newQ = newQ[:0]

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
					addPath(&r, path)
					newQ = append(newQ, r)
				}
			}
		}

		queue = append(queue, newQ...)
		if queue[0].Visited {
			queue = queue[1:]
		}
	}

	return nil
}

func MazeSolver(maze string) string {
	r := Room{}

	json.Unmarshal([]byte(maze), &r.NextMaze)

	rooms := []Room{r}

	res := BFS(rooms)

	if len(res) == 0 {
		fmt.Println("Sorry")
		return "Sorry"
	}

	return formatOutput(res)
}

func formatOutput(ans []string) string {
	s := "["

	for i, v := range ans {
		s += "\"" + v + "\""
		if i < len(ans)-1 {
			s += ", "
		}
	}

	s += "]"

	return s
}

//func main() {
//
//	example := `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`
//	//example := `{"forward": "exit"}`
//	//example := `{"forward": "tiger", "left": "ogre", "right": "demon"}`
//
//	res := MazeSolver(example)
//
//	fmt.Println(res)
//}
