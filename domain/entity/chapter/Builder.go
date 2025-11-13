package chapter

import (
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
	"remez_story/infrastructure/errors"
)

type ChapterBuilder struct {
	chapter *Chapter
	errors  *errors.Errors
}

func NewChapterBuilder() *ChapterBuilder {
	return &ChapterBuilder{
		chapter: &Chapter{},
		errors:  errors.NewErrors(),
	}
}

func (b *ChapterBuilder) ID(id ChapterID) *ChapterBuilder {
	b.chapter.ID = id
	return b
}

func (b *ChapterBuilder) Title(titleText string) *ChapterBuilder {
	t, err := title.NewTitle(titleText)
	if err != nil {
		b.errors.AddError(err)
		return b
	}
	b.chapter.Title = t
	return b
}

func (b *ChapterBuilder) TitleValue(t title.Title) *ChapterBuilder {
	b.chapter.Title = t
	return b
}

func (b *ChapterBuilder) Description(descText string) *ChapterBuilder {
	d, err := description.NewDescription(descText)
	if err != nil {
		b.errors.AddError(err)
		return b
	}
	b.chapter.Description = d
	return b
}

func (b *ChapterBuilder) DescriptionValue(d description.Description) *ChapterBuilder {
	b.chapter.Description = d
	return b
}

func (b *ChapterBuilder) Build() (*Chapter, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.chapter, nil
}

func (b *ChapterBuilder) checkRequiredFields() {
	if b.chapter.ID.IsZero() {
		b.errors.AddError(ErrChapterIDRequired)
	}

	if b.chapter.Title.IsZero() {
		b.errors.AddError(ErrTitleRequired)
	}
}
