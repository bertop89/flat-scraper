package flat

import (
  "fmt"
)

type Flat struct {
	Name string
	Price int
	Rooms int
	Size int
	Store int
	Elevator bool
	Link string
}

func (f *Flat) ToString() string {
	return fmt.Sprintf("[ %d € ] | %d m² | %s https://www.idealista.com%s \n", f.Price, f.Size, f.Name, f.Link)
}