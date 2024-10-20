package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func main() {
	//learnDataTypes()

	//learnIfElseAndFunctions()

	//learnArraysAndSlices()

	//learnMap()

	//learnStrings()

	//learnStructsAndInterfaces()

	//learnPointers()

	//learnGoRoutines()

	//learnChannels()

	//learnGenerics()
}

func learnDataTypes() {
	var intNum uint8 = 255
	fmt.Println(intNum)

	const floatPi float64 = 3.1415926
	fmt.Println(floatPi)

	var printValue string = "Hello, world"
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func learnIfElseAndFunctions() {
	var numerator int = 10
	var denominator int = 3
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("The result if %v", remainder)
	} else {
		fmt.Printf("The result if %v, and the remainder is %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("denominator should not be 0")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func learnArraysAndSlices() {
	var intArr = [3]int32{1, 2, 3}
	fmt.Println(intArr)
	intArr[0] = 123
	fmt.Println(intArr)
	fmt.Println(intArr[1])
	fmt.Println(intArr[2:3])

	var intSlice = []int32{1, 2, 3}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)

	var intSlice2 = []int32{5, 6}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 = make([]int32, 4, 8)
	fmt.Println(intSlice3, len(intSlice3), cap(intSlice3))
}

func learnMap() {
	var myMap = map[string]uint8{"UID": 1, "Brush": 1}

	myMap["Age"] = 20

	var val, present = myMap["Age"]
	if present {
		fmt.Println(val)
	} else {
		fmt.Println("Not present")
	}

	delete(myMap, "Age")

	for key, val := range myMap {
		fmt.Println(key, val)
	}
}

func learnStrings() {
	var myString string = "résumé"

	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)

	for i, v := range myString {
		fmt.Println(i, v)
	}

	fmt.Printf("the length of résumé is %v\n", len(myString))

	// runes are alias to int32
	var myString2 = []rune("résumé")

	var indexed2 = myString2[0]
	fmt.Printf("%v, %T\n", indexed2, indexed2)

	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	fmt.Printf("the length of résumé (runes) is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("myRune = %v\n", myRune)

	var strSlice = []string{"s", "u", "r", "a", "j"}
	var catStr string
	for i := 0; i < len(strSlice); i++ {
		catStr += strSlice[i]
	}
	fmt.Println(catStr)

	// strings are immutable in GO

	// a faster way to do the above
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr2 = strBuilder.String()
	fmt.Println(catStr2)
}

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	// this is a method, not a function
	// (e gasEngine) assigns this function to the gasEngine type, making it a method
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("\nyou can make it!")
	} else {
		fmt.Println("\nyou need to fuel up!")
	}
}

func learnStructsAndInterfaces() {
	// structs are way of defining our own types
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 10, ownerInfo: owner{"Suraj"}, owner: owner{"Pooraj"}, int: 10}
	myEngine.gallons = 15
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.name, myEngine.int)
	fmt.Printf("total miles left in tank is %v", myEngine.milesLeft())
	canMakeIt(myEngine, 50)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("memory location of thing2 is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	return thing2
}

func square_pointer(thing2 *[5]float64) [5]float64 {
	fmt.Printf("memory location of thing2 is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}
	return *thing2
}

func learnPointers() {
	var p *int32 = new(int32) // if you don't initialize it, you will get error
	var i int32
	fmt.Printf("the value p points to is %v\n", *p)
	fmt.Printf("the value of i is %v\n", i)
	*p = 10
	fmt.Printf("the value p points to is %v\n", *p)
	p = &i
	fmt.Printf("the value p points to is %v\n", *p)
	*p = 1
	fmt.Printf("the value p points to is %v\n", *p)
	fmt.Printf("the value of i is %v\n", i)

	var Slice = []int32{1, 2, 3}
	var SliceCopy = Slice
	SliceCopy[0] = 4
	fmt.Println(Slice)
	fmt.Println(SliceCopy) // slices contain pointers to an underlying array

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("memory location of thing1 is: %p\n", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("result is %v\n", result)
	fmt.Printf("the value of thing1 is: %v\n", thing1)

	var result2 [5]float64 = square_pointer(&thing1)
	fmt.Printf("result2 is %v\n", result2)
	fmt.Printf("the value of thing2 is: %v\n", thing1)
}

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results []string

func learnGoRoutines() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		dbCall(i, false)
	}
	fmt.Printf("total execution time has been: %v\n", time.Since(t0))

	// now, lets do it concurrently
	results = []string{}
	t0 = time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i, true)
	}
	wg.Wait()
	fmt.Printf("total execution time has been: %v\n", time.Since(t0))
	fmt.Println("the results slice looks like this:", results)
}

func dbCall(i int, done bool) {
	// stimulate database call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	if done {
		wg.Done()
	}
}

func save(str string) {
	m.Lock()
	results = append(results, str)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("the current results are: ", results)
	m.RUnlock()
}

func learnChannels() {
	// channels -> are used to hold data (integers, slices, etc.), are thread safe (avoids data races), can listen for data so that we can block code execution until data is added or removed

	//var c = make(chan int)
	var c = make(chan int, 5)
	go process(c)
	for i := range c {
		fmt.Println(i)
	}

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrice(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func process(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
	}
	fmt.Println("process end")
}

var maxChickenPrice float32 = 5
var maxTofuPrice float32 = 3

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= maxChickenPrice {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrice(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= maxTofuPrice {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nfound deal on chicken at %v", website)
	case website := <-tofuChannel:
		fmt.Printf("\nfound deal on tofu at %v", website)
	}
}

func learnGenerics() {
	var intSlice = []int{1, 2, 3}
	fmt.Println(sum[int](intSlice))

	var float32Slice = []float32{1.1, 2.2, 3.3}
	fmt.Println(sum[float32](float32Slice))

	fmt.Println(isEmpty(float32Slice))
}

func sum[T int | float32 | float64](slice []T) T {
	var tmp T
	for _, v := range slice {
		tmp += v
	}
	return tmp
}

func isEmpty[T any](slice []T) bool {
	if len(slice) == 0 {
		return true
	} else {
		return false
	}
}
