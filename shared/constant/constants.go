// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/droidrun/mobilerun-sdk-go/v6/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type CancelEvent string               // Always "CancelEvent"
type CreatedEvent string              // Always "CreatedEvent"
type ExceptionEvent string            // Always "ExceptionEvent"
type ExecutorActionEvent string       // Always "ExecutorActionEvent"
type ExecutorActionResultEvent string // Always "ExecutorActionResultEvent"
type ExecutorContextEvent string      // Always "ExecutorContextEvent"
type ExecutorInputEvent string        // Always "ExecutorInputEvent"
type ExecutorResponseEvent string     // Always "ExecutorResponseEvent"
type ExecutorResultEvent string       // Always "ExecutorResultEvent"
type FastAgentEndEvent string         // Always "FastAgentEndEvent"
type FastAgentExecuteEvent string     // Always "FastAgentExecuteEvent"
type FastAgentInputEvent string       // Always "FastAgentInputEvent"
type FastAgentOutputEvent string      // Always "FastAgentOutputEvent"
type FastAgentResponseEvent string    // Always "FastAgentResponseEvent"
type FastAgentResultEvent string      // Always "FastAgentResultEvent"
type FastAgentToolCallEvent string    // Always "FastAgentToolCallEvent"
type FinalizeEvent string             // Always "FinalizeEvent"
type List string                      // Always "list"
type ManagerContextEvent string       // Always "ManagerContextEvent"
type ManagerInputEvent string         // Always "ManagerInputEvent"
type ManagerPlanDetailsEvent string   // Always "ManagerPlanDetailsEvent"
type ManagerPlanEvent string          // Always "ManagerPlanEvent"
type ManagerResponseEvent string      // Always "ManagerResponseEvent"
type Model string                     // Always "model"
type QueuedEvent string               // Always "QueuedEvent"
type RecordUiStateEvent string        // Always "RecordUIStateEvent"
type ResultEvent string               // Always "ResultEvent"
type ScreenshotEvent string           // Always "ScreenshotEvent"
type Socks5 string                    // Always "socks5"
type StartEvent string                // Always "StartEvent"
type StopEvent string                 // Always "StopEvent"
type ToolExecutionEvent string        // Always "ToolExecutionEvent"
type UserMessageEvent string          // Always "UserMessageEvent"
type Wireguard string                 // Always "wireguard"

func (c CancelEvent) Default() CancelEvent                 { return "CancelEvent" }
func (c CreatedEvent) Default() CreatedEvent               { return "CreatedEvent" }
func (c ExceptionEvent) Default() ExceptionEvent           { return "ExceptionEvent" }
func (c ExecutorActionEvent) Default() ExecutorActionEvent { return "ExecutorActionEvent" }
func (c ExecutorActionResultEvent) Default() ExecutorActionResultEvent {
	return "ExecutorActionResultEvent"
}
func (c ExecutorContextEvent) Default() ExecutorContextEvent       { return "ExecutorContextEvent" }
func (c ExecutorInputEvent) Default() ExecutorInputEvent           { return "ExecutorInputEvent" }
func (c ExecutorResponseEvent) Default() ExecutorResponseEvent     { return "ExecutorResponseEvent" }
func (c ExecutorResultEvent) Default() ExecutorResultEvent         { return "ExecutorResultEvent" }
func (c FastAgentEndEvent) Default() FastAgentEndEvent             { return "FastAgentEndEvent" }
func (c FastAgentExecuteEvent) Default() FastAgentExecuteEvent     { return "FastAgentExecuteEvent" }
func (c FastAgentInputEvent) Default() FastAgentInputEvent         { return "FastAgentInputEvent" }
func (c FastAgentOutputEvent) Default() FastAgentOutputEvent       { return "FastAgentOutputEvent" }
func (c FastAgentResponseEvent) Default() FastAgentResponseEvent   { return "FastAgentResponseEvent" }
func (c FastAgentResultEvent) Default() FastAgentResultEvent       { return "FastAgentResultEvent" }
func (c FastAgentToolCallEvent) Default() FastAgentToolCallEvent   { return "FastAgentToolCallEvent" }
func (c FinalizeEvent) Default() FinalizeEvent                     { return "FinalizeEvent" }
func (c List) Default() List                                       { return "list" }
func (c ManagerContextEvent) Default() ManagerContextEvent         { return "ManagerContextEvent" }
func (c ManagerInputEvent) Default() ManagerInputEvent             { return "ManagerInputEvent" }
func (c ManagerPlanDetailsEvent) Default() ManagerPlanDetailsEvent { return "ManagerPlanDetailsEvent" }
func (c ManagerPlanEvent) Default() ManagerPlanEvent               { return "ManagerPlanEvent" }
func (c ManagerResponseEvent) Default() ManagerResponseEvent       { return "ManagerResponseEvent" }
func (c Model) Default() Model                                     { return "model" }
func (c QueuedEvent) Default() QueuedEvent                         { return "QueuedEvent" }
func (c RecordUiStateEvent) Default() RecordUiStateEvent           { return "RecordUIStateEvent" }
func (c ResultEvent) Default() ResultEvent                         { return "ResultEvent" }
func (c ScreenshotEvent) Default() ScreenshotEvent                 { return "ScreenshotEvent" }
func (c Socks5) Default() Socks5                                   { return "socks5" }
func (c StartEvent) Default() StartEvent                           { return "StartEvent" }
func (c StopEvent) Default() StopEvent                             { return "StopEvent" }
func (c ToolExecutionEvent) Default() ToolExecutionEvent           { return "ToolExecutionEvent" }
func (c UserMessageEvent) Default() UserMessageEvent               { return "UserMessageEvent" }
func (c Wireguard) Default() Wireguard                             { return "wireguard" }

func (c CancelEvent) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c CreatedEvent) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ExceptionEvent) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ExecutorActionEvent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ExecutorActionResultEvent) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c ExecutorContextEvent) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ExecutorInputEvent) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ExecutorResponseEvent) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ExecutorResultEvent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c FastAgentEndEvent) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c FastAgentExecuteEvent) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c FastAgentInputEvent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c FastAgentOutputEvent) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c FastAgentResponseEvent) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c FastAgentResultEvent) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c FastAgentToolCallEvent) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c FinalizeEvent) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ManagerContextEvent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ManagerInputEvent) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ManagerPlanDetailsEvent) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ManagerPlanEvent) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c ManagerResponseEvent) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c QueuedEvent) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c RecordUiStateEvent) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c ResultEvent) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c ScreenshotEvent) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Socks5) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c StartEvent) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c StopEvent) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ToolExecutionEvent) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c UserMessageEvent) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Wireguard) MarshalJSON() ([]byte, error)                 { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
