package intro

import (
	"testing"

	"fmt"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPrintInt(t *testing.T) {
	Convey("We can test printing an integer from an interface", t, func() {
		So(testPrintInt, ShouldNotPanic)
	})

	Convey("We can test printing a string from a stringprinter", t, func() {
		So(testPrintString, ShouldNotPanic)
	})

	Convey("We can print newString", t, func() {
		fmt.Println(newString)
		newString = "smallll joke"
		fmt.Println(newString)
	})

	Convey("Test iota!", t, func() {
		fmt.Println(code200)
		fmt.Println(code201)
	})
}

func TestCollections(t *testing.T) {
	Convey("We can test a slice of ints", t, func() {
		Iter(intSlc)
	})

	Convey("We can test a map", t, func() {
		MapIter(nnm)
	})

	Convey("We can test a channel", t, func() {
		ChanIter(sc)
	})
}
