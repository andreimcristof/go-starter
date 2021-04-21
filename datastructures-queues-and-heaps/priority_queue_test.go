package queue

import "testing"

func TestPriorityQueuePopWhenQueueEmptyShouldNotCrash(t *testing.T) {
	er := NewEmergencyRoom()

	next := er.HandleNextPatient()
	if next != nil {
		t.Error("expected empty queue to return nil")
	}
}

func TestPrioQueueWhenTwoPatientsReturnHighestPrio(t *testing.T) {
	er := NewEmergencyRoom()

	er.AdmitPatient(&Patient{name: "John", severity: Mild})
	er.AdmitPatient(&Patient{name: "Sam", severity: Critical})

	next := er.HandleNextPatient()

	if next.name != "Sam" {
		t.Error("expected Sam to be the next handled patient")
	}
}

func TestPrioQueueAfterUpdatePatientStatusReturnHighestPrio(t *testing.T) {
	er := NewEmergencyRoom()

	diana := &Patient{name: "Diana", severity: Mild}
	er.AdmitPatient(diana)

	liam := &Patient{name: "Liam", severity: Mild}
	er.AdmitPatient(liam)

	er.AdmitPatient(&Patient{name: "Jamal", severity: Mild})

	er.UpdatePatientStatus(liam, Critical)

	next := er.HandleNextPatient()
	if next.name != "Liam" {
		t.Error("Liam should be next returned")
	}

	er.UpdatePatientStatus(diana, Moderate)

	next2 := er.HandleNextPatient()

	if next2.name != "Diana" {
		t.Error("Diana should be next returned")
	}

	next3 := er.HandleNextPatient()
	if next3.name != "Jamal" {
		t.Error("Jamal should be next returned")
	}
}
