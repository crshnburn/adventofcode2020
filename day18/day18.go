package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"
)

func Calculate(input string) int {
	input = strings.ReplaceAll(input, "*", "-")
	expr, _ := parser.ParseExpr(input)
	return evalExpression(expr)
}

func Calculate2(input string) int {
	input = strings.ReplaceAll(input, "*", "-")
	input = strings.ReplaceAll(input, "+", "*")
	expr, _ := parser.ParseExpr(input)
	return evalExpression(expr)
}

func evalExpression(expr ast.Expr) int {
	switch node := expr.(type) {
	case *ast.BasicLit:
		val, _ := strconv.Atoi(node.Value)
		// fmt.Println(val)
		return val
	case *ast.ParenExpr:
		return evalExpression(node.X)
	case *ast.BinaryExpr:
		switch node.Op {
		case token.ADD, token.MUL:
			x := evalExpression(node.X)
			y := evalExpression(node.Y)
			// fmt.Println(x, "+", y)
			return x + y
		case token.SUB:
			x := evalExpression(node.X)
			y := evalExpression(node.Y)
			// fmt.Println(x, "*", y)
			return x * y
		}
	}
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += Calculate(line)
		sum2 += Calculate2(line)
	}
	fmt.Println("Part 1:", sum)
	fmt.Println("Part 2:", sum2)
}
