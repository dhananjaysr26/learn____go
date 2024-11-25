package main

import (
	"fmt"
	"sync"
	"time"
)

/*

With Out Go Routine

*/

// func Task(i int) {
//     defer fmt.Printf("END:Task No. %d\n", i)
//     time.Sleep(1 * time.Second)
// }

// func main() {
//     fmt.Println("!!  Task Runner !!")
//     count := 10
//     for i := 1; i <= count; i++ {
//         fmt.Printf("START: Task No %d\n", i)
//         Task(i)
//     }
// }

/*

With Go Routine

*/

func Task(i int, wg *sync.WaitGroup) {  
    defer wg.Done() 
    defer fmt.Printf("END: Task No. %d\n", i)  
    time.Sleep(1 * time.Second)    
}

func main() {
    fmt.Println("!!  Task Runner !!")
    count := 1000000

    // Create WaitGroup to track goroutines
    var wg sync.WaitGroup              
    
    for i := 1; i <= count; i++ {  
        wg.Add(1)                   
        fmt.Printf("START: Task No %d\n", i)  
        go Task(i, &wg)               
    }
    
    wg.Wait()                       
    fmt.Println("All tasks completed!")
}