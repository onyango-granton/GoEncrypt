package testdrivenenvt

import "testing"

func TestGreet(t *testing.T){

	type test struct{
		input string
		want string
	}

	testArr := []test{
		{
			input: "Fred",
			want: "Welcome, Fred",
		},
		{
			input: "",
			want: "Welcome, Stranger",
		},
	}

	for _,singleTest := range testArr{
		got := greet(singleTest.input)

		if (got != singleTest.want){
			t.Errorf("greet(%v) got (%v) expected (%v)", singleTest.input, got, singleTest.want)
		}
	}

}