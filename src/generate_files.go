package main

import (
	"math/rand"
	"time"
	"os"
	"log"
)

func Init() {
    rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ,.!?><-+=:;|][")
var word = "Hello world!\nGo is the best language in the world???"

func RandStringRunes(n int) string {
	// generate a string with designated length
	// reference: https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

func max(n1 int, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}

func GenerateFile(word_count int, line_count int, filename string) {
	// function to make random files with designated count of lines and desired grep -c result
	// for the potential queries
	if _, err := os.Stat(filename); err == nil {
		os.Remove(filename)
	}
	remain_word_count := word_count
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0755)
	if word_count == 0 {
		f.WriteString(RandStringRunes(100 * line_count))
		f.Close()
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < line_count; i++ {
		insert_word_count := rand.Intn(max(1, word_count / line_count))
		remain_word_count -= insert_word_count
		if insert_word_count != 0 {
			for j := 0; j < insert_word_count; j++ {
				text := RandStringRunes(rand.Intn(50))
				text += word
				text += RandStringRunes(rand.Intn(50)) + "\n"
				f.WriteString(text)
			}
		} else {
			text := RandStringRunes(rand.Intn(100)) + "\n"
			f.WriteString(text)
		}
	}

	if remain_word_count != 0 {
		for {
			f.WriteString(word + "\n")
			remain_word_count -= 1
			if remain_word_count == 0 {
				break
			}
		}
	}
	log.Println("File created!")
	f.Close()
}

func main() {
	Init()

	// we control the lines to be the same and change the frequency of the expression in our created files
	GenerateFile(100, 1000000, "/home/hangy6/test_logs/low_freq.log") //create low frequency
	GenerateFile(10000, 1000000, "/home/hangy6/test_logs/medium_freq.log") //create medium frequency
	GenerateFile(10000000, 1000000, "/home/hangy6/test_logs/high_freq.log") //create high frequency
	
	// we control the frequency to be the same and compare between different file size
	GenerateFile(3, 10, "/home/hangy6/test_logs/mini.log") //create mini file
	GenerateFile(30, 100, "/home/hangy6/test_logs/small.log") //create small file
	GenerateFile(300, 1000, "/home/hangy6/test_logs/medium.log") //create medium file
	GenerateFile(30000, 100000, "/home/hangy6/test_logs/large.log") //create large file
	GenerateFile(300000, 1000000, "/home/hangy6/test_logs/huge.log") //create huge file
}

