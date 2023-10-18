package main

type User struct {
	field3 bool
	field4 struct {
		nested1 int
		nested2 []string
		nested3 struct {
			nested1 int
			nested2 []string
			nested3 string
		}
	}
	field5 []int
	field1 int
	field2 string
}
