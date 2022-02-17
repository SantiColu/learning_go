package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const zeros = 7
const content = "We dance like Marionettes Swaying to the symphony ....."

func main() {
	limit := int(math.Pow(10, 12))
	c := []byte(content)
	zstr := strings.Repeat("0", zeros)
	start := time.Now()
	lastupd := start
	for nonce := 0; nonce <= limit; nonce++ {
		hasher := sha256.New()
		hasher.Write(append(c, []byte(strconv.Itoa(nonce))...))
		hash := hex.EncodeToString(hasher.Sum(nil))

		if hash[:zeros] == zstr {
			fmt.Printf("\nFound nonce %v in %v, hash: %v\n", nonce, time.Since(start), hash)
			break
		}

		d := time.Since(start)
		if d > time.Minute*10 {
			fmt.Printf("\nTime limit!! max nonce: %v\n", nonce)
			break
		}

		if time.Since(lastupd) > time.Second {
			lastupd = time.Now()
			fmt.Printf("\r%vs | %v%% | %v | %v", int(d.Seconds()), int(float64(nonce)/float64(limit)*100), hash[:zeros], nonce)
		}
	}
}
