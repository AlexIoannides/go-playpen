// Code generated by http://github.com/gojuno/minimock (v3.4.0). DO NOT EDIT.

package mocks

//go:generate minimock -i alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit.habitCreator -o habit_creator_mock.go -n HabitCreatorMock -p mocks

import (
	mm_habit "alexioannides/go-playpen/learn-go-with-pocket-sized-projects/09-grpc-api/internal/habit"
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// HabitCreatorMock implements mm_habit.habitCreator
type HabitCreatorMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcAdd          func(ctx context.Context, habit mm_habit.Habit) (err error)
	funcAddOrigin    string
	inspectFuncAdd   func(ctx context.Context, habit mm_habit.Habit)
	afterAddCounter  uint64
	beforeAddCounter uint64
	AddMock          mHabitCreatorMockAdd
}

// NewHabitCreatorMock returns a mock for mm_habit.habitCreator
func NewHabitCreatorMock(t minimock.Tester) *HabitCreatorMock {
	m := &HabitCreatorMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddMock = mHabitCreatorMockAdd{mock: m}
	m.AddMock.callArgs = []*HabitCreatorMockAddParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mHabitCreatorMockAdd struct {
	optional           bool
	mock               *HabitCreatorMock
	defaultExpectation *HabitCreatorMockAddExpectation
	expectations       []*HabitCreatorMockAddExpectation

	callArgs []*HabitCreatorMockAddParams
	mutex    sync.RWMutex

	expectedInvocations       uint64
	expectedInvocationsOrigin string
}

// HabitCreatorMockAddExpectation specifies expectation struct of the habitCreator.Add
type HabitCreatorMockAddExpectation struct {
	mock               *HabitCreatorMock
	params             *HabitCreatorMockAddParams
	paramPtrs          *HabitCreatorMockAddParamPtrs
	expectationOrigins HabitCreatorMockAddExpectationOrigins
	results            *HabitCreatorMockAddResults
	returnOrigin       string
	Counter            uint64
}

// HabitCreatorMockAddParams contains parameters of the habitCreator.Add
type HabitCreatorMockAddParams struct {
	ctx   context.Context
	habit mm_habit.Habit
}

// HabitCreatorMockAddParamPtrs contains pointers to parameters of the habitCreator.Add
type HabitCreatorMockAddParamPtrs struct {
	ctx   *context.Context
	habit *mm_habit.Habit
}

// HabitCreatorMockAddResults contains results of the habitCreator.Add
type HabitCreatorMockAddResults struct {
	err error
}

// HabitCreatorMockAddOrigins contains origins of expectations of the habitCreator.Add
type HabitCreatorMockAddExpectationOrigins struct {
	origin      string
	originCtx   string
	originHabit string
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmAdd *mHabitCreatorMockAdd) Optional() *mHabitCreatorMockAdd {
	mmAdd.optional = true
	return mmAdd
}

// Expect sets up expected params for habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Expect(ctx context.Context, habit mm_habit.Habit) *mHabitCreatorMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{}
	}

	if mmAdd.defaultExpectation.paramPtrs != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by ExpectParams functions")
	}

	mmAdd.defaultExpectation.params = &HabitCreatorMockAddParams{ctx, habit}
	mmAdd.defaultExpectation.expectationOrigins.origin = minimock.CallerInfo(1)
	for _, e := range mmAdd.expectations {
		if minimock.Equal(e.params, mmAdd.defaultExpectation.params) {
			mmAdd.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAdd.defaultExpectation.params)
		}
	}

	return mmAdd
}

// ExpectCtxParam1 sets up expected param ctx for habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) ExpectCtxParam1(ctx context.Context) *mHabitCreatorMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{}
	}

	if mmAdd.defaultExpectation.params != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Expect")
	}

	if mmAdd.defaultExpectation.paramPtrs == nil {
		mmAdd.defaultExpectation.paramPtrs = &HabitCreatorMockAddParamPtrs{}
	}
	mmAdd.defaultExpectation.paramPtrs.ctx = &ctx
	mmAdd.defaultExpectation.expectationOrigins.originCtx = minimock.CallerInfo(1)

	return mmAdd
}

