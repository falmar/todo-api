// Copyright 2016 David Lavieri.  All rights reserved.
// Use of this source code is governed by a MIT License
// License that can be found in the LICENSE file.

package main

// PageURL returns a formatted url
type pageURL struct {
	Page int64  `json:"page"`
	Link string `json:"link"`
}

type paging struct {
	Init         int64              `json:"-"`
	Max          int64              `json:"-"`
	CurrentPage  int64              `json:"current_page"`
	PageSize     int64              `json:"page_size"`
	Pages        []int64            `json:"-"`
	PagesURL     []pageURL          `json:"pages"`
	Results      int64              `json:"results"`
	TotalResults int64              `json:"total_results"`
	TotalPages   int64              `json:"total_pages"`
	Links        map[string]pageURL `json:"links"`
}

func (p *paging) calc(results int64) {
	if results <= 0 {
		return
	}

	if p.PageSize <= 0 {
		p.PageSize = 15
	}

	if p.CurrentPage <= 0 {
		p.CurrentPage = 1
	}

	for p.CurrentPage > 1 && (p.CurrentPage-1)*p.PageSize > results {
		p.CurrentPage--
	}

	pageMax := p.PageSize
	var pageInit int64

	p.TotalResults = results
	p.TotalPages = p.getTotalPages()

	p.Max = p.PageSize
	pageMax = pageInit + p.PageSize

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
	}

	if pageMax >= results {
		pageMax = results
	}

	p.Results = pageMax - pageInit
}

func (p paging) getPages() []int64 {
	if p.TotalPages > 0 {
		if (p.CurrentPage-2) > 0 && (p.CurrentPage+2) < p.TotalPages {
			return p.rng(p.CurrentPage-2, p.CurrentPage+2)
		}

		if p.TotalPages >= 5 {
			if p.CurrentPage-2 <= 0 {
				return p.rng(1, 4)
			}

			return p.rng(p.TotalPages-3, p.TotalPages)
		}

		return p.rng(1, p.TotalPages)
	}

	return nil
}

func (p paging) getTotalPages() int64 {
	if p.TotalResults > 0 {
		if p.TotalResults <= p.PageSize {
			return 1
		}

		if (p.TotalResults % p.PageSize) > 0 {
			return (p.TotalResults / p.PageSize) + 1
		}

		return (p.TotalResults / p.PageSize)
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

func (p paging) getLinks(f func(int64) string) map[string]pageURL {
	np := p.CurrentPage + 1
	pp := p.CurrentPage - 1

	if p.CurrentPage == 1 {
		pp = 1
	}

	if p.CurrentPage >= p.TotalPages {
		np = p.TotalPages
	}

	return map[string]pageURL{
		"previous": pageURL{
			Page: pp,
			Link: f(pp),
		},
		"current": pageURL{
			Page: p.CurrentPage,
			Link: f(p.CurrentPage),
		},
		"next": pageURL{
			Page: np,
			Link: f(np),
		},
	}
}

func (p paging) getPagesWithURL(f func(int64) string) []pageURL {
	pages := p.getPages()
	urlPages := make([]pageURL, len(pages))

	for i, v := range pages {
		urlPages[i] = pageURL{
			Page: v,
			Link: f(v),
		}
	}

	return urlPages
}
