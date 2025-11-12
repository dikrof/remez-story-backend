package chapter

import "remez_story/infrastructure/errors"

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

func (b *ChapterBuilder) Title(title string) *ChapterBuilder {
	b.chapter.Title = title
	return b
}

func (b *ChapterBuilder) Description(description string) *ChapterBuilder {
	b.chapter.Description = description
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

	if b.chapter.Title == "" {
		b.errors.AddError(ErrTitleRequired)
	}
}
