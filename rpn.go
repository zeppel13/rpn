/* This is ##NAME##(working title rpn), an easy commandline RPN-based calculator
 *
 * Copyright (C) Sebastian Kind 2015
 *
 *
 *
 * mail@sebastiankind.de
 * https://sebastiankind.de/
 *
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// ************************************************************************** //

//Calc contains the stack of an RPN-based calculator and the variables map

// ads
type Calc struct {
	stack     [4]float64
	variables map[string]float64 // ... nil map
}

const (
	e  = 2.71828182846
	pi = 3.14159265359
	//VERSION (obvious informatoin)
	VERSION = "0.1.1"
)

func (ca *Calc) enter(xValue float64) {
	ca.stack[3] = ca.stack[2]
	ca.stack[2] = ca.stack[1]
	ca.stack[1] = ca.stack[0]
	ca.stack[0] = xValue
}

//wtf OOP?
func (ca Calc) getStack() [4]float64 {
	a := ca.stack
	return (a)
}

func (ca *Calc) add() {
	ca.stack[0] += ca.stack[1]
	ca.refitStack()
}
func (ca *Calc) sub() {
	ca.stack[0] = ca.stack[1] - ca.stack[0]
	ca.refitStack()
}
func (ca *Calc) mul() {
	ca.stack[0] *= ca.stack[1]
	ca.refitStack()
}

func (ca *Calc) div() {
	ca.stack[0] = ca.stack[1] / ca.stack[0]
	ca.refitStack()
}
func (ca *Calc) pow() {
	ca.stack[0] = math.Pow(ca.stack[1], ca.stack[0])
	ca.refitStack()
}

func (ca *Calc) swapXY() {
	tmp := ca.stack[0]
	ca.stack[0] = ca.stack[1]
	ca.stack[1] = tmp
}

func (ca *Calc) xpow2() {
	ca.stack[0] = ca.stack[0] * ca.stack[0]
}

func (ca *Calc) sqrt() {
	ca.stack[0] = math.Sqrt(ca.stack[0])
}

func (ca *Calc) cbrt() {
	ca.stack[0] = math.Cbrt(ca.stack[0])
}

// xrt ...
func (ca *Calc) xrt() {
	ca.stack[0] = math.Pow(ca.stack[1], 1/ca.stack[0])
	ca.refitStack()
}

func (ca *Calc) modulo() {
	ca.stack[0] = (float64)((int64)(ca.stack[1]) % (int64)(ca.stack[0]))
	ca.refitStack()
}

func (ca *Calc) overX() {
	ca.stack[0] = 1.0 / ca.stack[0]

}

func (ca *Calc) sin() {
	ca.stack[0] = math.Sin(ca.stack[0] * (pi / 180))
}

func (ca *Calc) cos() {
	ca.stack[0] = math.Cos(ca.stack[0] * (pi / 180))
}

func (ca *Calc) tan() {
	ca.stack[0] = math.Tan(ca.stack[0] * (pi / 180))
}

func (ca *Calc) asin() {
	ca.stack[0] = math.Asin(ca.stack[0] * (pi / 180))
}

func (ca *Calc) acos() {
	ca.stack[0] = math.Acos(ca.stack[0] * (pi / 180))
}

func (ca *Calc) atan() {
	ca.stack[0] = math.Atan(ca.stack[0] * (pi / 180))
}

func (ca *Calc) exponent() {
	// ca.stack[1] Value
	// ca.stack[0] Exponent
	ca.stack[0] = ca.stack[1] * math.Pow(10, ca.stack[0])
	ca.refitStack()
}

// Basis = 10
func (ca *Calc) lg() {
	ca.stack[0] = math.Log10(ca.stack[0])
}

// Basis = e
func (ca *Calc) ln() {
	ca.stack[0] = math.Log(ca.stack[0])
}

// log_b n = x
// ca.stack[0] = b
// ca.stack[1] = n
func (ca *Calc) logbe() {
	ca.stack[0] = math.Log10(ca.stack[1]) / math.Log10(ca.stack[0])
	ca.refitStack()
}

// useless
func (ca *Calc) zTest() {
	// This functoin checks weather the value in ca.stack[0] is eaqual to zero
	//
	// It is not going to be used soon.
	if ca.stack[0] == 0 {
		ca.stack[0] = 0
	} else {
		ca.stack[0] = 1
	}
}

// Was habe ich hier gemacht?????
// FIX: useless
func (ca *Calc) deinMutter(haus int64, fisch float64) {
	fmt.Printf("Scientists found a ring around Uranos")
}

func (ca *Calc) pqSolver() {
	p := ca.stack[1]
	q := ca.stack[0]

	if (math.Pow(p/2, 2) - q) > 0 {
		ca.stack[0] = -p/2 + math.Sqrt(math.Pow(p/2, 2)-q)
		ca.stack[1] = -p/2 - math.Sqrt(math.Pow(p/2, 2)-q)
	} else {
		ca.stack[0] = -p / 2
		ca.stack[1] = -p / 2
	}
}

// Orthodromen-Strecke zwischen zwei Koordinaten ausrechnen.
func (ca *Calc) calcDistance() {
	// good old coord :)
	alat := ca.stack[3]
	along := ca.stack[2]
	blat := ca.stack[1]
	blong := ca.stack[0]
	var (
		distance float64
		angle    float64

		cosDeltaLambda float64
		deltaLong      float64

		aSin = math.Sin(alat * (pi / 180))
		bSin = math.Sin(blat * (pi / 180))
		aCos = math.Cos(alat * (pi / 180))
		bCos = math.Cos(blat * (pi / 180))
	)

	if along < 0.0 || blong < 0.0 {
		if along < 0.0 {
			along *= -1
		}
		if blong < 0.0 {
			blong *= -1
		}
		deltaLong = along + blong
	} else {
		deltaLong = along - blong
	}

	cosDeltaLambda = math.Cos(deltaLong * (pi / 180))
	angle = math.Acos(aSin*bSin + aCos*bCos*cosDeltaLambda)
	distance = 2 * pi * 6371 * ((angle * 180 / pi) / 360)

	ca.stack[0] = distance
}

func (ca *Calc) printBin() {
	binValue := (int64)(ca.stack[0])
	fmt.Printf("%b\n", binValue)
}

func (ca *Calc) rotateStack() {
	tmpStack := ca.getStack()
	ca.stack[0] = tmpStack[1]
	ca.stack[1] = tmpStack[2]
	ca.stack[2] = tmpStack[3]
	ca.stack[3] = tmpStack[0]
}

func (ca *Calc) drop() {
	ca.rotateStack()
	ca.stack[len(ca.stack)-1] = 0
}

func (ca *Calc) resetStack() {
	for i := range ca.getStack() {
		ca.stack[i] = 0
	}
}

func (ca *Calc) refitStack() {
	ca.stack[1] = ca.stack[2]
	ca.stack[2] = ca.stack[3]
	ca.stack[3] = 0
}

func (ca *Calc) setVariable(varName string) {
	ca.variables[varName] = ca.stack[0]
}

func (ca *Calc) getVariable(varName string) {
	ca.enter(ca.variables[varName])
}

func (ca Calc) getVariableMap() map[string]float64 {
	return ca.variables
}

// ************************************************************************** //
type rpnCommands struct {
	commands []string
	prioCmd  string
	inUse    bool
	index    int
}

func (r *rpnCommands) setCommands(script []string) {
	//script = strings.Trim(script, "\n")
	//r.commands = strings.Split(script, " ")
	r.commands = script
	r.index = -1
	/*for _, value := range r.commands {
		fmt.Println(value)
	}*/
}

