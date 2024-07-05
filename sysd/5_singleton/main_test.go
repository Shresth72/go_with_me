package main_test

import (
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
	err := suite.boiler.Fill()
	assert.Nil(suite.T(), err, "fill should not timeout")
	time.Sleep(100 * time.Millisecond)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err = suite.boiler.Fill()
		assert.NotNil(suite.T(), err, "fill should timeout when boiler is not empty")
	}()

	wg.Wait()
	assert.False(suite.T(), suite.boiler.IsEmpty(), "boiler should not be empty after attempted fill")
	assert.False(suite.T(), suite.boiler.IsBoiled(), "boiler should not be boiled after attempted fill")
}

func (suite *ChocolateBoilerTestSuite) TestBoilWhenNotFilled() {
	suite.boiler.SetEmpty(true)
	suite.boiler.SetBoiled(false)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := suite.boiler.Boil()
		assert.NotNil(suite.T(), err, "boil should timout when boiler is empty")
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Wait()
	assert.True(suite.T(), suite.boiler.IsEmpty(), "Boiler should be empty after attempted boil")
	assert.False(suite.T(), suite.boiler.IsBoiled(), "Boiler should not be boiled after attempted boil")
}