// ExpectHabitParam2 sets up expected param habit for habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) ExpectHabitParam2(habit mm_habit.Habit) *mHabitCreatorMockAdd {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{}
	}

	if mmAdd.defaultExpectation.params != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Expect")
	}

	if mmAdd.defaultExpectation.paramPtrs == nil {
		mmAdd.defaultExpectation.paramPtrs = &HabitCreatorMockAddParamPtrs{}
	}
	mmAdd.defaultExpectation.paramPtrs.habit = &habit
	mmAdd.defaultExpectation.expectationOrigins.originHabit = minimock.CallerInfo(1)

	return mmAdd
}

// Inspect accepts an inspector function that has same arguments as the habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Inspect(f func(ctx context.Context, habit mm_habit.Habit)) *mHabitCreatorMockAdd {
	if mmAdd.mock.inspectFuncAdd != nil {
		mmAdd.mock.t.Fatalf("Inspect function is already set for HabitCreatorMock.Add")
	}

	mmAdd.mock.inspectFuncAdd = f

	return mmAdd
}

// Return sets up results that will be returned by habitCreator.Add
func (mmAdd *mHabitCreatorMockAdd) Return(err error) *HabitCreatorMock {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	if mmAdd.defaultExpectation == nil {
		mmAdd.defaultExpectation = &HabitCreatorMockAddExpectation{mock: mmAdd.mock}
	}
	mmAdd.defaultExpectation.results = &HabitCreatorMockAddResults{err}
	mmAdd.defaultExpectation.returnOrigin = minimock.CallerInfo(1)
	return mmAdd.mock
}

// Set uses given function f to mock the habitCreator.Add method
func (mmAdd *mHabitCreatorMockAdd) Set(f func(ctx context.Context, habit mm_habit.Habit) (err error)) *HabitCreatorMock {
	if mmAdd.defaultExpectation != nil {
		mmAdd.mock.t.Fatalf("Default expectation is already set for the habitCreator.Add method")
	}

	if len(mmAdd.expectations) > 0 {
		mmAdd.mock.t.Fatalf("Some expectations are already set for the habitCreator.Add method")
	}

	mmAdd.mock.funcAdd = f
	mmAdd.mock.funcAddOrigin = minimock.CallerInfo(1)
	return mmAdd.mock
}

// When sets expectation for the habitCreator.Add which will trigger the result defined by the following
// Then helper
func (mmAdd *mHabitCreatorMockAdd) When(ctx context.Context, habit mm_habit.Habit) *HabitCreatorMockAddExpectation {
	if mmAdd.mock.funcAdd != nil {
		mmAdd.mock.t.Fatalf("HabitCreatorMock.Add mock is already set by Set")
	}

	expectation := &HabitCreatorMockAddExpectation{
		mock:               mmAdd.mock,
		params:             &HabitCreatorMockAddParams{ctx, habit},
		expectationOrigins: HabitCreatorMockAddExpectationOrigins{origin: minimock.CallerInfo(1)},
	}
	mmAdd.expectations = append(mmAdd.expectations, expectation)
	return expectation
}

// Then sets up habitCreator.Add return parameters for the expectation previously defined by the When method
func (e *HabitCreatorMockAddExpectation) Then(err error) *HabitCreatorMock {
	e.results = &HabitCreatorMockAddResults{err}
	return e.mock
}

