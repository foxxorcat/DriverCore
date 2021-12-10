package bilibili

// 检查链接是否有效
func (b *BiLiBiLi) CheckUrl(metaurl string) bool {
	res, err := b.client.R().SetContext(b.ctx).Head(b.formatUrl(metaurl))
	if err != nil || res.StatusCode() != 200 {
		return false
	}
	return true
}
