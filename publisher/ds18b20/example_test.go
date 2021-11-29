package ds18b20

import (
	"fmt"
	"log"
)

func main() {
	d, err := New()
	if err != nil {
		log.Fatalf("failed new: %v", err)
	}

	e, err := d.Read()
	if err != nil {
		log.Fatalf("fataled read: %v", err)
	}
	fmt.Println(e.Temperature)
	// Output:
	// 28.750°C
}
