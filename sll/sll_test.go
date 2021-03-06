package sll_test

import (
	"testing"

	"github.com/jamiealquiza/bicache/sll"
)

func TestHead(t *testing.T) {
	s := sll.New(5)

	node := s.PushHead("value")
	if s.Head() != node {
		t.Error("Unexpected head node")
	}
}

func TestTail(t *testing.T) {
	s := sll.New(5)

	node := s.PushTail("value")
	if s.Tail() != node {
		t.Error("Unexpected tail node")
	}
}

func TestRead(t *testing.T) {
	s := sll.New(5)

	s.PushHead("value")
	if s.Head().Read() != "value" {
		t.Error("Read method failed")
	}
}

func TestPushHead(t *testing.T) {
	s := sll.New(5)

	s.PushHead("value")
	if s.Head().Read() != "value" {
		t.Errorf(`Expected value "value", got "%s"`, s.Head().Read())
	}
}

func TestPushTail(t *testing.T) {
	s := sll.New(5)

	s.PushTail("value")
	if s.Tail().Read() != "value" {
		t.Errorf(`Expected value "value", got "%s"`, s.Tail().Read())
	}
}

func TestNext(t *testing.T) {
	s := sll.New(5)

	firstVal := "first"
	secondVal := "second"

	first := s.PushTail(firstVal)
	second := s.PushTail(secondVal)

	if second.Next() != first {
		t.Errorf("Expected node with value %s next, got %s", firstVal, second.Next().Read())
	}
}

func TestPrev(t *testing.T) {
	s := sll.New(5)

	firstVal := "first"
	secondVal := "second"

	first := s.PushTail(firstVal)
	second := s.PushTail(secondVal)

	if first.Prev() != second {
		t.Errorf("Expected node with value %s next, got %s", firstVal, second.Next().Read())
	}
}

func TestLen(t *testing.T) {
	s := sll.New(10)

	for i := 0; i < 5; i++ {
		s.PushTail(i)
	}

	if s.Len() != 5 {
		t.Errorf("Expected len 5, got %d", s.Len())
	}
}

func TestHighScores(t *testing.T) {
	s := sll.New(10)

	nodes := map[int]*sll.Node{}

	for i := 0; i < 5; i++ {
		nodes[i] = s.PushTail(i)
	}

	nodes[3].Read()
	nodes[3].Read()
	nodes[3].Read()

	nodes[4].Read()
	nodes[4].Read()

	// Should result in [2, 4, 3, 1, 5] with read scores
	// 3, 2, 0, 0, 0 respectively.
	scores := s.HighScores(3)

	if scores[0] != nodes[2] {
		t.Errorf("Expected scores position 0 node with value 2, got %d", scores[0].Read())
	}

	if scores[1] != nodes[4] {
		t.Errorf("Expected scores position 1 node with value 4, got %d", scores[1].Read())
	}

	if scores[2] != nodes[3] {
		t.Errorf("Expected scores position 2 node with value 3, got %d", scores[2].Read())
	}
}

func TestLowScores(t *testing.T) {
	s := sll.New(3)

	nodes := map[int]*sll.Node{}

	for i := 0; i < 3; i++ {
		nodes[i] = s.PushTail(i)
	}

	nodes[1].Read()
	nodes[1].Read()
	nodes[1].Read()

	nodes[0].Read()
	nodes[0].Read()

	// Should result in [2, 0, 1]
	// with read scores of 0, 2, 3 respectively.
	scores := s.LowScores(3)

	if scores[0] != nodes[2] {
		t.Errorf("Expected scores position 0 node with value 2, got %d", scores[2].Read())
	}

	if scores[1] != nodes[0] {
		t.Errorf("Expected scores position 1 node with value 4, got %d", scores[0].Read())
	}

	if scores[2] != nodes[1] {
		t.Errorf("Expected scores position 2 node with value 3, got %d", scores[1].Read())
	}
}

func TestMoveToHead(t *testing.T) {
	s := sll.New(10)

	for i := 0; i < 10; i++ {
		s.PushTail(i)
	}

	node := s.Tail().Next().Next()
	s.MoveToHead(node)

	// Check head method.
	if s.Head() != node {
		t.Errorf(`Expected node with value "%d" at head, got "%d"`,
			node.Read(), s.Head().Read())
	}

	// Check total order.
	expected := []int{9, 8, 6, 5, 4, 3, 2, 1, 0, 7}
	var i int
	for n := s.Tail(); n != nil; n = n.Next() {
		if n.Read() != expected[i] {
			t.Errorf(`Expected node with vallue "%d", got "%d"`, expected[i], n.Read())
		}
		i++
	}
}

func TestMoveToTail(t *testing.T) {
	s := sll.New(10)

	for i := 0; i < 10; i++ {
		s.PushTail(i)
	}

	node := s.Tail().Next().Next()
	s.MoveToTail(node)

	// Check tail method.
	if s.Tail() != node {
		t.Errorf(`Expected node with value "%d" at tail, got "%d"`,
			node.Read(), s.Tail().Read())
	}

	// Check total order.
	expected := []int{7, 9, 8, 6, 5, 4, 3, 2, 1, 0}
	var i int
	for n := s.Tail(); n != nil; n = n.Next() {
		if n.Read() != expected[i] {
			t.Errorf(`Expected node with value "%d", got "%d"`, expected[i], n.Read())
		}
		i++
	}
}

