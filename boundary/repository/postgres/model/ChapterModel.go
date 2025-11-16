package postgresModel

import (
	"database/sql"
	"remez_story/common/domainPrimitive/primitive/description"
	"remez_story/common/domainPrimitive/primitive/title"
	"remez_story/domain/entity/chapter"
	"remez_story/infrastructure/errors"
)

type ChapterModel struct {
	ID          int64          `db:"id"`
	Title       string         `db:"title"`
	Description sql.NullString `db:"description"`
	OrderIndex  int            `db:"order_index"`
}

func ChapterFromModel(model *ChapterModel) (*chapter.Chapter, error) {
	errs := errors.NewErrors()

	chapterID, err := chapter.NewChapterID(model.ID)
	if err != nil {
		errs.AddError(err)
	}

	chapterTitle, err := title.NewTitle(model.Title)
	if err != nil {
		errs.AddError(err)
	}

	if errs.IsPresent() {
		return nil, errs
	}

	builder := chapter.NewChapterBuilder().
		ID(chapterID).
		TitleValue(chapterTitle).
		OrderIndex(model.OrderIndex)

	if model.Description.Valid {
		desc, err := description.NewDescription(model.Description.String)
		if err != nil {
			return nil, err
		}
		builder.DescriptionValue(desc)
	}

	return builder.Build()
}
