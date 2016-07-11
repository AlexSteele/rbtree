// Package rbtree provides an implementation of a red-black tree.
//
// Example:
//
//        package main  
//
//        import (
//                "fmt"
//
//                "github.com/alexsteele/rbtree"
//        )
//
//        func main() {
//                tree := rbTree.New(rbTree.IntComparator)
//                tree.Add(100)
//                tree.Add(50)
//                tree.Add(150)
//                first := tree.First()
//                last := tree.Last()
//                containsHundred := tree.Contains(100)
//                removedHundred := tree.Remove(100)
//                size := tree.Size()
//                isEmpty := tree.IsEmpty()
//                tree.ForEach(func(elem interface{}) { fmt.Println(elem); })
//
//                fmt.Println("First: ", first
//                fmt.Println("Last: ", last)
//                fmt.Println("Contains 100?: ", containsHundred)
//                fmt.Println("Removed 100?: ", removedHundred)
//                fmt.Println("Size: ", size)
//                fmt.Println("Empty?: ", isEmpty)
//        }
//        
package rbtree
