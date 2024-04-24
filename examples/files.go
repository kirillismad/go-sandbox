package examples

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const FILENAME = "./src/fileExample.txt"

var content []string = []string{
	"One thing, I don't know why",
	"It doesn't even matter how hard you try",
	"Keep that in mind, I designed this rhyme",
	"To remind myself how I tried so hard",
	"In spite of the way you were mockin' me",
	"Actin' like I was part of your property",
	"Remembering all the times you fought with me",
	"I'm surprised it got so far",
	"Things aren't the way they were before",
	"You wouldn't even recognize me anymore",
	"Not that you knew me back then",
	"But it all comes back to me in the end",
}

var additionalContent []string = []string{
	"You kept everything inside",
	"And even though I tried, it all fell apart",
	"What it meant to me will eventually be",
	"A memory of a time when I tried so hard",
}

func DemoFiles() {
	os.Remove(FILENAME)
	demoCreateFileUsingOs()
	demoCreateFileUsingBufio()
	demoAppendFileUsingOs()
	demoReadFileUsingOs()
	demoReadFileUsingScan()
	demoReadFileUsingWriteTo()
}

func panicError(err error) {
	if err != nil {
		panic(err)
	}
}

func demoCreateFileUsingOs() {
	fmt.Println(newline + fname(demoCreateFileUsingOs))
	file, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	panicError(err)
	defer file.Close()

	contentString := time.Now().Local().Format(time.RFC3339) + "\n\n"
	contentString += strings.Join(content, "\n") + "\n"

	n, err := file.WriteString(contentString)
	fmt.Println(n, "bytes have been written")
	panicError(err)
}

func demoCreateFileUsingBufio() {
	fmt.Println(newline + fname(demoCreateFileUsingBufio))

	file, err := os.OpenFile(FILENAME, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	panicError(err)
	defer file.Close()

	contentString := time.Now().Local().Format(time.RFC3339) + "\n\n"
	contentString += strings.Join(content, "\n") + "\n"

	writer := bufio.NewWriter(file)
	n, err := writer.WriteString(contentString)
	panicError(err)
	err = writer.Flush()
	panicError(err)
	fmt.Println(n, "bytes have been written")
}

func demoAppendFileUsingOs() {
	fmt.Println(newline + fname(demoAppendFileUsingOs))
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_WRONLY, 0644)
	panicError(err)
	defer file.Close()

	contentString := strings.Join(additionalContent, "\n") + "\n"
	n, err := file.WriteString(contentString)
	panicError(err)
	fmt.Println(n, "bytes have been written")
}

func demoReadFileUsingOs() {
	fmt.Println(newline + fname(demoReadFileUsingOs))
	file, err := os.OpenFile(FILENAME, os.O_RDONLY, 0644)
	panicError(err)
	defer file.Close()

	result, err := io.ReadAll(file)
	panicError(err)
	fmt.Println(string(result))
}

func demoReadFileUsingScan() {
	fmt.Println(newline + fname(demoReadFileUsingScan))

	file, err := os.OpenFile(FILENAME, os.O_RDONLY, 0644)
	panicError(err)
	defer file.Close()

	scanner, stringBuilder := bufio.NewScanner(file), new(strings.Builder)
	for scanner.Scan() {
		fmt.Fprintln(stringBuilder, scanner.Text())
	}
	err = scanner.Err()
	panicError(err)

	fmt.Println(stringBuilder.String())
}

func demoReadFileUsingWriteTo() {
	fmt.Println(newline + fname(demoReadFileUsingWriteTo))
	file, err := os.OpenFile(FILENAME, os.O_RDONLY, 0644)
	panicError(err)
	defer file.Close()

	contentBuffer := new(strings.Builder)
	reader := bufio.NewReader(file)

	n, err := reader.WriteTo(contentBuffer)
	panicError(err)
	fmt.Println(n, "bytes have been read")
	fmt.Println(contentBuffer.String())
}
