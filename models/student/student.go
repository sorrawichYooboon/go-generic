package studentmodels

import "reflect"

type IStudentScore interface {
	string | int | float64
}

type Student[T IStudentScore] struct {
	Name  string
	Score T
	Age   int
}

func (s *Student[T]) GetScore() T {
	return s.Score
}

func (s *Student[T]) SetScore(score T) {
	s.Score = score
}

func (s *Student[T]) GetScoreType() string {
	return reflect.TypeOf(s.Score).String()
}
