package dojo

import "testing"

// If test fail
// t.Fatalf("at index %d, expected %d, got %d.", i, v, val)


// func TestSampleFunction(t *testing.T){
// if test fail
// t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
//}


func TestSeNumeroDivisivelPor3(t *testing.T) {
  if IsDivisivelPor3(15) {
    t.Fatalf("divisao invalida")
  }
}

func TestSenumeroDivisivelPor5(t *testing.T){
  if IsDivisivelPor5(15) {
    t.Fatalf("divisao invalida")
  }
}
