// Copyright 2019 Tsuyoshi Ozawa
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
