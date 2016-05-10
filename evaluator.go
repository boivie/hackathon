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
	"golang.org/x/crypto/pbkdf2"
	"crypto/sha1"
	"time"
	"encoding/gob"
	"encoding/hex"
)

var (
	hashFile = flag.String("hash-file", "", "Generate hashed answers file")
)

var salt = []byte("ff262653190b95a5a485993c5dc924753b59b5c8")

type Mapping struct {
	Iterations int
	Hashes     map[string]bool
}

func generate(fname string, iterations int) Mapping {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	var mapping Mapping
	mapping.Iterations = iterations
	mapping.Hashes = make(map[string]bool)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		var b bytes.Buffer
		b.Write([]byte("hackathon!:"))
		for _, word := range record {
			b.Write([]byte(strings.TrimSpace(word)))
			b.Write([]byte(":"))
		}

		dk := pbkdf2.Key(b.Bytes(), salt, iterations, 32, sha1.New)
		mapping.Hashes[hex.EncodeToString(dk)] = true
	}
	return mapping
}

func main() {
	flag.Parse()

	if *hashFile != "" {
		fmt.Fprintln(os.Stderr, "Generating hash")
		i := 1000
		var mapping Mapping
		for try := 0; try < 5; try++ {
			start := time.Now()
			mapping = generate(*hashFile, i)

			d := float32(time.Since(start) / time.Millisecond)
			fmt.Fprintf(os.Stderr, "%d iterations took %d ms\n", i, int(d))

			i = int(float32(i) * 1000.0 / d)
		}

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

		var test Mapping = generate(flag.Arg(0), answer.Iterations)

		correct := 0
		for k, _ := range test.Hashes {
			_, exists := answer.Hashes[k]
			if exists {
				correct++;
			}
		}

		fmt.Printf("%d%% correct\n", int(100.0 * float32(correct) / float32(len(answer.Hashes))))
	}
}
