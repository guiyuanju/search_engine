package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func search(dir string, keyword string) {
	type FileName = string
	type Keyword = string
	type Item struct {
		count   int
		context []string
	}

	m := map[Keyword]map[FileName]Item{}
	type mChVal struct {
		kw       string
		filename string
		line     string
	}
	mCh := make(chan mChVal)
	go func() {
		for v := range mCh {
			if m[v.kw] == nil {
				m[v.kw] = map[FileName]Item{}
			}
			item := m[v.kw][v.filename]
			item.count++
			item.context = append(item.context, v.line)
			m[v.kw][v.filename] = item
		}
	}()

	var mWg sync.WaitGroup
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Fatal("error during walk: ", err)
		}

		if info.IsDir() {
			return nil
		}

		mWg.Add(1)
		go func() {
			defer mWg.Done()

			bs, err := os.ReadFile(path)
			if err != nil {
				log.Fatal("error when read file: ", err)
			}

			str := string(bs)
			words := []string{}
			for _, kw := range strings.Fields(str) {
				words = append(words, kw)
				if len(words) > 10 {
					words = words[1:]
				}
				mCh <- mChVal{kw, info.Name(), strings.Join(words, " ")}
			}
		}()

		return nil
	})
	mWg.Wait()

	type Pair struct {
		filename string
		item     Item
	}
	type FinamMChVal struct {
		kw    string
		pairs []Pair
	}
	finalM := map[Keyword][]Pair{}
	finalMCh := make(chan FinamMChVal)
	go func() {
		for v := range finalMCh {
			finalM[v.kw] = v.pairs
		}
	}()

	var wg sync.WaitGroup
	for kw, fileItemMap := range m {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := []Pair{}
			for filename, item := range fileItemMap {
				tmp = append(tmp, Pair{filename: filename, item: item})
			}
			slices.SortFunc(tmp, func(a, b Pair) int { return b.item.count - a.item.count })
			finalMCh <- FinamMChVal{kw, tmp}
		}()
	}

	wg.Wait()
	close(finalMCh)

	for _, pair := range finalM[keyword] {
		fmt.Printf("[%d]    %s\n", pair.item.count, pair.filename)
		for _, line := range pair.item.context {
			fmt.Printf("        %s\n", line)
		}
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Fatal("expect at least 3 argumets\n")
	}

	dir := os.Args[1]
	kw := os.Args[2]
	search(dir, kw)
}
