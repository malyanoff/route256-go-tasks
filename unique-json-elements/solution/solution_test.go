package solution


import (
	"testing"
)


var jsonStr = `
[
  {
    "data": {
      "key": "two"
    }
  },
  {
    "data": {
      "key": "three"
    }
  },
  {
    "data": {
      "key": "four"
    },
    "meta": {
      "key": "hello"
    }
  },
  {
	  "data":{
		  "key":"two"
	  },
	  "data1":{
		  "key":"5454"
	  }
  }
]
`

func TestCountUniqueKeys(t *testing.T){

    got := CountUniqueKeys(jsonStr)
    want := 3

    if got != want {
        t.Errorf("got %v, wanted %v", got, want)
    }
}