// ?
func (r *rpnCommands) setPrio(task string) {
	r.prioCmd = task
}

func (r *rpnCommands) nextCommand() string {
	r.index++
	if r.index < len(r.commands) {
		return r.commands[r.index]
	}
	r.inUse = false
	return ""

}

func (r *rpnCommands) use() {
	r.inUse = true
}

func (r *rpnCommands) stop() {
	r.inUse = false
}

// ************************************************************************** //
func loadAndRunScript(filePath string, cmd *rpnCommands) {
	scriptFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer scriptFile.Close()

	var lines []string
	scanner := bufio.NewScanner(scriptFile)
	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
	cmd.setCommands(lines)
	cmd.use()
}

func saveMemory(fileName string, ca Calc) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for i := range ca.stack {
		file.WriteString(strconv.FormatFloat(ca.stack[3-i], 'f', 6, 64) + "\n")
	}
	for i, v := range ca.variables {
		file.WriteString(strconv.FormatFloat(v, 'f', 6, 64) + "\n->\n" + i + "\ndrop\n")
	}

}

func printStack(calc *Calc) {
	fmt.Println("-----------------------")
	fmt.Println()
	for i := range calc.getStack() {
		stackValue := (int64)(calc.getStack()[3-i])
		fmt.Printf("Stack %d: %v\t\t%x\n", 3-i, calc.getStack()[3-i], stackValue)
	}
}

