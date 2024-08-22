package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/nqxcode/platform_common/client/db.TxManager -o ./mocks\tx_manager_minimock.go -n TxManagerMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	mm_db "github.com/nqxcode/platform_common/client/db"
)

// TxManagerMock implements db.TxManager
type TxManagerMock struct {
	t minimock.Tester

	funcReadCommitted          func(ctx context.Context, f mm_db.Handler) (err error)
	inspectFuncReadCommitted   func(ctx context.Context, f mm_db.Handler)
	afterReadCommittedCounter  uint64
	beforeReadCommittedCounter uint64
	ReadCommittedMock          mTxManagerMockReadCommitted
}

// NewTxManagerMock returns a mock for db.TxManager
func NewTxManagerMock(t minimock.Tester) *TxManagerMock {
	m := &TxManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.ReadCommittedMock = mTxManagerMockReadCommitted{mock: m}
	m.ReadCommittedMock.callArgs = []*TxManagerMockReadCommittedParams{}

	return m
}

type mTxManagerMockReadCommitted struct {
	mock               *TxManagerMock
	defaultExpectation *TxManagerMockReadCommittedExpectation
	expectations       []*TxManagerMockReadCommittedExpectation

	callArgs []*TxManagerMockReadCommittedParams
	mutex    sync.RWMutex
}

// TxManagerMockReadCommittedExpectation specifies expectation struct of the TxManager.ReadCommitted
type TxManagerMockReadCommittedExpectation struct {
	mock    *TxManagerMock
	params  *TxManagerMockReadCommittedParams
	results *TxManagerMockReadCommittedResults
	Counter uint64
}

// TxManagerMockReadCommittedParams contains parameters of the TxManager.ReadCommitted
type TxManagerMockReadCommittedParams struct {
	ctx context.Context
	f   mm_db.Handler
}

// TxManagerMockReadCommittedResults contains results of the TxManager.ReadCommitted
type TxManagerMockReadCommittedResults struct {
	err error
}

// Expect sets up expected params for TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Expect(ctx context.Context, f mm_db.Handler) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{}
	}

	mmReadCommitted.defaultExpectation.params = &TxManagerMockReadCommittedParams{ctx, f}
	for _, e := range mmReadCommitted.expectations {
		if minimock.Equal(e.params, mmReadCommitted.defaultExpectation.params) {
			mmReadCommitted.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmReadCommitted.defaultExpectation.params)
		}
	}

	return mmReadCommitted
}

// Inspect accepts an inspector function that has same arguments as the TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Inspect(f func(ctx context.Context, f mm_db.Handler)) *mTxManagerMockReadCommitted {
	if mmReadCommitted.mock.inspectFuncReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("Inspect function is already set for TxManagerMock.ReadCommitted")
	}

	mmReadCommitted.mock.inspectFuncReadCommitted = f

	return mmReadCommitted
}

// Return sets up results that will be returned by TxManager.ReadCommitted
func (mmReadCommitted *mTxManagerMockReadCommitted) Return(err error) *TxManagerMock {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	if mmReadCommitted.defaultExpectation == nil {
		mmReadCommitted.defaultExpectation = &TxManagerMockReadCommittedExpectation{mock: mmReadCommitted.mock}
	}
	mmReadCommitted.defaultExpectation.results = &TxManagerMockReadCommittedResults{err}
	return mmReadCommitted.mock
}

// Set uses given function f to mock the TxManager.ReadCommitted method
func (mmReadCommitted *mTxManagerMockReadCommitted) Set(f func(ctx context.Context, f mm_db.Handler) (err error)) *TxManagerMock {
	if mmReadCommitted.defaultExpectation != nil {
		mmReadCommitted.mock.t.Fatalf("Default expectation is already set for the TxManager.ReadCommitted method")
	}

	if len(mmReadCommitted.expectations) > 0 {
		mmReadCommitted.mock.t.Fatalf("Some expectations are already set for the TxManager.ReadCommitted method")
	}

	mmReadCommitted.mock.funcReadCommitted = f
	return mmReadCommitted.mock
}

// When sets expectation for the TxManager.ReadCommitted which will trigger the result defined by the following
// Then helper
func (mmReadCommitted *mTxManagerMockReadCommitted) When(ctx context.Context, f mm_db.Handler) *TxManagerMockReadCommittedExpectation {
	if mmReadCommitted.mock.funcReadCommitted != nil {
		mmReadCommitted.mock.t.Fatalf("TxManagerMock.ReadCommitted mock is already set by Set")
	}

	expectation := &TxManagerMockReadCommittedExpectation{
		mock:   mmReadCommitted.mock,
		params: &TxManagerMockReadCommittedParams{ctx, f},
	}
	mmReadCommitted.expectations = append(mmReadCommitted.expectations, expectation)
	return expectation
}

