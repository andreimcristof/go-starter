package queue

import "container/heap"

type EmergencyRoom struct {
	patientsQueue PatientsQueue
}

func NewEmergencyRoom() *EmergencyRoom {
	er := &EmergencyRoom{patientsQueue: make(PatientsQueue, 0)}
	heap.Init(&er.patientsQueue)
	return er
}

// AdmitPatient adds new patients in ER's priority queue
func (er *EmergencyRoom) AdmitPatient(patient interface{}) {
	potentialPatient := patient.(*Patient)
	heap.Push(&er.patientsQueue, potentialPatient)
}

// HandleNextPatient remove highest prio from the patients queue
func (er *EmergencyRoom) HandleNextPatient() *Patient {
	if er.patientsQueue.Len() == 0 {
		return nil
	}
	nextPatient := heap.Pop(&er.patientsQueue)
	return nextPatient.(*Patient)
}

func (er *EmergencyRoom) UpdatePatientStatus(patient *Patient, newStatus SeverityStatus) {
	patient.severity = newStatus
	heap.Fix(&er.patientsQueue, patient.index)
}
