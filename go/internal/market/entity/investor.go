package entity

type Investor struct {
	Id            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

type InvestorAssetPosition struct {
	AssetId string
	Shares  int
}

func NewInvestor(id string) *Investor {
	return &Investor{
		Id:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (investor *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	investor.AssetPosition = append(investor.AssetPosition, assetPosition)
}

func (investor *Investor) UpdateAssetPosition(assetId string, qtdShares int) {
	assetPosition := investor.GetAssetPosition(assetId)

	if assetPosition == nil {
		investor.AssetPosition = append(investor.AssetPosition, NewInvestorAssetPosition(assetId, qtdShares))
	} else {
		assetPosition.Shares += qtdShares
	}
}

func (investor *Investor) GetAssetPosition(assetId string) *InvestorAssetPosition {
	for _, assetPosition := range investor.AssetPosition {
		if assetPosition.AssetId == assetId {
			return assetPosition
		}
	}

	return nil
}

func NewInvestorAssetPosition(assetId string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		AssetId: assetId,
		Shares:  shares,
	}
}
