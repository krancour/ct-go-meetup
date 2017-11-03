package meetup

import (
	"fmt"
)

type ExperiencedGopher struct {
	Name            string
	Company         string
	YearsExperience int
}

func (e *ExperiencedGopher) GetGreeting() (string, error) {
	return fmt.Sprintf(
		"Hello! My name is %s and I work at %s. I've been using Go for %d years.",
		e.Name,
		e.Company,
		e.YearsExperience,
	), nil
}
