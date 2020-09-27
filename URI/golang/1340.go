package main

import (
  "io"
  "fmt"
  "container/list"
  "container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
  value int
  index int
}

// PriorityQueue type
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { 
  return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
  return pq[i].value > pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}

// Push add a element into a PQ
func (pq *PriorityQueue) Push(element interface{}) {
  size := len(*pq)
  item := element.(*Item)
  item.index = size
  
  *pq = append(*pq, item)
}

// Pop remove the less element of a PQ
func (pq *PriorityQueue) Pop() interface{} {
  old := *pq
  n := len(old)
  item := old[n - 1]
  
  old[n - 1] = nil  // avoid memory leak
  item.index = -1 // for safety
  
  *pq = old[0:n - 1]

  return item
}

// Q1340 Incomplete
// https://www.urionlinejudge.com.br/judge/pt/problems/view/1340
func Q1340() {
  var n int
  
  for {
    _, err := fmt.Scanf("%d", &n)

    if err == io.EOF {
      break
    }
    
    stack := list.New()
    queue := list.New()
    var priorityQueue PriorityQueue

    var command, num int
    isQueue, isStack, isPQ := true, true, true

    for ; n > 0; n-- {
      fmt.Scanf("%d %d", &command, &num)

      if !isQueue && !isPQ && !isStack {
        continue
      }

      if (command == 1) {
        stack.PushBack(num)
        queue.PushBack(num)
        heap.Push(&priorityQueue, &Item{ value: num })
      } else {
        stackElement := stack.Back()
        isStack = isStack && stackElement != nil && stackElement.Value == num
        stack.Remove(stackElement)

        queueElement := queue.Front()
        isQueue = isQueue && queueElement != nil && queueElement.Value == num
        queue.Remove(queueElement)

        pqElement := heap.Pop(&priorityQueue).(*Item)
        isPQ = isPQ && pqElement != nil && pqElement.value == num
      }
    }

    if !isPQ && !isStack && !isQueue {
      fmt.Printf("impossible\n")
    } else if (isQueue && isStack) || (isQueue && isPQ) || (isStack && isPQ) {
      fmt.Printf("not sure\n")
    } else if isQueue {
      fmt.Printf("queue\n")
    } else if isPQ {
      fmt.Printf("priority queue\n")
    } else {
      fmt.Printf("stack\n")
    }
  }
}