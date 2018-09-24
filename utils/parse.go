package utils

import (
	"strconv"
	"fmt"
)

 func ParseInt(value string) (int64){
	result,_ := strconv.ParseInt(value, 10,64)
	return result	
 }

 func ParseFloat(value string) (float64){
	result,_ := strconv.ParseFloat(value, 64)
	return result	
 }

 func ParseFloat2String(value float64) string{
	result := fmt.Sprintf("%v", value)
	return result	
 }