// Then sets up TxManager.ReadCommitted return parameters for the expectation previously defined by the When method
func (e *TxManagerMockReadCommittedExpectation) Then(err error) *TxManagerMock {
	e.results = &TxManagerMockReadCommittedResults{err}
	return e.mock
}

// ReadCommitted implements db.TxManager
func (mmReadCommitted *TxManagerMock) ReadCommitted(ctx context.Context, f mm_db.Handler) (err error) {
	mm_atomic.AddUint64(&mmReadCommitted.beforeReadCommittedCounter, 1)
	defer mm_atomic.AddUint64(&mmReadCommitted.afterReadCommittedCounter, 1)

	if mmReadCommitted.inspectFuncReadCommitted != nil {
		mmReadCommitted.inspectFuncReadCommitted(ctx, f)
	}

	mm_params := &TxManagerMockReadCommittedParams{ctx, f}

	// Record call args
	mmReadCommitted.ReadCommittedMock.mutex.Lock()
	mmReadCommitted.ReadCommittedMock.callArgs = append(mmReadCommitted.ReadCommittedMock.callArgs, mm_params)
	mmReadCommitted.ReadCommittedMock.mutex.Unlock()

	for _, e := range mmReadCommitted.ReadCommittedMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmReadCommitted.ReadCommittedMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmReadCommitted.ReadCommittedMock.defaultExpectation.Counter, 1)
		mm_want := mmReadCommitted.ReadCommittedMock.defaultExpectation.params
		mm_got := TxManagerMockReadCommittedParams{ctx, f}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmReadCommitted.t.Errorf("TxManagerMock.ReadCommitted got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmReadCommitted.ReadCommittedMock.defaultExpectation.results
		if mm_results == nil {
			mmReadCommitted.t.Fatal("No results are set for the TxManagerMock.ReadCommitted")
		}
		return (*mm_results).err
	}
	if mmReadCommitted.funcReadCommitted != nil {
		return mmReadCommitted.funcReadCommitted(ctx, f)
	}
	mmReadCommitted.t.Fatalf("Unexpected call to TxManagerMock.ReadCommitted. %v %v", ctx, f)
	return
}

// ReadCommittedAfterCounter returns a count of finished TxManagerMock.ReadCommitted invocations
func (mmReadCommitted *TxManagerMock) ReadCommittedAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadCommitted.afterReadCommittedCounter)
}

// ReadCommittedBeforeCounter returns a count of TxManagerMock.ReadCommitted invocations
func (mmReadCommitted *TxManagerMock) ReadCommittedBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmReadCommitted.beforeReadCommittedCounter)
}

// Calls returns a list of arguments used in each call to TxManagerMock.ReadCommitted.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmReadCommitted *mTxManagerMockReadCommitted) Calls() []*TxManagerMockReadCommittedParams {
	mmReadCommitted.mutex.RLock()

	argCopy := make([]*TxManagerMockReadCommittedParams, len(mmReadCommitted.callArgs))
	copy(argCopy, mmReadCommitted.callArgs)

	mmReadCommitted.mutex.RUnlock()

	return argCopy
}

// MinimockReadCommittedDone returns true if the count of the ReadCommitted invocations corresponds
// the number of defined expectations
func (m *TxManagerMock) MinimockReadCommittedDone() bool {
	for _, e := range m.ReadCommittedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadCommittedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadCommittedCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReadCommitted != nil && mm_atomic.LoadUint64(&m.afterReadCommittedCounter) < 1 {
		return false
	}
	return true
}

// MinimockReadCommittedInspect logs each unmet expectation
func (m *TxManagerMock) MinimockReadCommittedInspect() {
	for _, e := range m.ReadCommittedMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to TxManagerMock.ReadCommitted with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ReadCommittedMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterReadCommittedCounter) < 1 {
		if m.ReadCommittedMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to TxManagerMock.ReadCommitted")
		} else {
			m.t.Errorf("Expected call to TxManagerMock.ReadCommitted with params: %#v", *m.ReadCommittedMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcReadCommitted != nil && mm_atomic.LoadUint64(&m.afterReadCommittedCounter) < 1 {
		m.t.Error("Expected call to TxManagerMock.ReadCommitted")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *TxManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockReadCommittedInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *TxManagerMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *TxManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockReadCommittedDone()
}
