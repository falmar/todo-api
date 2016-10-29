// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

import "testing"

func TestPaging3(t *testing.T) {
	init := int64(0)
	max := int64(1)
	totalPages := int64(3)

	p := &paging{CurrentPage: 1, PageSize: 1}

	p.calc(3)

	if p.Init != init {
		t.Fatalf("Expected init to be %d; got: %d", init, p.Init)
	}

	if p.Max != max {
		t.Fatalf("Expected max to be %d; got: %d", max, p.Max)
	}

	if p.TotalPages != totalPages {
		t.Fatalf("Expected totalPages to be %d; got: %d", totalPages, p.TotalPages)
	}

}

func TestPaging5(t *testing.T) {
	init := int64(0)
	max := int64(5)
	totalPages := int64(1)

	p := &paging{CurrentPage: 1, PageSize: 15}

	p.calc(5)

	if p.Init != init {
		t.Fatalf("Expected init to be %d; got: %d", init, p.Init)
	}

	if p.Max != max {
		t.Fatalf("Expected max to be %d; got: %d", max, p.Max)
	}

	if p.TotalPages != totalPages {
		t.Fatalf("Expected totalPages to be %d; got: %d", totalPages, p.TotalPages)
	}

}

func TestPaging13(t *testing.T) {
	init := int64(0)
	max := int64(13)
	totalPages := int64(1)

	p := &paging{CurrentPage: 2, PageSize: 15}

	p.calc(13)

	if p.Init != init {
		t.Fatalf("Expected init to be %d; got: %d", init, p.Init)
	}

	if p.Max != max {
		t.Fatalf("Expected max to be %d; got: %d", max, p.Max)
	}

	if p.TotalPages != totalPages {
		t.Fatalf("Expected totalPages to be %d; got: %d", totalPages, p.TotalPages)
	}

}

func TestPaging25(t *testing.T) {
	init := int64(15)
	max := int64(10)
	totalPages := int64(2)

	p := &paging{CurrentPage: 2, PageSize: 15}

	p.calc(25)

	if p.Init != init {
		t.Fatalf("Expected init to be %d; got: %d", init, p.Init)
	}

	if p.Max != max {
		t.Fatalf("Expected max to be %d; got: %d", max, p.Max)
	}

	if p.TotalPages != totalPages {
		t.Fatalf("Expected totalPages to be %d; got: %d", totalPages, p.TotalPages)
	}

}

func TestPaging30(t *testing.T) {
	init := int64(30)
	max := int64(10)
	totalPages := int64(3)

	p := &paging{CurrentPage: 3, PageSize: 15}

	p.calc(40)

	if p.Init != init {
		t.Fatalf("Expected init to be %d; got: %d", init, p.Init)
	}

	if p.Max != max {
		t.Fatalf("Expected max to be %d; got: %d", max, p.Max)
	}

	if p.TotalPages != totalPages {
		t.Fatalf("Expected totalPages to be %d; got: %d", totalPages, p.TotalPages)
	}

}

func TestPages3(t *testing.T) {
	pages := []int64{1, 2, 3}

	p := &paging{CurrentPage: 2, PageSize: 1}

	p.calc(3)
	pgs := p.getPages()

	for i, v := range pages {
		if pgs[i] != v {
			t.Fatalf("Expected Pages[%d] to be %d; got: %d", i, v, p.Pages[i])
		}
	}

}

func TestPages20(t *testing.T) {
	pages := []int64{1, 2, 3, 4}

	p := &paging{CurrentPage: 4, PageSize: 5}

	p.calc(20)
	pgs := p.getPages()

	for i, v := range pages {
		if pgs[i] != v {
			t.Fatalf("Expected Pages[%d] to be %d; got: %d", i, v, p.Pages[i])
		}
	}

}

func TestPages40(t *testing.T) {
	pages := []int64{5, 6, 7, 8}

	p := &paging{CurrentPage: 6, PageSize: 5}

	p.calc(40)
	pgs := p.getPages()

	for i, v := range pages {
		if pgs[i] != v {
			t.Fatalf("Expected Pages[%d] to be %d; got: %d", i, v, p.Pages[i])
		}
	}

}

func TestPages30(t *testing.T) {
	pages := []int64{3, 4, 5, 6}

	p := &paging{CurrentPage: 4, PageSize: 5}

	p.calc(30)
	pgs := p.getPages()

	for i, v := range pages {
		if pgs[i] != v {
			t.Fatalf("Expected Pages[%d] to be %d; got: %d", i, v, p.Pages[i])
		}
	}

}

func TestPages45(t *testing.T) {
	pages := []int64{2, 3, 4, 5}

	p := &paging{CurrentPage: 4, PageSize: 10}

	p.calc(45)
	pgs := p.getPages()

	for i, v := range pages {
		if pgs[i] != v {
			t.Fatalf("Expected Pages[%d] to be %d; got: %d", i, v, p.Pages[i])
		}
	}

}
