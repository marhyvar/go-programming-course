package main

import "testing"

func TestArea(t *testing.T) {
        result := Area(2, 2)
        if result != 4 {
                t.Errorf("Function Area(2, 2) = %d; want 4", result)
        }
}

