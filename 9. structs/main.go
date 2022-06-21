package main1

import "fmt"

type s_struct struct {
	On		bool
	Ammo	int
	Power	int
}

func (s *s_struct) Shoot() bool {
	if s.On == false {
		return false
	} else if s.Ammo > 0 {
		s.Ammo -= 1
		return true
	} else {
		return false
	}
}

func (s *s_struct) RideBike() bool {
	if s.On == false {
		return false
	} else if s.Power > 0 {
		s.Power -= 1
		return true
	} else {
		return false
	}
}