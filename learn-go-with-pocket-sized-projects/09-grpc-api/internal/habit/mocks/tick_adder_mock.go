// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mocks

//go:generate minimock -i alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit.tickAdder -o tick_adder_mock.go -n TickAdderMock -p mocks

import (
	mm_habit "alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit"
	"context"
	"sync"
	mm_atomic "sync/atomic"
	"time"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// TickAdderMock implements mm_habit.tickAdder
type TickAdderMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAddTick          func(ctx context.Context, id mm_habit.ID, t time.Time) (err error)
	funcAddTickOrigin    string
	inspectFuncAddTick   func(ctx context.Context, id mm_habit.ID, t time.Time)
	afterAddTickCounter  uint64
	beforeAddTickCounter uint64
	AddTickMock          mTickAdderMockAddTick
}

// NewTickAdderMock returns a mock for mm_habit.tickAdder
func NewTickAdderMock(t minimock.Tester) *TickAdderMock {
	m := &TickAdderMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddTickMock = mTickAdderMockAddTick{mock: m}
	m.AddTickMock.callArgs = []*TickAdderMockAddTickParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mTickAdderMockAddTick struct {
	optional           bool
	mock               *TickAdderMock
	defaultExpectation *TickAdderMockAddTickExpectation
	expectations       []*TickAdderMockAddTickExpectation

	callArgs []*TickAdderMockAddTickParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// TickAdderMockAddTickExpectation specifies expectation struct of the tickAdder.AddTick
type TickAdderMockAddTickExpectation struct {
	mock               *TickAdderMock
	params             *TickAdderMockAddTickParams
	paramPtrs          *TickAdderMockAddTickParamPtrs
	expectationOrigins TickAdderMockAddTickExpectationOrigins
	results            *TickAdderMockAddTickResults
	returnOrigin       string
	Counter            uint64
}

// TickAdderMockAddTickParams contains parameters of the tickAdder.AddTick
type TickAdderMockAddTickParams struct {
	ctx context.Context
	id  mm_habit.ID
	t   time.Time
}

// TickAdderMockAddTickParamPtrs contains pointers to parameters of the tickAdder.AddTick
type TickAdderMockAddTickParamPtrs struct {
	ctx *context.Context
	id  *mm_habit.ID
	t   *time.Time
}

// TickAdderMockAddTickResults contains results of the tickAdder.AddTick
type TickAdderMockAddTickResults struct {
	err error
}

// TickAdderMockAddTickOrigins contains origins of expectations of the tickAdder.AddTick
type TickAdderMockAddTickExpectationOrigins struct {
	origin    string
	originCtx string
	originId  string
	originT   string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAddTick *mTickAdderMockAddTick) Optional() *mTickAdderMockAddTick {
	mmAddTick.optional = true
	return mmAddTick
}

// Expect sets up expected params for tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Expect(ctx context.Context, id mm_habit.ID, t time.Time) *mTickAdderMockAddTick {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{}
	}

	if mmAddTick.defaultExpectation.paramPtrs != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by ExpectParams functions")
	}

	mmAddTick.defaultExpectation.params = &TickAdderMockAddTickParams{ctx, id, t}
	mmAddTick.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmAddTick.expectations {
		if minimock.Equal(e.params, mmAddTick.defaultExpectation.params) {
			mmAddTick.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddTick.defaultExpectation.params)
		}
	}

	return mmAddTick
}

// ExpectCtxParam1 sets up expected param ctx for tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) ExpectCtxParam1(ctx context.Context) *mTickAdderMockAddTick {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{}
	}

	if mmAddTick.defaultExpectation.params != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Expect")
	}

	if mmAddTick.defaultExpectation.paramPtrs == nil {
		mmAddTick.defaultExpectation.paramPtrs = &TickAdderMockAddTickParamPtrs{}
	}
	mmAddTick.defaultExpectation.paramPtrs.ctx = &ctx
	mmAddTick.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmAddTick
}

// ExpectIdParam2 sets up expected param id for tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) ExpectIdParam2(id mm_habit.ID) *mTickAdderMockAddTick {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{}
	}

	if mmAddTick.defaultExpectation.params != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Expect")
	}

	if mmAddTick.defaultExpectation.paramPtrs == nil {
		mmAddTick.defaultExpectation.paramPtrs = &TickAdderMockAddTickParamPtrs{}
	}
	mmAddTick.defaultExpectation.paramPtrs.id = &id
	mmAddTick.defaultExpectation.expectationOrigins.originId = minimock.CallerInfo(1)

	return mmAddTick
}

