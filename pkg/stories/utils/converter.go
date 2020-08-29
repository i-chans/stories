package utils

import (
	"stories/pkg/stories/contract"
	"stories/pkg/stories/domain"
)

func DomainToContract(st domain.Story) contract.Story {
	return contract.Story{
		ID:    st.GetID(),
		Title: st.GetTitle(),
		Body:  st.GetBody(),
	}
}
