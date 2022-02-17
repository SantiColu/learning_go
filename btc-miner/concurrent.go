package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

const zeros = 8
const workers = 1000
const content = "We dance like Marionettes Swaying to the symphony ....."

var c []byte = []byte(content)
var zstr string = strings.Repeat("0", zeros)
var start time.Time = time.Now()

func main() {
	wg := sync.WaitGroup{}
	limit := int(math.Pow(10, 12))
	each := limit / workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		from := each + each*(i-1)
		to := each + each*i - 1
		// fmt.Printf("Starting worker #%v from: %v to %v...\n", i, from, to)
		go checkRange(from, to, &wg)
	}

	wg.Wait()
}

func checkRange(from, to int, wg *sync.WaitGroup) {
	wg.Add(1)
	for nonce := from; nonce <= to; nonce++ {
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

	}

	wg.Done()
}