// ExpectTParam3 sets up expected param t for tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) ExpectTParam3(t time.Time) *mTickAdderMockAddTick {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{}
	}

	if mmAddTick.defaultExpectation.params != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Expect")
	}

	if mmAddTick.defaultExpectation.paramPtrs == nil {
		mmAddTick.defaultExpectation.paramPtrs = &TickAdderMockAddTickParamPtrs{}
	}
	mmAddTick.defaultExpectation.paramPtrs.t = &t
	mmAddTick.defaultExpectation.expectationOrigins.originT = minimock.CallerInfo(1)

	return mmAddTick
}

// Inspect accepts an inspector function that has same arguments as the tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Inspect(f func(ctx context.Context, id mm_habit.ID, t time.Time)) *mTickAdderMockAddTick {
	if mmAddTick.mock.inspectFuncAddTick != nil {
		mmAddTick.mock.t.Fatalf("Inspect function is already set for TickAdderMock.AddTick")
	}

	mmAddTick.mock.inspectFuncAddTick = f

	return mmAddTick
}

// Return sets up results that will be returned by tickAdder.AddTick
func (mmAddTick *mTickAdderMockAddTick) Return(err error) *TickAdderMock {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	if mmAddTick.defaultExpectation == nil {
		mmAddTick.defaultExpectation = &TickAdderMockAddTickExpectation{mock: mmAddTick.mock}
	}
	mmAddTick.defaultExpectation.results = &TickAdderMockAddTickResults{err}
	mmAddTick.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmAddTick.mock
}

// Set uses given function f to mock the tickAdder.AddTick method
func (mmAddTick *mTickAdderMockAddTick) Set(f func(ctx context.Context, id mm_habit.ID, t time.Time) (err error)) *TickAdderMock {
	if mmAddTick.defaultExpectation != nil {
		mmAddTick.mock.t.Fatalf("Default expectation is already set for the tickAdder.AddTick method")
	}

	if len(mmAddTick.expectations) > 0 {
		mmAddTick.mock.t.Fatalf("Some expectations are already set for the tickAdder.AddTick method")
	}

	mmAddTick.mock.funcAddTick = f
	mmAddTick.mock.funcAddTickOrigin = minimock.CallerInfo(1)
	return mmAddTick.mock
}

