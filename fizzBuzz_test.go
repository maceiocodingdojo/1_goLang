package dojo

import "testing"

// If test fail
// t.Fatalf("at index %d, expected %d, got %d.", i, v, val)


// func TestSampleFunction(t *testing.T){
// if test fail
// t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
//}


func TestSeNumeroDivisivelPor3(t *testing.T) {
  if IsDivisivelPor3(3) {
    t.Fatalf("divisao invalida")
  }
  if IsDivisivelPor3(6) {
    t.Fatalf("divisao invalida")
  }
  if IsDivisivelPor3(9) {
    t.Fatalf("divisao invalida")
  }
}

/*func TestSeNumeroDivisivelPor5(t *testing.T) {
  if
}*/
