package query

import (
	"errors"

	"ecommerce-microservice/membership/src/model"
)

// memberQueryInMemory model
type memberQueryInMemory struct {
	db map[string]*model.Member
}

// NewMemberQueryInMemory for initialise memberQueryInMemory model
func NewMemberQueryInMemory(db map[string]*model.Member) MembershipQuery {
	return &memberQueryInMemory{db}
}

// FindByEmail will return Member by its email
func (q *memberQueryInMemory) FindByEmail(email string) <-chan QueryResult {
	output := make(chan QueryResult)
	go func() {
		defer close(output)

		var member *model.Member
		for _, v := range q.db {
			if v.Email == email {
				member = v
				break
			} else {
				output <- QueryResult{Error: errors.New("MEMBER_NOT_FOUND")}
				return
			}
		}

		output <- QueryResult{Result: member}
	}()
	return output
}
