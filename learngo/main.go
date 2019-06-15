package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/xiongyejun/xyjgo/tree"
)

func main() {
	fmt.Println(NonRepeatingSubStr("abc测试测试abc"))
}

func testAdder() {
	a := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Println(i, s)
	}
}

func testRetriverPoster() {
	var r RetriverPoster
	r = &Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf("%T %v\r\n", r, r)
	//	fmt.Println(downLoad(r))

	post(r)
}

const URL string = "http://www.imooc.com"

func post(poster IPost) {
	fmt.Println(poster.Post(URL, url.Values{
		"name":   []string{"ccmouse"},
		"course": []string{"golang"},
	}))
}

func downLoad(r IRetriever) string {
	return r.Get("http://www.imooc.com")
}

func doPrint(n *tree.Node) {
	fmt.Println(n.Value)
}
func testTree() {
	root := tree.NewNode(3)
	root.Left = &tree.Node{Value: 0}
	root.Rright = &tree.Node{Value: 5}
	root.Rright.Left = new(tree.Node)
	root.Left.Rright = tree.NewNode(2)

	root.Traversal(doPrint)
}

func NonRepeatingSubStr(s string) (string, int) {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0

	for i, ch := range []rune(s) {
		if lasti, ok := lastOccurred[ch]; ok && lasti >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			maxStart = start
		}
		lastOccurred[ch] = i
	}
	return string([]rune(s)[maxStart : maxStart+maxLength]), maxLength
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangele() {
	var a, b int = 3, 4
	var c int = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)

		if err != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}

			http.Error(writer, http.StatusText(code), code)
		}
	}
}
