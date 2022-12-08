package structures

import (
	"errors"
)

type PQueueInterface interface {
	Enqueue(value int, priority int)
	Dequeue(priority int) int
}

type PQueue struct {
	PiorityLists map[int]LinkedList
}

type LinkedList struct {
	Value int
	Next  *LinkedList
	Last  *LinkedList
}

func (pq *PQueue) Enqueue(value int, priority int) {
	if len(pq.PiorityLists) == 0 {
		pq.PiorityLists = map[int]LinkedList{
			priority: LinkedList{
				Value: value,
			},
		}
		return
	}

	if linkedList, ok := pq.PiorityLists[priority]; ok {
		newLast := LinkedList{
			Value: value,
		}
		if linkedList.Last == nil {
			linkedList.Next = &newLast
			linkedList.Last = &newLast
		} else {
			linkedList.Last = &newLast
			pq.PiorityLists[priority].Last.Next = &newLast
		}

		pq.PiorityLists[priority] = linkedList
	} else {
		pq.PiorityLists[priority] = LinkedList{
			Value: value,
		}
	}
	return
}

func (pq *PQueue) Dequeue(priority int) (int, error) {
	if _, ok := pq.PiorityLists[priority]; !ok {
		return 0, errors.New("Elements by priority not found")
	}

	result := pq.PiorityLists[priority].Value
	if pq.PiorityLists[priority].Next != nil {
		pq.PiorityLists[priority] = *pq.PiorityLists[priority].Next
	} else {
		delete(pq.PiorityLists, priority)
	}
	return result, nil
}
