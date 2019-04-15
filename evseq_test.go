package evseq

import (
	"testing"
)

func TestPushEventFormFront(t *testing.T) {
	es := NewEventSequence()
	es.PushEventFromFront(ClusterAddNode, 1.0, nil)
	es.PushEventFromFront(ClusterDelNode, 3.0, nil)
	es.PushEventFromFront(ClusterAddPod, 2.0, nil)
	if es.GetNextEvent().evtype != ClusterAddNode {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent().evtype != ClusterAddPod {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent().evtype != ClusterDelNode {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent() != nil {
		t.Fatal("an unexpected event is in evseq.")
	}

}

func TestPushEventFormBack(t *testing.T) {
	es := NewEventSequence()
	es.PushEventFromBack(ClusterAddNode, 1.0, nil)
	es.PushEventFromBack(ClusterDelNode, 3.0, nil)
	es.PushEventFromBack(ClusterAddPod, 2.0, nil)
	if es.GetNextEvent().evtype != ClusterAddNode {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent().evtype != ClusterAddPod {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent().evtype != ClusterDelNode {
		t.Fatal("the event order is unexpected.")
	}

	if es.GetNextEvent() != nil {
		t.Fatal("an unexpected event is in evseq.")
	}

}
