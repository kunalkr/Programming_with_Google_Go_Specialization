package main

import (
"fmt"
"math/rand"
"os"
)

func swap(sli []int,i int,j int){
    if sli[i] > sli[j]{
         temp := sli[j]
         sli[j] = sli[i]
         sli[i] = temp
    }
}

func bubble_sort(sli []int){
   for i:=0;i<len(sli);i++{
	for j:=i+1;j<len(sli);j++{
              swap(sli,i,j)
	}
   }

}

func main(){

sli := make([]int,0,10)
if os.Args[1] == "rand"{
for i:=0;i<10;i++{
    var num int
    num = rand.Intn(100)
    sli = append(sli,num)
}
}else{
for j:=0;j<10;j++{
    var num int
    fmt.Printf("Enter num%d:",j)
    _,err := fmt.Scanf("%d",&num)
    if err == nil{
    sli = append(sli,num)
    }
}
}

fmt.Println(sli)
bubble_sort(sli)
fmt.Println(sli)
}
