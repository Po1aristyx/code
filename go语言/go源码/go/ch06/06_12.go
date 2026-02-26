package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	randTest()
	randAnswer()
}
func randTest() {
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Float64())
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(10)
	fmt.Println(randnum)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float64())
	num := rand.Intn(7) + 5
	fmt.Println(num)
}
func randAnswer() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	rand.Seed(time.Now().UnixNano())
	randnum := rand.Intn(len(answers))
	fmt.Println("随机回答", answers[randnum])
}
