package node

import (
	"remez_story/infrastructure/errors"
)

type ChoiceBuilder struct {
	choice *Choice
	errors *errors.Errors
}

func NewChoiceBuilder() *ChoiceBuilder {
	return &ChoiceBuilder{
		choice: &Choice{},
		errors: errors.NewErrors(),
	}
}

func (b *ChoiceBuilder) ID(id ChoiceID) *ChoiceBuilder {
	b.choice.ID = id
	return b
}

func (b *ChoiceBuilder) Text(text string) *ChoiceBuilder {
	b.choice.Text = text
	return b
}

func (b *ChoiceBuilder) Effects(effects []Effect) *ChoiceBuilder {
	b.choice.Effects = effects
	return b
}

func (b *ChoiceBuilder) ToNodeID(toNodeID NodeID) *ChoiceBuilder {
	b.choice.ToNodeID = toNodeID
	return b
}

func (b *ChoiceBuilder) Build() (*Choice, error) {
	b.checkRequiredFields()
	if b.errors.IsPresent() {
		return nil, b.errors
	}

	return b.choice, nil
}

func (b *ChoiceBuilder) checkRequiredFields() {
	if b.choice.ID.IsZero() {
		b.errors.AddError(ErrChoiceIDRequired)
	}

	if b.choice.Text == "" {
		b.errors.AddError(ErrChoiceTextRequired)
	}

	if b.choice.ToNodeID.IsZero() {
		b.errors.AddError(ErrChoiceToNodeRequired)
	}
}
