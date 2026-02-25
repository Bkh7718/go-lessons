package main

import (
	"lesson-4/pkg/billing"
	"log"
)

func main() {
	
	result := billing.Calculate10000(9999.0)
	log.Print(result, " diram")

}

