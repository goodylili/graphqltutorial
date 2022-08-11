package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"graphqltutorial/graph/generated"
	"graphqltutorial/graph/model"
)

// CreateBio is the resolver for the CreateBio field.
func (r *mutationResolver) CreateBio(ctx context.Context, input model.AddBio) (*model.BioData, error) {

	person, err := r.BioRepository.CreatePerson()
	person_ := &model.BioData{
		ID:   person.ID,
		Name: person.Name,
		Age:  person.Age,
	}
	if err != nil {
		return nil, err
	}
	return person_, err
}

// DeleteBio is the resolver for the DeleteBio field.
func (r *mutationResolver) DeleteBio(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdateBio is the resolver for the UpdateBio field.
func (r *mutationResolver) UpdateBio(ctx context.Context, id int) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetBio is the resolver for the GetBio field.
func (r *queryResolver) GetBio(ctx context.Context, id int) (*model.BioData, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetBios is the resolver for the GetBios field.
func (r *queryResolver) GetBios(ctx context.Context) ([]*model.BioData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

//func (r *mutationResolver) CreateBio(ctx context.Context, input model.BioData) (*model.BioData, error) {
//	person, err := r.BioRepository.CreatePerson(input)
//	person_ := &model.BioData{
//		ID:   person.ID,
//		Name: person.Name,
//		Age:  person.Age,
//	}
//	if err != nil {
//		return nil, err
//	}
//	return person_, err
//
//}
//
//// Bios is the resolver for the todos field.
//func (r *queryResolver) Bios(ctx context.Context) ([]model.BioData, error) {
//	bio, err := r.BioRepository.GetAllPersons()
//	if err != nil {
//		return nil, err
//	}
//	return bio, nil
//}
//
//func (r *mutationResolver) UpdateBook(ctx context.Context, id int, input model.BioData) (string, error) {
//	err := r.BioRepository.UpdatePerson(input, id)
//	if err != nil {
//		return "nil", err
//	}
//	message := "successfully updated row"
//
//	return message, nil
//}
//
//func (r *mutationResolver) DeleteBook(ctx context.Context, id int) (string, error) {
//	err := r.BioRepository.DeletePerson(id)
//	if err != nil {
//		return "", err
//	}
//	message := "successfully deleted row"
//	return message, nil
//}
