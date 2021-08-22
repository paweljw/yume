package game

import (
	par "github.com/prataprc/goparsec"
	"fmt"
)

func makeCommandY(ast *par.AST) par.Parser {
	var command par.Parser

	// Commons
	rest := par.Token(`[^<]+`, "REST")
	maybeRest := ast.Maybe("maybe_rest", nil, rest)
	target := par.Token(`[^<]+`, "TARGET")

	// AT modifier
	atToken := par.Token("(at|on|into)", "TARGET_MODIFIER")
	maybeAt := ast.Maybe("maybe_at", nil, atToken)

	// quit command
	quitVerb := par.Atom("quit", "VERB")
	quitCommand := ast.And("quit_command", nil, quitVerb, maybeRest)

	// look command
	lookVerb := par.Atom("look", "VERB")
	lookCommand := ast.And("look_command", nil, lookVerb, maybeAt, target)

	// set_flag
	setFlagVerb := par.Atom("set_flag", "VERB")
	player := par.Token("[a-zA-Z]+", "PLAYER")
	flagName := par.Token("[a-z-A-Z_]+", "FLAG")
	setFlagCommand := ast.And("set_flag_command", nil, setFlagVerb, player, flagName)

	// unset_flag
	unsetFlagVerb := par.Atom("unset_flag", "VERB")
	unsetFlagCommand := ast.And("unset_flag_command", nil, unsetFlagVerb, player, flagName)

	// DIRECTIONS
	direction := par.Atom("(w|west|e|east|n|north|s|south|ne|northeast|se|southeast|sw|southwest|nw|northwest|u|up|d|down)", "DIRECTION")
	directionCommand := ast.And("direction_command", nil, direction)

	// go command
	// TODO: imply this verb if no other found in implementation
	goVerb := par.Atom("go", "VERB")
	goCommand := ast.And("go_command", nil, goVerb, direction)

	// failure verb
	failureVerb := par.Token(`[^<]+`, "VERB")
	failureCommand := ast.And("failure_command", nil, failureVerb, maybeRest)

	// All commands together
	command = ast.OrdChoice(
		"cmd",
		nil,
		quitCommand,
		lookCommand,
		setFlagCommand,
		unsetFlagCommand,
		directionCommand,
		goCommand,
		failureCommand,
	)

	return command
}

func GetFirst(ast *par.AST, kind string) (string, bool) {
	ch := make(chan par.Queryable, 100)
	go ast.Query(kind, ch)
	node, ok := <-ch

	if ok {
		return string(node.GetValue()), true
	} else {
		return "", false
	}
}

func TestLexer() {
	data := []byte(`west`)

	ast := par.NewAST("cmd", 100)
	y := makeCommandY(ast)
	s := par.NewScanner(data).TrackLineno()
	ast.Parsewith(y, s)

	tok, _ := GetFirst(ast, "VERB")
	fmt.Println(tok)
	tok, _ = GetFirst(ast, "TARGET_MODIFIER")
	fmt.Println(tok)
	tok, _ = GetFirst(ast, "TARGET")
	fmt.Println(tok)

	differentData := []byte(`go west`)
	ast = par.NewAST("cmds", 100)
	y = makeCommandY(ast)
	q := par.NewScanner(differentData)
	ast.Parsewith(y, q)

	fmt.Println("---")

	tok, _ = GetFirst(ast, "VERB")
	fmt.Println(tok)
	tok, _ = GetFirst(ast, "DIRECTION")
	fmt.Println(tok)
	tok, _ = GetFirst(ast, "TARGET_MODIFIER")
	fmt.Println(tok)
	tok, _ = GetFirst(ast, "TARGET")
	fmt.Println(tok)

}
