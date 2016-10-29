// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

type paging struct {
	Init         int64
	Max          int64
	CurrentPage  int64
	PageSize     int64
	Pages        []int64
	Results      int64
	TotalResults int64
	TotalPages   int64
}

func (p *paging) calc(results int64) {
	if results <= 0 {
		return
	}

	if p.PageSize <= 0 {
		p.PageSize = 15
	}

	for p.CurrentPage > 0 && (p.CurrentPage-1)*p.PageSize > results {
		p.CurrentPage--
	}

	pageMax := p.PageSize
	var pageInit int64

	p.TotalResults = results
	p.TotalPages = p.getTotalPages(results, p.PageSize)
	p.Pages = p.getPages(p.CurrentPage, p.TotalPages)

	p.Max = p.PageSize

	if p.Max > results {
		p.Max = results
	}

	if p.CurrentPage > 1 {
		p.Init = (p.CurrentPage - 1) * p.PageSize

		if results-p.Init < p.Max {
			p.Max = results - p.Init
		}

		pageInit = p.Init
		pageMax = pageInit + p.Max
	} else {
		pageMax = pageInit + p.PageSize
	}

	if pageMax >= results {
		pageMax = results
	}

	p.Results = pageMax - pageInit

}

func (p paging) getPages(currentPage, totalPages int64) []int64 {
	if totalPages > 0 {
		if (currentPage-2) > 0 && (currentPage+2) < totalPages {
			return p.rng(currentPage-2, currentPage+2)
		}

		if totalPages >= 5 {
			if currentPage-2 <= 0 {
				return p.rng(1, 4)
			}

			return p.rng(totalPages-3, totalPages)
		}

		return p.rng(1, totalPages)
	}

	return nil
}

func (p paging) getTotalPages(results, pageSize int64) int64 {
	if results > 0 {
		if results <= pageSize {
			return 1
		}

		if (results % pageSize) > 0 {
			return (results / pageSize) + 1
		}

		return (results / pageSize)
	}

	return 0
}

func (p paging) rng(init, max int64) []int64 {
	pages := make([]int64, (max-init)+1)

	i := 0
	for f := init; f <= max; f++ {
		pages[i] = f
		i++
	}

	return pages
}
