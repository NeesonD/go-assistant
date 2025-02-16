package news

type NewsCache struct {
    articles map[string]Article
}

type Article struct {
    ID      string
    Title   string
    Content string
    Source  string
}

func NewNewsCache() *NewsCache {
    return &NewsCache{
        articles: make(map[string]Article),
    }
}

func (nc *NewsCache) AddArticle(article Article) {
    nc.articles[article.ID] = article
}

func (nc *NewsCache) GetArticle(id string) (Article, bool) {
    article, exists := nc.articles[id]
    return article, exists
}

func (nc *NewsCache) GetAllArticles() []Article {
    allArticles := make([]Article, 0, len(nc.articles))
    for _, article := range nc.articles {
        allArticles = append(allArticles, article)
    }
    return allArticles
}

func (nc *NewsCache) RemoveArticle(id string) {
    delete(nc.articles, id)
}