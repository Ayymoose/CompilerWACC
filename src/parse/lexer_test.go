package parse

import (
	"io/ioutil"
	"testing"

	. "grammar"
)

func TestBasicExitExit1(t *testing.T) {
	// [BEGIN : "begin\r" LINE_FEED : "\n" EXIT : "exit " NEG : "-" NUMBER : "1" CARRIAGE_RETURN : "\r" LINE_FEED : "\n" END : "end\r" LINE_FEED : "\n"]
	want := []Item{Item{BEGIN, "begin\r"}, Item{LINE_FEED, "\n"},
		Item{EXIT, "exit "}, Item{NEG, "-"}, Item{NUMBER, "1"},
		Item{CARRIAGE_RETURN, "\r"}, Item{LINE_FEED, "\n"},
		Item{END, "end\r"}, Item{LINE_FEED, "\n"}}

	var got []Item

	b, err := ioutil.ReadFile("../../wacc_examples/valid/basic/exit/exit-1.wacc")
	if err != nil {
		t.Error(err)
	}
	s := string(b)
	lex := Lex("exit-1.wacc", s)
	for item := range lex.Items {
		got = append(got, item)
	}
	for i := 0; i < len(got) || i < len(want); i++ {
		if got[i] != want[i] {
			t.Errorf("expected: %v\ngot:%v", want, got)
		}
	}
}
