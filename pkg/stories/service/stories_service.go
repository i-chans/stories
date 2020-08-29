package service

import (
	"go.uber.org/zap"
	"stories/cmd/config"
	"stories/pkg/stories/contract"
	"stories/pkg/stories/domain"
	"stories/pkg/stories/store"
	"stories/pkg/stories/utils"
)

type StoriesService interface {
	Add(addReq contract.AddStoryRequest) (contract.AddStoryResponse, error)
	Search(searchReq contract.SearchStoryRequest) (contract.SearchStoryResponse, error)
	Update(updateReq contract.UpdateStoryRequest) (contract.UpdateStoryResponse, error)
	Delete(deleteReq contract.DeleteStoryRequest) (contract.DeleteStoryResponse, error)
}

type defaultStoriesService struct {
	cfg config.StoryConfig
	lgr *zap.Logger
	str *store.Store
}

func (dss *defaultStoriesService) Add(addReq contract.AddStoryRequest) (contract.AddStoryResponse, error) {
	story, err := domain.NewVanillaStory(dss.cfg.GetTitleMaxLength(), dss.cfg.GetBodyMaxLength(), addReq.Title, addReq.Body)
	if err != nil {
		dss.lgr.Error(err.Error())
		return contract.AddStoryResponse{}, err
	}

	id, err := dss.str.StoriesStore().Add(story)
	if err != nil {
		return contract.AddStoryResponse{}, err
	}

	return contract.AddStoryResponse{StoryID: id}, nil
}

func (dss *defaultStoriesService) Search(searchReq contract.SearchStoryRequest) (contract.SearchStoryResponse, error) {
	stories, err := dss.str.StoriesStore().Search(searchReq.Query)
	if err != nil {
		return contract.SearchStoryResponse{}, err
	}

	resp := contract.SearchStoryResponse{}

	for _, story := range stories {
		resp.Stories = append(resp.Stories, utils.DomainToContract(story))
	}

	return resp, nil
}

func (dss *defaultStoriesService) Update(updateReq contract.UpdateStoryRequest) (contract.UpdateStoryResponse, error) {
	story, err := domain.NewStory(dss.cfg.GetTitleMaxLength(), dss.cfg.GetBodyMaxLength(), updateReq.ID, updateReq.Title, updateReq.Body)
	if err != nil {
		dss.lgr.Error(err.Error())
		return contract.UpdateStoryResponse{}, err
	}

	_, err = dss.str.StoriesStore().Update(story)
	if err != nil {
		return contract.UpdateStoryResponse{}, err
	}

	return contract.UpdateStoryResponse{Success: true}, nil
}

func (dss *defaultStoriesService) Delete(deleteReq contract.DeleteStoryRequest) (contract.DeleteStoryResponse, error) {
	_, err := dss.str.StoriesStore().Delete(deleteReq.ID)
	if err != nil {
		return contract.DeleteStoryResponse{}, err
	}

	return contract.DeleteStoryResponse{Success: true}, nil
}

func NewStoriesService(cfg config.StoryConfig, lgr *zap.Logger, str *store.Store) StoriesService {
	return &defaultStoriesService{
		cfg: cfg,
		lgr: lgr,
		str: str,
	}
}
