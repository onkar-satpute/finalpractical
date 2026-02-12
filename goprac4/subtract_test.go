package main
import "testing"

func TestSubtract(t *testing.T){
	tests := []struct{a,b,res int}{{10,5,5},{20,10,10}}
	for _,tt := range tests{
		if Subtract(tt.a,tt.b)!=tt.res{
			t.Error("Failed")
		}
	}
}
