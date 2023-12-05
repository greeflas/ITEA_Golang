package cmd

import (
	"context"
	"flag"
	"github.com/google/uuid"
	"github.com/greeflas/itea_golang/repository"
)

type UpdateArticleCommand struct {
	articleRepository *repository.ArticleRepository

	id    string
	title string
	body  string
}

func NewUpdateArticleCommand(articleRepository *repository.ArticleRepository) *UpdateArticleCommand {
	return &UpdateArticleCommand{articleRepository: articleRepository}
}

func (c *UpdateArticleCommand) Name() string {
	return "update_article"
}

func (c *UpdateArticleCommand) InitFlags(flags *flag.FlagSet) {
	flags.StringVar(&c.id, "id", "", "ID of the article")
	flags.StringVar(&c.title, "title", "", "title of the article")
	flags.StringVar(&c.body, "body", "", "body of the article")
}

func (c *UpdateArticleCommand) Run(ctx context.Context) error {
	article, err := c.articleRepository.FindById(ctx, uuid.MustParse(c.id))
	if err != nil {
		return err
	}

	if article == nil {
		return nil
	}

	if c.title != "" {
		article.ChangeTitle(c.title)
	}

	if c.body != "" {
		article.SetBody(c.body)
	}

	return c.articleRepository.Update(ctx, article)
}
