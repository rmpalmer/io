package main 

import (
	"fmt"
	"dscout"
	"dscin"
	"records"
)

func main() {
	fmt.Printf("io test starting.\n")
    d := dscout.NewDscout("bar.gz")
    t := records.NewTrace(6,16)
    for i:=0; i<4; i++ {
    	t.Header[i] = i
    }
    for i:=0; i<10; i++ {
    	t.Data[i] = 10.5 - float64(i) 
    }
    for i:= 0; i<5; i++ {
    	t.Header[0] = i
	    d.HandleTrace(t)
    }
    d.HandleEod()
    
    e := dscin.NewDscin("bar.gz")
    for {
    	t = e.HandleTrace()
    	if (t == nil) {
    		break
    	} else {
			t.Summarize()
			t.Detail()
    	}
    }
    
	fmt.Printf("io testy done\n")
}

