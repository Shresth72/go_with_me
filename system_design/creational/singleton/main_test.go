package main_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/Shresth72/sysd/sysd/5_singleton"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ChocolateBoilerTestSuite struct {
	suite.Suite
	boiler *main.ChocolateBoiler
}

func (suite *ChocolateBoilerTestSuite) SetupTest() {
	suite.boiler = main.NewChocolateBoiler()
	// Reset boiler state
	suite.boiler.SetEmpty(true)
	suite.boiler.SetBoiled(false)
}

func TestChocolateBoilerTestSuite(t *testing.T) {
	suite.Run(t, new(ChocolateBoilerTestSuite))
}

func (suite *ChocolateBoilerTestSuite) TestFill() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := suite.boiler.Fill()
		assert.Nil(suite.T(), err, "fill should not timout")
	}()

	wg.Wait()
	assert.False(suite.T(), suite.boiler.IsEmpty(), "boiler should not be empty after filling")
	assert.False(suite.T(), suite.boiler.IsBoiled(), "boiler should not be boiled after filling")
}

func (suite *ChocolateBoilerTestSuite) TestBoil() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := suite.boiler.Fill()
		assert.Nil(suite.T(), err, "fill should not timeout")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		err := suite.boiler.Boil()
		assert.Nil(suite.T(), err, "boil should not timeout")
	}()

	wg.Wait()
	assert.True(suite.T(), suite.boiler.IsBoiled(), "boiler should be boiled after boiling")
}

func (suite *ChocolateBoilerTestSuite) TestDrain() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		err := suite.boiler.Fill()
		assert.Nil(suite.T(), err, "fill should not timeout")
	}()

	go func() {
		defer wg.Done()
		err := suite.boiler.Boil()
		assert.Nil(suite.T(), err, "boil should not timeout")
	}()

	go func() {
		defer wg.Done()
		err := suite.boiler.Drain()
		assert.Nil(suite.T(), err, "drain should not timeout")
	}()

	wg.Wait()
	assert.True(suite.T(), suite.boiler.IsEmpty(), "boiler should be empty after draining")
	assert.False(suite.T(), suite.boiler.IsBoiled(), "boiler should not be boiled after draining")
}

func (suite *ChocolateBoilerTestSuite) TestFillWhenNotEmpty() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := suite.boiler.Fill()
	assert.Nil(suite.T(), err, "fill should not timeout")
	time.Sleep(100 * time.Millisecond)

	done := make(chan bool)
	go func() {
		suite.boiler.Fill()
		done <- true
	}()

	select {
	case <-ctx.Done():
		assert.True(suite.T(), true, "fill should timeout after already being filled")
		assert.False(suite.T(), suite.boiler.IsEmpty(), "boiler should not be empty after attempted fill")
		assert.False(suite.T(), suite.boiler.IsBoiled(), "boiler should not be boiled after attempted fill")
	case <-done:
		assert.Fail(suite.T(), "fill function returned unexpectedly")
	}
}

func (suite *ChocolateBoilerTestSuite) TestBoilWhenNotFilled() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	done := make(chan bool)

	go func() {
		suite.boiler.Boil()
		done <- true
	}()

	select {
	case <-ctx.Done():
		assert.True(suite.T(), true, "Boil request should timeout after 3 seconds")
		assert.True(suite.T(), suite.boiler.IsEmpty(), "Boiler should be empty after attempted boil")
		assert.False(suite.T(), suite.boiler.IsBoiled(), "Boiler should not be boiled after attempted boil")
	case <-done:
		assert.Fail(suite.T(), "Boil function returned unexpectedly")
	}
}