// When sets expectation for the tickAdder.AddTick which will trigger the result defined by the following
// Then helper
func (mmAddTick *mTickAdderMockAddTick) When(ctx context.Context, id mm_habit.ID, t time.Time) *TickAdderMockAddTickExpectation {
	if mmAddTick.mock.funcAddTick != nil {
		mmAddTick.mock.t.Fatalf("TickAdderMock.AddTick mock is already set by Set")
	}

	expectation := &TickAdderMockAddTickExpectation{
		mock:               mmAddTick.mock,
		params:             &TickAdderMockAddTickParams{ctx, id, t},
		expectationOrigins: TickAdderMockAddTickExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmAddTick.expectations = append(mmAddTick.expectations, expectation)
	return expectation
}

// Then sets up tickAdder.AddTick return parameters for the expectation previously defined by the When method
func (e *TickAdderMockAddTickExpectation) Then(err error) *TickAdderMock {
	e.results = &TickAdderMockAddTickResults{err}
	return e.mock
}

// Times sets number of times tickAdder.AddTick should be invoked
func (mmAddTick *mTickAdderMockAddTick) Times(n uint64) *mTickAdderMockAddTick {
	if n == 0 {
		mmAddTick.mock.t.Fatalf("Times of TickAdderMock.AddTick mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAddTick.expectedInvocations, n)
	mmAddTick.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmAddTick
}

func (mmAddTick *mTickAdderMockAddTick) invocationsDone() bool {
	if len(mmAddTick.expectations) == 0 && mmAddTick.defaultExpectation == nil && mmAddTick.mock.funcAddTick == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAddTick.mock.afterAddTickCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAddTick.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// AddTick implements mm_habit.tickAdder
func (mmAddTick *TickAdderMock) AddTick(ctx context.Context, id mm_habit.ID, t time.Time) (err error) {
	mm_atomic.AddUint64(&mmAddTick.beforeAddTickCounter, 1)
	defer mm_atomic.AddUint64(&mmAddTick.afterAddTickCounter, 1)

	mmAddTick.t.Helper()

	if mmAddTick.inspectFuncAddTick != nil {
		mmAddTick.inspectFuncAddTick(ctx, id, t)
	}

	mm_params := TickAdderMockAddTickParams{ctx, id, t}

	// Record call args
	mmAddTick.AddTickMock.mutex.Lock()
	mmAddTick.AddTickMock.callArgs = append(mmAddTick.AddTickMock.callArgs, &mm_params)
	mmAddTick.AddTickMock.mutex.Unlock()

	for _, e := range mmAddTick.AddTickMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddTick.AddTickMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddTick.AddTickMock.defaultExpectation.Counter, 1)
		mm_want := mmAddTick.AddTickMock.defaultExpectation.params
		mm_want_ptrs := mmAddTick.AddTickMock.defaultExpectation.paramPtrs

		mm_got := TickAdderMockAddTickParams{ctx, id, t}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmAddTick.t.Errorf("TickAdderMock.AddTick got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddTick.AddTickMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.id != nil && !minimock.Equal(*mm_want_ptrs.id, mm_got.id) {
				mmAddTick.t.Errorf("TickAdderMock.AddTick got unexpected parameter id, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddTick.AddTickMock.defaultExpectation.expectationOrigins.originId, *mm_want_ptrs.id, mm_got.id, minimock.Diff(*mm_want_ptrs.id, mm_got.id))
			}

			if mm_want_ptrs.t != nil && !minimock.Equal(*mm_want_ptrs.t, mm_got.t) {
				mmAddTick.t.Errorf("TickAdderMock.AddTick got unexpected parameter t, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAddTick.AddTickMock.defaultExpectation.expectationOrigins.originT, *mm_want_ptrs.t, mm_got.t, minimock.Diff(*mm_want_ptrs.t, mm_got.t))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddTick.t.Errorf("TickAdderMock.AddTick got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmAddTick.AddTickMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddTick.AddTickMock.defaultExpectation.results
		if mm_results == nil {
			mmAddTick.t.Fatal("No results are set for the TickAdderMock.AddTick")
		}
		return (*mm_results).err
	}
	if mmAddTick.funcAddTick != nil {
		return mmAddTick.funcAddTick(ctx, id, t)
	}
	mmAddTick.t.Fatalf("Unexpected call to TickAdderMock.AddTick. %v %v %v", ctx, id, t)
	return
}

// AddTickAfterCounter returns a count of finished TickAdderMock.AddTick invocations
func (mmAddTick *TickAdderMock) AddTickAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddTick.afterAddTickCounter)
}

// AddTickBeforeCounter returns a count of TickAdderMock.AddTick invocations
func (mmAddTick *TickAdderMock) AddTickBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddTick.beforeAddTickCounter)
}

// Calls returns a list of arguments used in each call to TickAdderMock.AddTick.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddTick *mTickAdderMockAddTick) Calls() []*TickAdderMockAddTickParams {
	mmAddTick.mutex.RLock()

	argCopy := make([]*TickAdderMockAddTickParams, len(mmAddTick.callArgs))
	copy(argCopy, mmAddTick.callArgs)

	mmAddTick.mutex.RUnlock()

	return argCopy
}

// MinimockAddTickDone returns true if the count of the AddTick invocations corresponds
// the number of defined expectations
func (m *TickAdderMock) MinimockAddTickDone() bool {
	if m.AddTickMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddTickMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddTickMock.invocationsDone()
}

// MinimockAddTickInspect logs each unmet expectation
func (m *TickAdderMock) MinimockAddTickInspect() {
	for _, e := range m.AddTickMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TickAdderMock.AddTick at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterAddTickCounter := mm_atomic.LoadUint64(&m.afterAddTickCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddTickMock.defaultExpectation != nil && afterAddTickCounter < 1 {
		if m.AddTickMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to TickAdderMock.AddTick at\n%s", m.AddTickMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to TickAdderMock.AddTick at\n%s with params: %#v", m.AddTickMock.defaultExpectation.expectationOrigins.origin, *m.AddTickMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddTick != nil && afterAddTickCounter < 1 {
		m.t.Errorf("Expected call to TickAdderMock.AddTick at\n%s", m.funcAddTickOrigin)
	}

	if !m.AddTickMock.invocationsDone() && afterAddTickCounter > 0 {
		m.t.Errorf("Expected %d calls to TickAdderMock.AddTick at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.AddTickMock.expectedInvocations), m.AddTickMock.expectedInvocationsOrigin, afterAddTickCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TickAdderMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddTickInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TickAdderMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *TickAdderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddTickDone()
}
