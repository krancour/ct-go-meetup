package meetup

import (
	"fmt"
)

type NoviceGopher struct {
	Name    string
	Company string
}

func (n *NoviceGopher) GetGreeting() (string, error) {
	return fmt.Sprintf(
		"Hello! My name is %s and I work at %s. I'm new to Go!",
		n.Name,
		n.Company,
	), nil
}
