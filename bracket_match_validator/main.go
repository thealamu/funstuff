/* Given a bracket sequence like '[{()}]',
this program validates each opening bracket
is matched with its closing bracket.
*/

package main

import (
	"bufio"
	"container/list"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		print("Enter a bracket sequence: ")
		bracketSeq, err := getBracketSequence(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}

		isValid := isValidSequence(bracketSeq)
		if isValid {
			fmt.Println("✔ Sequence is valid")
		} else {
			fmt.Println("⚠ Not a valid sequence")
		}
	}
}

func isValidSequence(bSeq string) bool {
	closers := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	//maintain a stack
	stack := list.New()

	for _, chr := range bSeq {
		opener, isCloser := closers[string(chr)]
		if !isCloser {
			//new opener, push to stack
			stack.PushFront(string(chr))
			continue
		}
		//a closer, top of the stack must match opener
		stackHead := stack.Front()
		if stackHead == nil {
			return false
		}
		stackHeadValue, ok := stackHead.Value.(string)
		if !ok {
			log.Fatalf("invalid type %T in stack", stackHead.Value)
		}

		if opener != stackHeadValue {
			return false
		}
		//pop the stack
		stack.Remove(stackHead)
	}

	if stack.Len() != 0 {
		return false
	}

	return true
}

func getBracketSequence(r io.Reader) (string, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	s.Scan()

	line := strings.TrimSpace(s.Text())

	if line == "" {
		return "", errors.New("got empty sequence")
	}
	if err := s.Err(); err != nil {
		return "", err
	}

	return line, nil
}
