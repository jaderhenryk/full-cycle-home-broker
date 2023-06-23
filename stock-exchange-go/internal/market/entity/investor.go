package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetsPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetsPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetsPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPosition(assetID string, qtyShares int) {
	assetPosition := i.GetAssetPosition(assetID)
	if assetPosition == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestorAssetPosition(assetID, qtyShares))
	} else {
		assetPosition.Shares += qtyShares
	}
}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetsPosition {
	for _, assetPosition := range i.AssetPosition {
		if assetPosition.AssetID == assetID {
			return assetPosition
		}
	}
	return nil
}

type InvestorAssetsPosition struct {
	AssetID string
	Shares  int
}

func NewInvestorAssetPosition(assetID string, shares int) *InvestorAssetsPosition {
	return &InvestorAssetsPosition{
		AssetID: assetID,
		Shares:  shares,
	}
}