// Times sets number of times habitCreator.Add should be invoked
func (mmAdd *mHabitCreatorMockAdd) Times(n uint64) *mHabitCreatorMockAdd {
	if n == 0 {
		mmAdd.mock.t.Fatalf("Times of HabitCreatorMock.Add mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmAdd.expectedInvocations, n)
	mmAdd.expectedInvocationsOrigin = minimock.CallerInfo(1)
	return mmAdd
}

func (mmAdd *mHabitCreatorMockAdd) invocationsDone() bool {
	if len(mmAdd.expectations) == 0 && mmAdd.defaultExpectation == nil && mmAdd.mock.funcAdd == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmAdd.mock.afterAddCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmAdd.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Add implements mm_habit.habitCreator
func (mmAdd *HabitCreatorMock) Add(ctx context.Context, habit mm_habit.Habit) (err error) {
	mm_atomic.AddUint64(&mmAdd.beforeAddCounter, 1)
	defer mm_atomic.AddUint64(&mmAdd.afterAddCounter, 1)

	mmAdd.t.Helper()

	if mmAdd.inspectFuncAdd != nil {
		mmAdd.inspectFuncAdd(ctx, habit)
	}

	mm_params := HabitCreatorMockAddParams{ctx, habit}

	// Record call args
	mmAdd.AddMock.mutex.Lock()
	mmAdd.AddMock.callArgs = append(mmAdd.AddMock.callArgs, &mm_params)
	mmAdd.AddMock.mutex.Unlock()

	for _, e := range mmAdd.AddMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAdd.AddMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAdd.AddMock.defaultExpectation.Counter, 1)
		mm_want := mmAdd.AddMock.defaultExpectation.params
		mm_want_ptrs := mmAdd.AddMock.defaultExpectation.paramPtrs

		mm_got := HabitCreatorMockAddParams{ctx, habit}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.ctx != nil && !minimock.Equal(*mm_want_ptrs.ctx, mm_got.ctx) {
				mmAdd.t.Errorf("HabitCreatorMock.Add got unexpected parameter ctx, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAdd.AddMock.defaultExpectation.expectationOrigins.originCtx, *mm_want_ptrs.ctx, mm_got.ctx, minimock.Diff(*mm_want_ptrs.ctx, mm_got.ctx))
			}

			if mm_want_ptrs.habit != nil && !minimock.Equal(*mm_want_ptrs.habit, mm_got.habit) {
				mmAdd.t.Errorf("HabitCreatorMock.Add got unexpected parameter habit, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
					mmAdd.AddMock.defaultExpectation.expectationOrigins.originHabit, *mm_want_ptrs.habit, mm_got.habit, minimock.Diff(*mm_want_ptrs.habit, mm_got.habit))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAdd.t.Errorf("HabitCreatorMock.Add got unexpected parameters, expected at\n%s:\nwant: %#v\n got: %#v%s\n",
				mmAdd.AddMock.defaultExpectation.expectationOrigins.origin, *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAdd.AddMock.defaultExpectation.results
		if mm_results == nil {
			mmAdd.t.Fatal("No results are set for the HabitCreatorMock.Add")
		}
		return (*mm_results).err
	}
	if mmAdd.funcAdd != nil {
		return mmAdd.funcAdd(ctx, habit)
	}
	mmAdd.t.Fatalf("Unexpected call to HabitCreatorMock.Add. %v %v", ctx, habit)
	return
}

// AddAfterCounter returns a count of finished HabitCreatorMock.Add invocations
func (mmAdd *HabitCreatorMock) AddAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.afterAddCounter)
}

// AddBeforeCounter returns a count of HabitCreatorMock.Add invocations
func (mmAdd *HabitCreatorMock) AddBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAdd.beforeAddCounter)
}

// Calls returns a list of arguments used in each call to HabitCreatorMock.Add.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAdd *mHabitCreatorMockAdd) Calls() []*HabitCreatorMockAddParams {
	mmAdd.mutex.RLock()

	argCopy := make([]*HabitCreatorMockAddParams, len(mmAdd.callArgs))
	copy(argCopy, mmAdd.callArgs)

	mmAdd.mutex.RUnlock()

	return argCopy
}

// MinimockAddDone returns true if the count of the Add invocations corresponds
// the number of defined expectations
func (m *HabitCreatorMock) MinimockAddDone() bool {
	if m.AddMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.AddMock.invocationsDone()
}

// MinimockAddInspect logs each unmet expectation
func (m *HabitCreatorMock) MinimockAddInspect() {
	for _, e := range m.AddMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to HabitCreatorMock.Add at\n%s with params: %#v", e.expectationOrigins.origin, *e.params)
		}
	}

	afterAddCounter := mm_atomic.LoadUint64(&m.afterAddCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.AddMock.defaultExpectation != nil && afterAddCounter < 1 {
		if m.AddMock.defaultExpectation.params == nil {
			m.t.Errorf("Expected call to HabitCreatorMock.Add at\n%s", m.AddMock.defaultExpectation.returnOrigin)
		} else {
			m.t.Errorf("Expected call to HabitCreatorMock.Add at\n%s with params: %#v", m.AddMock.defaultExpectation.expectationOrigins.origin, *m.AddMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAdd != nil && afterAddCounter < 1 {
		m.t.Errorf("Expected call to HabitCreatorMock.Add at\n%s", m.funcAddOrigin)
	}

	if !m.AddMock.invocationsDone() && afterAddCounter > 0 {
		m.t.Errorf("Expected %d calls to HabitCreatorMock.Add at\n%s but found %d calls",
			mm_atomic.LoadUint64(&m.AddMock.expectedInvocations), m.AddMock.expectedInvocationsOrigin, afterAddCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *HabitCreatorMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockAddInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *HabitCreatorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *HabitCreatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddDone()
}
