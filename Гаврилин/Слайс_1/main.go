package main

import (
    "fmt"
    "math/rand"
    "time"
)

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    arr := make([]int, 10)
    for i := range arr {
        arr[i] = rand.Intn(100)
    }
    fmt.Println("До", arr)
    bubbleSort(arr)
    fmt.Println("После", arr)
}