func TestPushHeadNode(t *testing.T) {
	s1 := sll.New(3)
	s2 := sll.New(3)

	s1.PushTail("target")
	node := s1.Tail()
	s1.Remove(node)

	s2.PushHead("value")
	s2.PushHead("value")
	s2.PushHeadNode(node)

	// Check from the head.
	if s2.Head() != node {
		t.Errorf(`Expected node with value "target", got "%s"`, s2.Head().Read())
	}

	// Ensure the links are correct.
	if s2.Tail().Next().Next() != node {
		t.Errorf(`Expected node with value "target", got "%s"`, s2.Tail().Next().Next().Read())
	}
}

func TestPushTailNode(t *testing.T) {
	s1 := sll.New(3)
	s2 := sll.New(3)

	s1.PushTail("target")
	node := s1.Tail()
	s1.Remove(node)

	s2.PushHead("value")
	s2.PushHead("value")
	s2.PushTailNode(node)

	// Check from the head.
	if s2.Tail() != node {
		t.Errorf(`Expected node with value "target", got "%s"`, s2.Tail().Read())
	}

	// Ensure the links are correct.
	if s2.Head().Prev().Prev() != node {
		t.Errorf(`Expected node with value "target", got "%s"`, s2.Head().Prev().Prev().Read())
	}
}

func TestRemove(t *testing.T) {
	s := sll.New(3)

	nodes := map[int]*sll.Node{}

	for i := 0; i < 3; i++ {
		nodes[i] = s.PushTail(i)
	}

	s.Remove(nodes[1])

	if s.Tail().Next().Read() != 0 {
		t.Errorf(`Expected node with value "0", got "%d"`, s.Tail().Next().Read())
	}

	scores := s.HighScores(3)

	if len(scores) != 2 {
		t.Errorf("Expected scores len 2, got %d", len(scores))
	}

	// This effectively tests the unexported
	// removeFromScores method.
	if scores[0] != s.Tail() {
		t.Error("Unexpected node in scores position 0")
	}

	if scores[1] != s.Tail().Next() {
		t.Error("Unexpected node in scores position 1")
	}
}

func TestSync(t *testing.T) {
	s := sll.New(3)

	first := s.PushTail("value")
	second := s.PushTail("value")
	s.PushTail("value")

	s.RemoveAsync(second)

	if s.Len() != 3 {
		t.Errorf("Expected len 3, got %d", s.Len())
	}

	if s.Tail().Next() != first {
		t.Error("Unexpected list order")
	}

	s.Sync()

	if s.Len() != 2 {
		t.Errorf("Expected len 2, got %d", s.Len())
	}
}

func TestRemoveAsync(t *testing.T) {
	s := sll.New(3)

	nodes := map[int]*sll.Node{}

	for i := 0; i < 3; i++ {
		nodes[i] = s.PushTail(i)
	}

	s.RemoveAsync(nodes[1])

	if s.Tail().Next().Read() != 0 {
		t.Errorf(`Expected node with value "0", got "%d"`, s.Tail().Next().Read())
	}

	if s.Len() != 3 {
		t.Errorf("Expected len 3, got %d", s.Len())
	}

	s.Sync()

	scores := s.HighScores(3)

	if s.Len() != 2 {
		t.Errorf("Expected len 2, got %d", s.Len())
	}

	if scores[0] != s.Tail() {
		t.Error("Unexpected node in scores position 0")
	}

	if scores[1] != s.Tail().Next() {
		t.Error("Unexpected node in scores position 1")
	}
}

func TestRemoveHead(t *testing.T) {
	s := sll.New(3)

	s.PushTail("value")
	target := s.PushTail("value")
	s.PushTail("value")

	s.RemoveHead()

	if s.Head() != target {
		t.Error("Unexpected head node")
	}
}

func TestRemoveTail(t *testing.T) {
	s := sll.New(3)

	s.PushTail("value")
	target := s.PushTail("value")
	s.PushTail("value")

	s.RemoveTail()

	if s.Tail() != target {
		t.Error("Unexpected tail node")
	}
}

func TestRemoveHeadAsync(t *testing.T) {
	s := sll.New(3)

	s.PushTail("value")
	target := s.PushTail("value")
	s.PushTail("value")

	s.RemoveHeadAsync()

	if s.Head() != target {
		t.Error("Unexpected head node")
	}

	if s.Len() != 3 {
		t.Errorf("Expected len 3, got %d", s.Len())
	}

	s.Sync()

	if s.Len() != 2 {
		t.Errorf("Expected len 2, got %d", s.Len())
	}
}

func TestRemoveTailAsync(t *testing.T) {
	s := sll.New(3)

	s.PushTail("value")
	target := s.PushTail("value")
	s.PushTail("value")

	s.RemoveTailAsync()

	if s.Tail() != target {
		t.Error("Unexpected tail node")
	}

	if s.Len() != 3 {
		t.Errorf("Expected len 3, got %d", s.Len())
	}

	s.Sync()

	if s.Len() != 2 {
		t.Errorf("Expected len 2, got %d", s.Len())
	}
}
