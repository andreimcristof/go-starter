package queue

type PatientsQueue []*Patient

// Push required by the heap.Interface
func (pq *PatientsQueue) Push(patientData interface{}) {
	patient := patientData.(*Patient)
	patient.index = len(*pq) // required by the heap.Interface
	*pq = append(*pq, patient)
}

// Pop required by the heap.Interface
func (pq *PatientsQueue) Pop() interface{} {
	currentQueue := *pq
	n := len(currentQueue)
	patient := currentQueue[n-1]
	patient.index = -1 // required by the heap
	*pq = currentQueue[0 : n-1]
	return patient
}

// Len required by sort.Interface
func (pq *PatientsQueue) Len() int {
	return len(*pq)
}

// Swap required by the sort.Interface
func (pq PatientsQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].index = a
	pq[b].index = b
}

// Less required by the sort.Interface
// we flip the comparer (to greater than) because we need
// the comparer to sort by highest prio, not lowest
func (pq PatientsQueue) Less(a, b int) bool {
	return pq[a].severity > pq[b].severity
}
