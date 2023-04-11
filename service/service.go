package service

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

func FinalCompute(req []uint8) {
	docs := []*Document{
		NewDoc(string(req)),
		NewDoc("Wondering where the quick brown fox had gone, the lazy dog howled at the moon. After the fox had jumped over the lazy dog, the dog let out a bark. The quick brown fox then ran away from the lazy dog and disappeared into the dense forest. The lazy dog was left alone, wondering where the fox had vanished to."),
		NewDoc("The lazy dog lay in the grass while the quick brown fox leaped over it. After the fox had successfully cleared the dog, the lazy canine let out a bark of surprise. The quick brown fox continued to run away from the dog and disappeared into the thick forest. Meanwhile, the lazy dog howled at the bright moon, puzzled as to where the fox had gone."),
	}

	for _, doc := range docs {
		for term := range doc.Terms {
			tfidf := ComputeTFIDF(doc, docs, term)
			if tfidf == 0.0 {
				continue
			} else {
				fmt.Printf("%s %f\n", term, tfidf)
			}
		}
	}

}
