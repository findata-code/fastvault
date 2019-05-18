package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDateNow(t *testing.T){
	now := time.Now().Unix()
	fmt.Println(now)
}