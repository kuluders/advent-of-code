package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  // "slices"
  "math"
  "strconv"
)

func main() {
  // Step 1: Open the file
  f, err := os.Open("input-test.txt")
  if err != nil {
    fmt.Println(err)
    return
  }

  // Step 2: Create a reader
  r := bufio.NewReader(f)

  var reports [][]string
  // Step 3: Read line by line
  for {
      line, err := r.ReadString('\n')
      if err != nil {
          break
      }
      reports = append(reports, strings.Fields(line))
  }

  count_safe := 0
  for i := 0; i < len(reports); i++ {
    last := 0.0
    up := 0.0
    safe := true
    skipped := false
    // fmt.Println(reports[i])
    for j := 0; j < len(reports[i]); j++{
      current := 0.0
      if s, err := strconv.ParseFloat(reports[i][j], 64); err == nil {
        current = s
      }
      if last == 0 {
        last = current
      }else{
        difference := current - last
        direction := 0.0
        
        if difference < 0 {
          direction = -1
        }else{
          direction = 1
        }
        
        fmt.Println(last, current, difference)
        if up == 0 {
          up = direction
        }else if up != direction{
          if !safe {
            unsafe = true
            fmt.Println("marking unsafe to true")
            break
          }
          safe = false
          fmt.Println("marking safe to false")
        }
        last = current
        difference = math.Abs(difference)
        if difference > 3 || difference < 1{
          if !safe {
            unsafe = true
            fmt.Println("marking unsafe to true")
            break
          }
          safe = false
          fmt.Println("marking safe to false")
        }
      }
    }
    fmt.Println(safe)
    fmt.Println(unsafe)
    if !unsafe {
      fmt.Println("adding")
      count_safe += 1
    }
  }

  // fmt.Println(reports)
  fmt.Println(count_safe)
  defer f.Close()
}