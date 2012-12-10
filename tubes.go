
package main

import (
  "fmt"
  "container/list"
  "time"
)

type Tube struct {
  channel chan string
  subscribers *list.List
  name string
}

func newTube(name string) *Tube {
  t := new(Tube)
  t.channel = make(chan string, 8)
  t.subscribers = list.New()
  t.name = name
  return t
}

func (t *Tube) Sub(target *Tube) {
  target.subscribers.PushBack(t)
}

func (t *Tube) Pub(s string) {
  fmt.Println("pub!")
  // send on this tube's channel
  t.channel <- s

  for e := t.subscribers.Front(); e != nil; e = e.Next() {
    fmt.Println("s?")
    t2 := e.Value.(*Tube)
    fmt.Println(t2.name)
    // send on the tube's subscribed channel
    t2.channel <- s
  }
}

func main() {
  fmt.Println("你好, 世界")

  t1 := newTube("t1")
  t2 := newTube("t2")

  t2.Sub(t1)

  t1.Pub("bracatinga")
  t2.Pub("pirarucu")
  t1.Pub("capivara")
  t2.Pub("ipe")

  go out(t2)

  time.Sleep(1 * 1e9)
  fmt.Println("再见！")

}


func out(t *Tube) {
  for i := 0; i < 4; i++ {
    fmt.Println(<-t.channel)
  }
}

