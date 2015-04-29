package dojo

//import "strconv"


// func SampleFunction(paramName paramType) returnType{
//
func IsDivisivelPor3(num int) int {
  if num % 3 == 0 {
    return "Fizz"


  }
}

func IsDivisivelPor5(num int) int  {
  if num % 5 == 0 {
    return "Buzz"
  }
}
func IsDivisivelPor3(num int) int {
  if (num%3==0) && (num%5==0){
    return"fizz-buzz"
  }
}
