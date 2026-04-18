package domain

type CodeGenerator interface { // esto es un Output Port (Driven Port)
	Generate() (string, error)
}
