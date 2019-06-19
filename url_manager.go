package main

type UrlManager struct {
	newUrls []string
	oldUrls []string
}

func NewUrlManager() *UrlManager {
	var result = new(UrlManager)
	result.newUrls = make([]string, 0)
	result.oldUrls = make([]string, 0)
	return result
}

func (c UrlManager) addNewUrl(url string) {

	if url == "" {
		return
	}

	p := func(urls []string) {
		for _, v := range urls {
			if v == url {
				return
			}
		}
	}

	p(c.oldUrls)
	p(c.newUrls)
	c.newUrls = append(c.newUrls, url)

}

func (c UrlManager) addNewUrls(urls []string) {
	if len(urls) == 0 {
		return
	}
	for _, url := range urls {
		c.addNewUrl(url)
	}
}

// 判断是否有newUrls中有URL条目
func (c UrlManager) hasNewUrl() bool {
	return len(c.newUrls) != 0
}

// 从URL条目拿取一条url，并从newUrls移除，在oldUrls添加
func (c UrlManager) getNewUrl() (url string) {
	const index = 0
	url = c.newUrls[index]
	c.newUrls = append(c.newUrls[:index], c.newUrls[index+1:]...)
	c.oldUrls = append(c.oldUrls, url)
	return url
}
