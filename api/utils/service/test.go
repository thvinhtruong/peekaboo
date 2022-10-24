package service

import (
	"log"
	"math/rand"
	"time"
)

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
