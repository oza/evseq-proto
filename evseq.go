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
	"container/list"
	"fmt"
)

//type EventContext interface {
//	GetValue() interface{}
//}

type Event struct {
	time float64
	evtype string
	ctx interface{}
}

const (
	ClusterAddNode = "ClusterAddNode"
	ClusterDelNode = "ClusterDelNode"
	ClusterDetectNodeFailure = "ClusterDetectNodeFailure"
	ClusterAddPod = "ClusterAddPod"
	ClusterDelPod = "ClusterDelPod"
	SchedulerStartJob = "SchedulerStartJob"
	SchedulerExitJob = "SchedulerExitJob"
	SchedulerPreemptJob = "SchedulerPreemptJob"
	SubmitterSubmitJob = "SubmitterSubmitJob"
	SubmitterDeleteJob = "SubmitterDeleteJob"
)


type KubeSim struct {
	curtime float64
}

type EventSequence struct {
	events *list.List
}

func (e *Event) setTime(time float64) {
	e.time = time
}

func NewEvent(evtype string, t float64, ctx interface{}) *Event  {
	return &Event {
		evtype: evtype,
		time: t,
		ctx: ctx,
	}
}

func NewEventSequence() *EventSequence {
	return &EventSequence {
		events: list.New(),
	}
}

func (es *EventSequence) GetNextEvent() *Event {
	if es.events.Len() == 0 {
		return nil
	}
	e := es.events.Front()
	es.events.Remove(e)
	return e.Value.(*Event)
}

func (es *EventSequence) PeekNextEvent() *Event {
	e := es.events.Front()
	return e.Value.(*Event)
}

func (es *EventSequence) PushEventFromFront(evtype string, t float64, ctx interface{}) {
	if es.events.Len() == 0 {
		es.events.PushBack(NewEvent(evtype, t, ctx))
		return
	}

	for e := es.events.Front(); e != nil; e = e.Next() {
		ev := e.Value.(*Event)
		// fmt.Println("cur %v t %v", e.Value.(*Event), t)
		if ev.time > t {
			es.events.InsertBefore(NewEvent(evtype, t, ctx), e)
			return
		}
	}
	es.events.PushBack(NewEvent(evtype, t, ctx))
}

func (es *EventSequence) PushEventFromBack(evtype string, t float64, ctx interface{}) {
	if es.events.Len() == 0 {
		es.events.PushFront(NewEvent(evtype, t, ctx))
		return
	}

	for e := es.events.Back(); e != nil; e = e.Prev() {
		// fmt.Println(e.Value.(*Event))
		ev := e.Value.(*Event)
		if ev.time < t {
			es.events.InsertAfter(NewEvent(evtype, t, ctx), e)
			return
		}
	}

	es.events.PushFront(NewEvent(evtype, t, ctx))
}

func (es *EventSequence) PrintEvents() {
	for e := es.events.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(*Event))
	}
}

func main() {



	/** 
	// Simulator's main loop.
	es := NewEventSequence()
	// Kick a first event.
	es.PushEventFromFront(ClusterAddNode, 1.0, nil)
	es.PushEventFromFront(SubmitterSubmitJob, initialSubmittionDelay, nil)
	for ev := es.GetNextEvent() ; ev != nil;  {
		curtime := ev.time
		switch ev.evtype {
		case ClusterAddNode:
			// time := ev.ctx.(*float64)
			//HandleClusterAddNode(time)
		case ClusterDelNode:
			//ctx := ev.ctx.GetValue().(*CustomType)
			//HandleClusterAddNode(time, ctx)
		case SubmitterSubmitJob:
			//ji := ev.ctx.GetValue().(*JobInfo) // :which implements GetValue)
			es.PushEventFromFront(SubmitterSubmitJob, curtime + submitionInterval, ji)
		case SchedulerStartJob:
			//cv := ev.ctx.GetValue().(*JobInfo)

		default:
			abort()
		}
	}
	**/
		
	//es.PrintEvents()
	//for e := l.Front(); e != nil; e = e.Next() {
	//	fmt.Println(e.Value)
	//}

}
