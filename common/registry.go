package common

type Solver func(string) string

var solvers = map[int]Solver{}

func RegisterSolver(day int, solver Solver) {
	solvers[day] = solver
}

func GetSolver(day int) (Solver, bool) {
	solver, exists := solvers[day]
	return solver, exists
}
