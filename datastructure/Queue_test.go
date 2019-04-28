package datastructure

import (
	"code.byted.org/gopkg/pkg/testing/assert"
	"testing"
)

func TestQueue(t *testing.T)  {

	queue := NewQueue(3)

	var err error
	err = queue.Push(&Node{"1"})
	assert.Nil(t, err)

	err = queue.Push(&Node{"2"})
	assert.Nil(t, err)

	err = queue.Push(&Node{"3"})
	assert.Nil(t, err)

	err = queue.Push(&Node{"4"})
	assert.True(t, err != nil)

	var value *Node
	value = queue.Pop()
	assert.Equal(t, value.value, "1")

	value = queue.Pop()
	assert.Equal(t, value.value, "2")

	value = queue.Pop()
	assert.Equal(t, value.value, "3")

	value = queue.Pop()
	assert.True(t, value == nil)

}
