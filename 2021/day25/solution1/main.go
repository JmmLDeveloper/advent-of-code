package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func cord_with_loop( yi, xi, rows, cols int) (yo ,xo int) {
	if yi == rows {
		yo = 0
	} else {
		yo = yi
	} 
	if xi == cols {
		xo = 0
	} else {
		xo = xi
	}

	return
}

func step(board [][]rune) (moves int) {
	
	rows := len(board)
	cols := len(board[0])

	for _,herd := range []rune{'>','v'} {
		for _,phase := range []string{"checking","moving"}  {
			for i := 0; i < rows ; i++ {
				for j := 0 ; j < cols ; j++ {
					if phase == "checking"   {
						x,y := -1,-1 
						if herd == '>' && board[i][j] == '>' {
							y,x = cord_with_loop(i,j + 1,rows,cols)
						} else if herd == 'v' && board[i][j] == 'v' {
							y,x = cord_with_loop(i + 1,j,rows,cols) 
						}
						if x != -1 && y != -1 && board[y][x] == '.' {
							moves++;
							board[y][x] = '*'
							board[i][j] = '-'	
						}
					} else if phase == "moving" {
						switch board[i][j] {
							case '-' :
								board[i][j] = '.'
							case '*' :
								board[i][j] = herd
						} 
					}
				}
			}
		}
	} 
	return
}

func print_board(board [][]rune) {
	for i := 0 ; i < len(board) ;i++ {
		for j := 0 ; j < len(board[0]) ; j++ {
			fmt.Printf("%c",board[i][j])
		}
		fmt.Println("")
	}
}

func main() {
	b, err :=  ioutil.ReadFile("input.txt");
	if err != nil {
		panic(err)
	}

	input := string(b)

	input_rows := strings.Split(input,"\n");

	for i := range input_rows {
		input_rows[i] = strings.Trim( input_rows[i] , "\r" )
	}

	rows := len(input_rows)
	cols := len(input_rows[0]) 

	board := make([][]rune,rows);


	for i := 0 ; i < rows ; i++ {
		board[i] = make([]rune,cols);
		for j,ch := range input_rows[i] {
			board[i][j] = ch
		}
	}

	var counter int
	for counter = 1 ; step(board) != 0 ; counter++ {}
	
	fmt.Printf("The answer to part 1 is : %d",counter)

	
	

}

