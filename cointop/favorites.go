package cointop

func (ct *Cointop) toggleFavorite() error {
	ct.portfoliovisible = false
	coin := ct.highlightedRowCoin()
	if coin == nil {
		return nil
	}
	_, ok := ct.favorites[coin.Name]
	if ok {
		delete(ct.favorites, coin.Name)
		coin.Favorite = false
	} else {
		ct.favorites[coin.Name] = true
		coin.Favorite = true
	}
	go ct.updateTable()
	return nil
}

func (ct *Cointop) toggleShowFavorites() error {
	ct.portfoliovisible = false
	ct.filterByFavorites = !ct.filterByFavorites
	go ct.updateTable()
	return nil
}
