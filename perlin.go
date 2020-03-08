package main

import (
   "fmt"
   "time"

   //"github.com/aquilax/go-perlin"
   "github.com/cjslep/noise"
)




func main() {
    var data []float64
   x := 10
   y := 10

   //p:= perlin.NewPerlin(2,2,10, time.Now().UnixNano())

   for i := 0; i < x; i++{
      for j := 0; j < y; j++{

         p:= noise.NewPerlin( time.Now().UnixNano())
         //p:= noise.NewPerlin(1)

         time.Sleep(1)
         calc := p.Noise(float64(i/8.),float64(j/8.))

         data = append(data,calc)

      }
   }
   fmt.Println(data)
}
