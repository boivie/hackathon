package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"log"
	"encoding/csv"
	"io"
	"bytes"
	"encoding/gob"
	"crypto/sha1"
	"encoding/hex"
)

var (
	hashFile = flag.String("hash-file", "", "Generate hashed answers file")
)

type Mapping map[string]string

func hash(value string) string {
	h := sha1.New()
	io.WriteString(h, "Cheating is not allowed")
	io.WriteString(h, value)
	io.WriteString(h, "You will have to present how you solved it later anyway")
	return hex.EncodeToString(h.Sum(nil))
}

func generate(fname string) Mapping {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	mapping := make(Mapping)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		key := hash(strings.ToLower(strings.TrimSpace(record[0])))
		value := hash(strings.ToLower(strings.TrimSpace(record[1])))
		mapping[key] = value
	}
	return mapping
}

func main() {
	flag.Parse()

	if *hashFile != "" {
		fmt.Fprintln(os.Stderr, "Generating hash")
		var mapping Mapping
		mapping = generate(*hashFile)

		output, err := os.Create("answers.hashed")
		if err != nil {
			log.Fatal(err)
		}
		defer output.Close()

		enc := gob.NewEncoder(output)
		enc.Encode(mapping)

	} else {
		var answer Mapping

		data, err := Asset("answers.hashed")
		if err != nil {
			log.Fatal(err)
		}
		dec := gob.NewDecoder(bytes.NewReader(data))
		err = dec.Decode(&answer)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		r := csv.NewReader(file)

		true_positive := 0
		false_positive := 0

		asked := make(map[string]bool)
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			u := hash(strings.ToLower(strings.TrimSpace(record[0])))
			k := hash(strings.ToLower(strings.TrimSpace(record[1])))

			_, exists := asked[u + ":" + k]
			if !exists {
				v, exists := answer[u]
				if exists {
					if v == k {
						true_positive++
					} else {
						false_positive++
					}
				} else {
					false_positive++
				}
				asked[u + ":" + k] = true
			}
		}

		precision := float32(true_positive) / float32(true_positive + false_positive)
		recall := float32(true_positive) / float32(len(answer))

		var F1 float32
		if precision + recall > 0 {
			F1 = 2 * (precision * recall) / (precision + recall)
		} else {
			F1 = 0
		}

		fmt.Printf("Precision %.5f\n", precision)
		fmt.Printf("Recall: %.5f\n", recall)
		fmt.Printf("F1-score: %.5f\n", F1)
	}
}
