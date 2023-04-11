package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"strings"
)

type Document struct {
	Content string
	Terms   map[string]int
}

func NewDoc(content string) *Document {
	re, err := regexp.Compile(`[^\w]`) //clean the string remove punctuations. There should be a better way to do this tbh
	if err != nil {
		log.Fatal(err)
	}
	content = re.ReplaceAllString(content, " ")
	doc := &Document{ //doc is defined
		Content: content,
		Terms:   make(map[string]int),
	}
	terms := strings.Fields(doc.Content)
	for _, term := range terms {
		doc.Terms[term]++
	}
	return doc
}

func ComputeIDF(docs []*Document, term string) float64 {
	n := float64(len(docs))
	df := 0.0
	for _, doc := range docs {
		if doc.Terms[term] > 0 {
			df++
		}
	}
	idf := math.Log(n / df)
	return idf
}

func ComputeTFIDF(doc *Document, docs []*Document, term string) float64 {
	tf := float64(doc.Terms[term]) / float64(len(doc.Terms))
	idf := ComputeIDF(docs, term)
	tfidf := tf * idf
	return tfidf
}

func main() {
	docs := []*Document{
		NewDoc("This is a sample document."),
		NewDoc("Another sample document."),
		NewDoc("And yet another sample document."),
	}
	for _, doc := range docs {
		for term := range doc.Terms {
			tfidf := ComputeTFIDF(doc, docs, term)
			fmt.Printf("%s %f\n", term, tfidf)
		}
	}

}
