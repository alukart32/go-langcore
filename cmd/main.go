// GO practice
// basic types: https://go101.org/article/basic-types-and-value-literals.html

package main

import (
	controlflow "alukart32.com/langcore/internal/controlFlow"
	"alukart32.com/langcore/internal/function"
	"alukart32.com/langcore/internal/scope"
	"alukart32.com/langcore/internal/types"
	"alukart32.com/langcore/tools"
)

var topicConsoleLogger = tools.ConsoleLogger{Format: "\n%s\n%s", Spliter: "==================================\n"}
var subTopicConsoleLogger = tools.ConsoleLogger{Format: "%s %s\n", Spliter: "::\n"}

func main() {
	topicConsoleLogger.Log("Basic Types")
	subTopicConsoleLogger.Log("Get uintSizes and int size")
	types.IntSizeBits()
	subTopicConsoleLogger.Log("Float point values")
	types.FpValues()
	subTopicConsoleLogger.Log("Float point hex values")
	types.FpHexValues()
	subTopicConsoleLogger.Log("Rune values")
	types.RuneValues()
	subTopicConsoleLogger.Log("String values")
	types.StringValues()
	// const
	subTopicConsoleLogger.Log("Const values")
	types.ConstValues()
	types.UntypedNamedConst()
	// var scopes
	subTopicConsoleLogger.Log("Scopes values")
	scope.Scopes()

	topicConsoleLogger.Log("Control flow")
	// if-else
	subTopicConsoleLogger.Log("IF-ELSE")
	controlflow.IfElse()
	controlflow.IfElseAndRand()
	controlflow.IfElseDayTimeRange()
	// for
	subTopicConsoleLogger.Log("FOR")
	controlflow.ForBaseForm()
	controlflow.ForCountDown()
	controlflow.ForBreak()
	controlflow.ForContinue()
	// switch
	subTopicConsoleLogger.Log("SWITCH")
	controlflow.SwitchBaseForm()
	controlflow.SwitchModN()
	controlflow.SwitchFallthrough()
	// goto
	subTopicConsoleLogger.Log("GOTO")
	controlflow.GoToInc()
	controlflow.GoTofindSmallestPrimeLargerThan()

	topicConsoleLogger.Log("FUNCTIONS")
	// anonymous func
	subTopicConsoleLogger.Log("Anonymous functions")
	function.AnonymousFunc()
	subTopicConsoleLogger.Log("Variadic functions")
	println("Variadic function Sum call: ", function.Sum(1, 5, 87))
	subTopicConsoleLogger.Log("Fucntion Value")
	function.FuncValue()

	// container types
	topicConsoleLogger.Log("CONTAINER TYPES: array, slice, map")
	types.ContainerTypeDeclaration()
	types.ArraySliceCompositeliterals()
	types.GetPointerFromContainerType()
	types.SimplifyCompositeLiterals()
	types.RetrieveAndModifyContainerTypes()
	types.ContainerAssignments()
	types.AppendAndDeleteFromMap()
	types.SliceAppend()
	types.MakeSliceAndMap()

	// slices
	//types.Reslicing()
	subTopicConsoleLogger.Log("SLICE")
	types.CopySlices()
	types.RangeOverArrayModificationHasNoEffects()
	types.RangeOverSliceModificationHasEffects()
	types.RangeOverMap()
	types.UseArrayPointer()
	types.OverArrayWithPanic()
	types.IndexArrayElementWithPointer()
	types.DeriveSlicesFromArrayPointer()

	subTopicConsoleLogger.Log("SLICE: delete elems by condition")
	isNumeralNegativeOrZero := func(n int) bool {
		return !(n <= 0)
	}
	types.DeleteElements([]int{1, -1, 2, -2, 3, 0}, isNumeralNegativeOrZero, true)

	subTopicConsoleLogger.Log("SLICE: stack operations")
	types.StackOper()

	// String
	topicConsoleLogger.Log("STRING type")
	types.Strings()
	subTopicConsoleLogger.Log("STRING operations")
	types.StringOper()
	types.StringLen()
	subTopicConsoleLogger.Log("STRING: conversion between ByteSlice & RuneSlice")
	types.StringConversionByteSliceToRune()

	// Struct
	topicConsoleLogger.Log("STRUCT type")
	subTopicConsoleLogger.Log("Book struct type example")
	types.BookStructExample()

	//Interfaces
	topicConsoleLogger.Log("INTERFACE type")
	subTopicConsoleLogger.Log("INTERFACE: value boxing")
	types.InterfacesBoxing()
	subTopicConsoleLogger.Log("INTERFACE: value boxing into blank interface")
	types.BoxingValuesWithBlankInterface()

	// Interfaces Polymorphism
	topicConsoleLogger.Log("INTERFACE: Polymorphism")
	types.Polymorphism()
}
