package scaffold

type Scaffold interface {
	Generate() error
}

type ScaffoldBuilder interface {
	Build() Scaffold
}
