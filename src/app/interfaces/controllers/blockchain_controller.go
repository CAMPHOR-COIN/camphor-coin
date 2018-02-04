package controllers

import (
	"app/domain"
)

type BlockchainController struct {}

func NewBlockchainController() *BlockchainController {
	return &BlockchainController{}
}

func (controller *BlockchainController) Show(c Context) {
	// TODO DBから取得した最新のチェーンを利用する - 20180204 andoshin11
	blockchain := domain.GetBlockchain()

	// blockchain, err := controller.interactor.GetBlockchain()
	// if err != nil {
	//   c.JSON(500, NewError(err))
	//   return
	// }

	c.JSON(200, blockchain)
}
