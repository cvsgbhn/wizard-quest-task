package maze

import "testing"

func TestMazeSolver(t *testing.T) {
	type args struct {
		maze string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1 from task text",
			args: args{maze: `{"forward": "tiger", "left": {"forward": {"upstairs": "exit"}, "left": "dragon"}, "right": {"forward": "dead end"}}`},
			want: "[\"left\", \"forward\", \"upstairs\"]",
		},
		{
			name: "Case 2 from task text",
			args: args{maze: `{"forward": "exit"}`},
			want: "[\"forward\"]",
		},
		{
			name: "Case 3 from task text",
			args: args{maze: `{"forward": "tiger", "left": "ogre", "right": "demon"}`},
			want: "Sorry",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MazeSolver(tt.args.maze); got != tt.want {
				t.Errorf("MazeSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}
