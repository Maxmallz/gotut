package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strings"
)

func main() {
	src, err := os.OpenFile("src.json", os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	dst, err := os.OpenFile("dst.json", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		panic(err)
	}

}

func printMostOccuringLinesFromFilesInADir() {
	dir := "/Users/mirogbele/dev/flux/flux-svc-core/Foundation/Processing"
	countByLine := make(map[string]int)
	maxSizes := []int{}

	files, err := ioutil.ReadDir(dir)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		f, err := os.Open(dir + "/" + file.Name())

		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			countByLine[strings.TrimSpace(scanner.Text())]++
		}
	}

	for _, v := range countByLine {
		maxSizes = append(maxSizes, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(maxSizes)))

	maxSizes = maxSizes[:10]

	for i := range maxSizes {
		for k, v := range countByLine {
			if maxSizes[i] == v {
				fmt.Printf("[%q | %v]\n", k, v)
			}
		}
	}
}

func printSimpleHTTPCall() {
	url := "https://flux-core-dev.tesla.com/api/v2/admin/tenant"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	fmt.Printf("%s\n", b)
}
