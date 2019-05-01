package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Point struct
{
	x, y, r float64
}

var points []Point = make([]Point, 0)

func Round(value float64) (newValue float64) {
	var round float64
	pow := math.Pow(10, float64(0))
	digit := pow * value
	_, div := math.Modf(digit)
	if div >= 0.5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newValue = round / pow
	return
}

func InRange(one Point, two Point) bool{
	dx := one.x - two.x
	dy := one.y - two.y
	distance := math.Sqrt(dx*dx + dy*dy)
	//fmt.Printf("\t%v\n", Round(distance))
	return (Round(distance) <= Round(one.r) || Round(distance) <= Round(two.r))
}

func AddPoint(newPoint Point) {

	var group []Point = make([]Point, 0)
	group = append(group, newPoint)
	
	k := 0
	
	for k < len(points){
		if InRange(newPoint, points[k]) {
			group = append(group, points[k])
			points = append(points[:k], points[k+1:]...)
		}else{
			k++
		}
	}
	
	//fmt.Printf("Group size: %v\n", len(group))
	
	if len(group) == 1 {
		points = append(points, newPoint)
	} else {
		var combined Point
		combined.x = 0.0
		combined.y = 0.0
		combined.r = 0.0
		for e := 0; e < len(group); e++{
			combined.x += group[e].x
			combined.y += group[e].y
			combined.r += group[e].r * group[e].r
		}
		combined.x = combined.x / float64(len(group))
		combined.y = combined.y / float64(len(group))
		combined.r = math.Sqrt(combined.r)
		AddPoint(combined)
	}
}

func main() {
	fileName := "crash.in"
	inFile, _ := os.Open(fileName)
	defer inFile.Close()
	
	reader := bufio.NewReader(inFile)
	scanner := bufio.NewScanner(reader)
	
	for scanner.Scan(){		
		point_count, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		
		if point_count == 0 {
			break
		}
		
		for i := 0; i < int(point_count); i++{
			var s Point
			
			scanner.Scan()
			
			reader2 := strings.NewReader(scanner.Text())			
			scanner2 := bufio.NewScanner(reader2)
			
			scanner2.Split(bufio.ScanWords)
			
			scanner2.Scan()
			x, _ := strconv.ParseFloat(scanner2.Text(), 64)
			
			scanner2.Scan()
			y, _ := strconv.ParseFloat(scanner2.Text(), 64)
			
			scanner2.Scan()
			r, _ := strconv.ParseFloat(scanner2.Text(), 64)
			
			//fmt.Printf("Index: %v ", i)
			//fmt.Printf("(%v, %v) : %v\n", x, y, r)
			//fmt.Printf("Going till: %v\n", int(point_count))
			
			s.x = x
			s.y = y
			s.r = r
			
			AddPoint(s)			
		}

		fmt.Println(len(points))
		
		points = make([]Point, 0)
	}
}