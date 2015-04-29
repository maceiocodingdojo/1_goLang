package dojo

import "strconv"


// func SampleFunction(paramName paramType) returnType{
//

func IsDivisivelPor3(num int) string {
  if num % 3 == 0 {
    return "Fizz"
  } else {
    return strconv.Itoa(num)
  }
}

func IsDivisivelPor5(num int) string  {
  if num % 5 == 0 {
    return "Buzz"
  } else {
    return strconv.Itoa(num)
  }
}

func IsDivisivelPor3e5(num int) string {
  if num%3==0 && num%5==0 {
    return"fizzBuzz"
  } else {
    return strconv.Itoa(num)
  }
}