func showVariables(ca Calc) {
	for i, v := range ca.getVariableMap() {
		fmt.Println(i, ":", v)
	}
}
func nextCommand(cmd *rpnCommands) string {
	if cmd.prioCmd != "" {
		next := cmd.prioCmd
		cmd.prioCmd = ""
		return next
	}
	if cmd.inUse == true {
		next := cmd.nextCommand()
		return next
	}
	input := ""
	fmt.Scanf("%v", &input)
	return input
}

// Useless? // Freitag, 20. Mai 2016 Yep
func helpM() {
	version()

}

func version() {
	fmt.Println("This is rpn Version", VERSION, "by Sebastian Kind (c) 2015 - 2016 #Yolo-Licence: use this Program at your own risk.")
}

// ***************** Mainloop
func inputLoop(calc *Calc, cmd rpnCommands) {
	stackView := 1
	for true {

		if stackView == 1 {
			printStack(calc)
		}

		input := nextCommand(&cmd)
		x, err := strconv.ParseFloat(input, 64)
		if err == nil {
			calc.enter(x)
		}

		switch input {
		case "q":
			return
		case "h":
			helpM()
		case "help":
			helpM()
		case "reset":
			calc.resetStack()
		case "drop":
			calc.drop()

		case "a":
			calc.add()
		case "+":
			calc.add()
		case "s":
			calc.sub()
		case "-":
			calc.sub()
		case "m":
			calc.mul()
		case "*":
			calc.mul()
		case "d":
			calc.div()
		case "/":
			calc.div()

		case "p":
			calc.pow()
		case "x2":
			calc.xpow2()
		case "sqrt":
			calc.sqrt()
		case "cbrt":
			calc.cbrt()
		case "xrt":
			calc.xrt()
		case "swap", "sw":
			calc.swapXY()
		case "r":
			calc.rotateStack()
		case "pi":
			calc.enter(pi)
		case "e":
			calc.enter(e)
		case "mod":
			calc.modulo()
		case "%":
			calc.modulo()
		case "ux":
			calc.overX()

		case "sin":
			calc.sin()
		case "cos":
			calc.cos()
		case "tan":
			calc.tan()
		case "asin":
			calc.asin()
		case "acos":
			calc.acos()
		case "atan":
			calc.atan()
		case "lg":
			calc.lg()
		case "ln":
			calc.ln()
		case "log":
			calc.logbe()
		case "pq":
			calc.pqSolver()
		case "ex":
			calc.exponent()
		case "lifetheuniverseandeverything":
			go fmt.Println("42")
		case "distance":
			calc.calcDistance()

		case "show":
			showVariables(*calc)
		case "->":
			inputName := nextCommand(&cmd)
			calc.setVariable(inputName)
		case "get":
			inputName := nextCommand(&cmd)
			calc.getVariable(inputName)

		case "store":
			//Should save Stack and Variables to a text file
			inputName := nextCommand(&cmd)
			saveMemory(inputName, *calc)
		case "run":
			// This will run a script or load stack variables from a text file
			inputPath := nextCommand(&cmd)
			loadAndRunScript(inputPath, &cmd)
		case "stop":
			// This should stop a script. I think it is unnecessary
			cmd.stop()

		case "view":
			stackView *= -1
		case "print":
			printStack(calc)
		case "bin":
			calc.printBin()
		}
	}
}

func main() {
	var (
		rpn      Calc
		commands rpnCommands
	)

	//Flagparsing
	pathPtr := flag.String("r", "", "./foo/bar Run a script. Type path like this")
	hidePtr := flag.Bool("hide", false, "Tell the programm to hide stack information, 'cuz u r like A B0ZZ \xF0\x9F\x98\x8E")
	versPtr := flag.Bool("version", false, "Displav Version")
	flag.Parse()
	//fmt.Println("run-flag:", *pathPtr)
	//fmt.Println("hide-flag:", *hidePtr)
	//fmt.Println("versoin-flag:", *versPtr)

	if *pathPtr != "" {
		loadAndRunScript(*pathPtr, &commands)
	}
	if *hidePtr == true {
		commands.setPrio("view")
	}
	if *versPtr == true {
		version()
	}

	rpn.variables = make(map[string]float64, 4)
	inputLoop(&rpn, commands)

}
