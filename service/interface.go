package service

import "cyolo-exercise/model"

type Processor interface {
	Process(data string) []error
	Histogram(top int) []model.Pair
}
