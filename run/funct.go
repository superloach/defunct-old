package run

type Funct interface {
	String() string
	Call(*Under, Queue) (Funct, error)
}
