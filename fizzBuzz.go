package dojo

//import "strconv"


// func SampleFunction(paramName paramType) returnType{
//
num:=[]int{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15};

func IsDivisivelPor3(num int) string {
  if num % 3 == 0 {
    return "Fizz"


  }
}

func IsDivisivelPor5(num int) string  {
  if num % 5 == 0 {
    return "Buzz"
  }
}
func IsDivisivelPor3(num int) string {
  if (num%3==0) && (num%5==0){
    return"fizz-buzz"
  }
}
