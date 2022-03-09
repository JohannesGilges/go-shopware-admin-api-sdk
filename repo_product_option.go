package go_shopware_admin_sdk

import (
	"net/http"
)

type ProductOptionRepository ClientService

func (t ProductOptionRepository) Search(ctx ApiContext, criteria Criteria) (*ProductOptionCollection, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search/product-option", criteria)

	if err != nil {
		return nil, nil, err
	}

	uResp := new(ProductOptionCollection)
	resp, err := t.Client.Do(ctx.Context, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

func (t ProductOptionRepository) SearchIds(ctx ApiContext, criteria Criteria) (*SearchIdsResponse, *http.Response, error) {
	req, err := t.Client.NewRequest(ctx, "POST", "/api/search-ids/product-option", criteria)

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

func (t ProductOptionRepository) Upsert(ctx ApiContext, entity []ProductOption) (*http.Response, error) {
	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_option": {
		Entity:  "product_option",
		Action:  "upsert",
		Payload: entity,
	}})
}

func (t ProductOptionRepository) Delete(ctx ApiContext, ids []string) (*http.Response, error) {
	payload := make([]entityDelete, 0)

	for _, id := range ids {
		payload = append(payload, entityDelete{Id: id})
	}

	return t.Client.Bulk.Sync(ctx, map[string]SyncOperation{"product_option": {
		Entity:  "product_option",
		Action:  "delete",
		Payload: payload,
	}})
}

type ProductOption struct {
	Option *PropertyGroupOption `json:"option,omitempty"`

	ProductId string `json:"productId,omitempty"`

	ProductVersionId string `json:"productVersionId,omitempty"`

	OptionId string `json:"optionId,omitempty"`

	Product *Product `json:"product,omitempty"`
}

type ProductOptionCollection struct {
	EntityCollection

	Data []ProductOption `json:"data"`
}
