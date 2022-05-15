# GEB Formal System Solver

This code was inspired by a challenge presented in Douglas Hofstadter's book, [Godel, Escher, Bach](https://en.wikipedia.org/wiki/G%C3%B6del,_Escher,_Bach). This challenge comes in the first chapter, and is called the MU Challenge.

The goal is, given a set of possible transformations, find the series of transformations that are need to get from a starting state, "MI", to an ending state, "MU". 

Written by [Alex Krentsel](www.Krentsel.com) in May 2022.

## Defining Transformations

Each transformation is encoded by a type adhering to the `Transform` interface. Defining a new Transform requires creating a new type and implementing the `apply(string) []string` method. 

Given an input string, this method should return a slice of the result of applying this transform a single time to any valid, matching place in the input. 

```go
type Transform1 string

func (t Transform1) apply(input string) []string {
	if len(input) > 0 && input[len(input)-1] == 'I' {
		return []string{input + "U"}
	}
	return []string{}
}
```

Note: be careful when programming arbitrary matching rules such as "III" -> "U", as you will need to correctly handle overlapping matches (following that rule, "IIII" -> "UI", "IU").

## Searching

Searching is implemented as a modified Breadth-First Search algorithm, with a "bag" mechanism that ensures we (1) keep track of how we got to any given state, and (2) do not apply rules in cycles. 

## Usage

Run the solver, and the solver will provide the sequence of rules it applied to get to a solution if it can find such a sequence, or otherwise let you know if it gives up after a configurable number of steps. 

```bash
go run solver.go
```

## Future Improvements

The largest area for improvement would be to avoid gumming up our search queue with rules that just expand forever but cannot possibly get to a solution. For example, the "Mx -> Mxx" rule will produces ever-increasing strings, which with the currently specified rules in the MIU game remain unsolvable in some cases ("MIU -> MIUIU -> MIUIUIU -> ...", no future application gets you to a state that can have any other transformation applied). 

Note that this is very hard though. There could be an arbitrary rule, such as `MIUIUIUIUIUIUIUIU -> MU`, which you couldn't apply until many applications of the `Mx->Mxx` rule. An optimization here requires more thinking (and may even be impossible). 

## License
[MIT](https://choosealicense.com/licenses/mit/)
