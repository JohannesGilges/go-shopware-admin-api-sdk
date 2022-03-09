package go_shopware_admin_sdk

import (
	"net/http"
	"time"
)

type MainCategoryRepository ClientService

func (t MainCategoryRepository) Search(ctx ApiContext, criteria Criteria) (*MainCategoryCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/main-category", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(MainCategoryCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MainCategoryRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/main-category", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(SearchIdsResponse)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t MainCategoryRepository) Upsert(ctx ApiContext, entity []MainCategory) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"main_category": {
		Entity:  "main_category",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t MainCategoryRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"main_category": {
		Entity:  "main_category",
		Action:  "delete",
		Payload: payload,
	}})
}

type MainCategory struct {
	CategoryId string `json:"categoryId,omitempty"`

	CategoryVersionId string `json:"categoryVersionId,omitempty"`

	Category *Category `json:"category,omitempty"`

	CreatedAt time.Time `json:"createdAt,omitempty"`

	UpdatedAt time.Time `json:"updatedAt,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	Product *Product `json:"product,omitempty"`

	SalesChannel *SalesChannel `json:"salesChannel,omitempty"`

	Id string `json:"id,omitempty"`

	SalesChannelId string `json:"salesChannelId,omitempty"`
}

type MainCategoryCollection struct {
	EntityCollection

	Data []MainCategory `json:"data"`
}
