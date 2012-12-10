
package main

import (
  "fmt"
  "container/list"
  "time"
)

type Tube struct {
  in chan string
  out chan string
  subscribers *list.List
  name string
}

func newTube(name string) *Tube {
  t := new(Tube)
  t.in = make(chan string, 8)
  t.out = make(chan string, 8)
  t.subscribers = list.New()
  t.name = name

  // takes from 'in' and put on all subscriber's 'out'
  go t.Multiplier()

  return t
}

func (t *Tube) Sub(target *Tube) {
  target.subscribers.PushBack(t)
}

// forever running multiplier goroutine
func (t *Tube) Multiplier() {
  for {
    // take from the in channel
    s := <-t.in
    // send on the out channel
    t.out <- s
    // and on the subscriber's in channel
    for e := t.subscribers.Front(); e != nil; e = e.Next() {
      t2 := e.Value.(*Tube)
      // send on the tube's subscribed channels
      t2.in <- s
    }
  }
}

func main() {
  fmt.Println("你好, 世界")

  // luigi brings the tubes
  t1 := newTube("t1")
  t2 := newTube("t2")
  t3 := newTube("t3")

  // mario conects them
  t2.Sub(t1)
  t3.Sub(t2)

  // por favor no hechar el papel en el baño. (taza)
  t1.in <- "bracatinga"
  t2.in <- "pirarucu"
  t3.in <- "capivara"

  go out(t3)

  time.Sleep(1 * 1e9)
  fmt.Println("再见！")

}


func out(t *Tube) {
  for i := 0; i < 4; i++ {
    fmt.Println(<-t.out)
  }
}

