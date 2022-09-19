// GO practice
// basic types: https://go101.org/article/basic-types-and-value-literals.html

package main

import (
	controlflow "alukart32.com/langcore/internal/controlFlow"
	"alukart32.com/langcore/internal/function"
	"alukart32.com/langcore/internal/scope"
	"alukart32.com/langcore/internal/types"
	"alukart32.com/tools"
)

func main() {
	tools.TopicConsoleLogger.Log("Basic Types")
	tools.SubTopicConsoleLogger.Log("Get uintSizes and int size")
	types.IntSizeBits()
	tools.SubTopicConsoleLogger.Log("Float point values")
	types.FpValues()
	tools.SubTopicConsoleLogger.Log("Float point hex values")
	types.FpHexValues()
	tools.SubTopicConsoleLogger.Log("Rune values")
	types.RuneValues()
	tools.SubTopicConsoleLogger.Log("String values")
	types.StringValues()
	// const
	tools.SubTopicConsoleLogger.Log("Const values")
	types.ConstValues()
	types.UntypedNamedConst()
	// var scopes
	tools.SubTopicConsoleLogger.Log("Scopes values")
	scope.Scopes()

	tools.TopicConsoleLogger.Log("Control flow")
	// if-else
	tools.SubTopicConsoleLogger.Log("IF-ELSE")
	controlflow.IfElse()
	controlflow.IfElseAndRand()
	controlflow.IfElseDayTimeRange()
	// for
	tools.SubTopicConsoleLogger.Log("FOR")
	controlflow.ForBaseForm()
	controlflow.ForCountDown()
	controlflow.ForBreak()
	controlflow.ForContinue()
	// switch
	tools.SubTopicConsoleLogger.Log("SWITCH")
	controlflow.SwitchBaseForm()
	controlflow.SwitchModN()
	controlflow.SwitchFallthrough()
	// goto
	tools.SubTopicConsoleLogger.Log("GOTO")
	controlflow.GoToInc()
	controlflow.GoTofindSmallestPrimeLargerThan()

	tools.TopicConsoleLogger.Log("FUNCTIONS")
	// anonymous func
	tools.SubTopicConsoleLogger.Log("Anonymous functions")
	function.AnonymousFunc()
	tools.SubTopicConsoleLogger.Log("Variadic functions")
	println("Variadic function Sum call: ", function.Sum(1, 5, 87))
	tools.SubTopicConsoleLogger.Log("Fucntion Value")
	function.FuncValue()

	// container types
	tools.TopicConsoleLogger.Log("CONTAINER TYPES: array, slice, map")
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
	tools.SubTopicConsoleLogger.Log("SLICE")
	types.CopySlices()
	types.RangeOverArrayModificationHasNoEffects()
	types.RangeOverSliceModificationHasEffects()
	types.RangeOverMap()
	types.UseArrayPointer()
	types.OverArrayWithPanic()
	types.IndexArrayElementWithPointer()
	types.DeriveSlicesFromArrayPointer()

	tools.SubTopicConsoleLogger.Log("SLICE: delete elems by condition")
	isNumeralNegativeOrZero := func(n int) bool {
		return !(n <= 0)
	}
	types.DeleteElements([]int{1, -1, 2, -2, 3, 0}, isNumeralNegativeOrZero, true)

	tools.SubTopicConsoleLogger.Log("SLICE: stack operations")
	types.StackOper()

	// String
	tools.TopicConsoleLogger.Log("STRING type")
	types.Strings()
	tools.SubTopicConsoleLogger.Log("STRING operations")
	types.StringOper()
	types.StringLen()
	tools.SubTopicConsoleLogger.Log("STRING: conversion between ByteSlice & RuneSlice")
	types.StringConversionByteSliceToRune()

	// Struct
	tools.TopicConsoleLogger.Log("STRUCT type")
	tools.SubTopicConsoleLogger.Log("Book struct type example")
	types.BookStructExample()

	//Interfaces
	tools.TopicConsoleLogger.Log("INTERFACE type")
	tools.SubTopicConsoleLogger.Log("INTERFACE: value boxing")
	types.InterfacesBoxing()
	tools.SubTopicConsoleLogger.Log("INTERFACE: value boxing into blank interface")
	types.BoxingValuesWithBlankInterface()

	// Interfaces Polymorphism
	tools.TopicConsoleLogger.Log("INTERFACE: Polymorphism")
	types.Polymorphism()

	tools.SubTopicConsoleLogger.Log("INTERFACE: type assertion")
	types.InterfaceTypeAssertion1()
	types.InterfaceTypeAssertion2()
	types.InterfceTypeSwitch()
}
