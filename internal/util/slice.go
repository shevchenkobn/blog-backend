package util

import "../services/onexit"

func RemoveCallbackOnOrder(s []onexit.Callback, callback onexit.Callback) []onexit.Callback {
	for i, c := range s {
		if &c == &callback {
			return RemoveINoOrder(s, i)
		}
	}
	return s
}

func RemoveINoOrder(s []onexit.Callback, i int) []onexit.Callback {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}
