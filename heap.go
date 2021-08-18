package main

import (
	"fmt"
)

//Maxheap struct has a slice that holds the array
type Maxheap struct{
	array []int
}

///Insert adds an element to the heap
func (h *Maxheap) Insert(key int){
	h.array=append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

/// Extract reaturn the largest key, and removes it from heap
func (h *Maxheap) Extract() int{
	extracted :=h.array[0]
	l:=len(h.array)-1

	if len(h.array)==0{
		fmt.Println("cannot extract bcz length is zero")
		return -1
	}

	h.array[0]=h.array[len(h.array)-1]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return extracted
}

// Maxheapify down top to bottom
func (h *Maxheap) maxHeapifyDown(index int){
	lastIndex:=len(h.array)-1
	l,r :=leftIndex(index),rightIndex(index)
	childCompare :=0
	for l <= lastIndex{
		if l == lastIndex {
			childCompare=l 
		} else if h.array[l] > h.array[r]{
			childCompare=l
		} else {
			childCompare = r
		}
		if h.array[index] < h.array[childCompare]{
			h.swap(index,childCompare)
			index=childCompare
			l,r = leftIndex(index),rightIndex(index)
		} else {
			return 
		}
	}
}


// Maxheapify will heapify from bottom top
func (h *Maxheap) maxHeapifyUp(index int){
	for h.array[parentIndex(index)] < h.array[index]{
		h.swap(parentIndex(index),index)
		index = parentIndex(index)
	}
}

func parentIndex(i int) int {
	return (i-1)/2
}

func leftIndex(i int) int {
	return 2*i + 1
}

func rightIndex (i int) int {
	return 2*i+2
}

func (h *Maxheap) swap(i1,i2 int) {
	h.array[i1],h.array[i2]=h.array[i2],h.array[i1]
}

func main(){
	m:=&Maxheap{}
	buildHeap:=[]int{10,20,30,3,1,6,50,90,14,37}
	for _,v :=range buildHeap{
		m.Insert(v)
		fmt.Println(m)
	}
	for i:=0;i<5;i++{
		m.Extract()
		fmt.Println(m)
	}
}