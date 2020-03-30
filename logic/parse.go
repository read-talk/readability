package readability

func (d *Document) Parse() {
	d.removeScript()
	d.removeStyle()
	d.removeLink()
	d.Title = d.getTitle()
	//d.Content = d.getArticle()
}

func (d *Document) removeScript() {
	d.htmlDoc.Find("script").Remove()
}

func (d *Document) removeStyle() {
	d.htmlDoc.Find("style").Remove()
}

func (d *Document) removeLink() {
	d.htmlDoc.Find("link").Remove()
}

func (d *Document) getTitle() string {
	return d.htmlDoc.Find("title").Text()
}
