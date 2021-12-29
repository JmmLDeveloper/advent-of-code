package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Reindeer struct {
	speed ,fly_time , rest_time , energy , distance , score int // distance if for the first problem , score if for the second
	name string
}

func NewReindeer(name string, speed int, fly_time int,rest_time int) Reindeer  {
	return Reindeer{
		name : name,
		speed: speed,
		fly_time : fly_time,
		rest_time : rest_time,
		// positive if flying negative if resting , this values jumps from 0 to fly_time or to rest_time when the other time is over
		//energy can never be 0
		energy : fly_time, 
		distance: 0, 
	}
}

func (r *Reindeer) RunsSecond() {

	if r.energy == 1 {
		r.energy = -r.rest_time;
		r.distance += r.speed
	} else if ( r.energy == -1 ){
		r.energy = r.fly_time
	} else if ( r.energy > 0 ) {
		r.distance += r.speed
		r.energy--
	} else {
		r.energy++;
	}

}  

func MustAtoi(str string) int {
	number,err := strconv.Atoi(str)


	if err != nil {
		panic(err)
	}

	return number
}

func ParseFile() []Reindeer {

	
	input_bytes,err := os.ReadFile("./input.txt")

	input := string(input_bytes)

	lines := strings.Split(input,"\n")

	if err != nil {
		panic(err)
	}

	// this is horrible but HEY!, i just want to practice
	name_regex := regexp.MustCompile(`^\S*`,)
	speed_regex := regexp.MustCompile(`fly (\d+)`)
	time_regex := regexp.MustCompile(`(\d*) seconds`)
	reindeers := make( []Reindeer, len(lines) )
	
	for i := range lines {
		name := name_regex.FindString(lines[i])
		speed := MustAtoi(speed_regex.FindString(lines[i])[4:])

		times := time_regex.FindAllString(lines[i],2)
		fly_time := MustAtoi(strings.TrimSuffix( times[0] , " seconds" ))
		rest_time := MustAtoi(strings.TrimSuffix( times[1] , " seconds" ))

		reindeers[i] = NewReindeer(name,speed,fly_time,rest_time)
//		fmt.Println(reindeers[i])
	}



	return reindeers
}

func (r Reindeer) String () string {
	return fmt.Sprintf(" { speed : %d , name : %s , rest_time : %d , fly_time : %d} ",r.speed,r.name,r.rest_time,r.fly_time)
}

func ReindeersWithMoreDistance(  reindeers []Reindeer  ) []int {

	best_reindeers := make([]int,0)


	best_distance := -1;

	for _,r := range reindeers {
		if r.distance > best_distance {
			best_distance = r.distance
		}
	}

	for idx,r := range reindeers {
		if r.distance == best_distance {
			best_reindeers = append(best_reindeers,idx)
		}
	}
	
	return best_reindeers
}



func main() {

	const SECONDS int = 2503

	reindeers := ParseFile()

	for s := 0 ; s < SECONDS ; s++ {
		for idx := range reindeers {
			reindeers[idx].RunsSecond()
		}
		for _,idx := range ReindeersWithMoreDistance(reindeers) {
			reindeers[idx].score++;
		} 
	}

	best_score := -1

	for _,r := range reindeers {
		if r.score > best_score {
			best_score = r.score
		}
	}

	


	

	fmt.Printf("the answer of part 1 is : %d\n", reindeers[ ReindeersWithMoreDistance(reindeers)[0] ].distance )
	fmt.Printf("the answer of part 2 is : %d", best_score )


}