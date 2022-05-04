package models

import ("time"; "errors")

var ErrNoRecord = errors.New("models: no matching record Found")

type Snippet struct {
  ID string
  Title string
  Content string
  Created time.Time
  Expires time.Time
  
}