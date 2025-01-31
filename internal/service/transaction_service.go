package service

import (
<<<<<<< HEAD
	"strconv"

=======
	"context"
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	r "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
<<<<<<< HEAD
=======
	"strconv"
	"time"
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
)

type TransactionService interface {
	GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error)
<<<<<<< HEAD
	TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error)
	Transfer(input *dto.TransferRequestBody) (*model.Transaction, error)
=======
	TopUp(ctx context.Context, input *dto.TopUpRequestBody) (*model.Transaction, error)
	Transfer(ctx context.Context, input *dto.TransferRequestBody) (*model.Transaction, error)
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	CountTransaction(userID int) (int64, error)
}

type transactionService struct {
	transactionRepository  r.TransactionRepository
	walletRepository       r.WalletRepository
	sourceOfFundRepository r.SourceOfFundRepository
}

type TSConfig struct {
	TransactionRepository  r.TransactionRepository
	WalletRepository       r.WalletRepository
	SourceOfFundRepository r.SourceOfFundRepository
}

func NewTransactionService(c *TSConfig) TransactionService {
	return &transactionService{
		transactionRepository:  c.TransactionRepository,
		walletRepository:       c.WalletRepository,
		sourceOfFundRepository: c.SourceOfFundRepository,
	}
}

func (s *transactionService) GetTransactions(userID int, query *dto.TransactionRequestQuery) ([]*model.Transaction, error) {
	transactions, err := s.transactionRepository.FindAll(userID, query)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

<<<<<<< HEAD
func (s *transactionService) TopUp(input *dto.TopUpRequestBody) (*model.Transaction, error) {
=======
func (s *transactionService) TopUp(ctx context.Context, input *dto.TopUpRequestBody) (*model.Transaction, error) {
	tx := s.transactionRepository.Begin(ctx)

>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	sourceOfFund, err := s.sourceOfFundRepository.FindById(input.SourceOfFundID)
	if err != nil {
		return &model.Transaction{}, err
	}
	if sourceOfFund.ID == 0 {
		return &model.Transaction{}, &custom_error.SourceOfFundNotFoundError{}
	}

	wallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if wallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	transaction := &model.Transaction{}
	transaction.SourceOfFundID = &sourceOfFund.ID
	transaction.UserID = input.User.ID
	transaction.DestinationID = wallet.ID
<<<<<<< HEAD
=======
	time.Sleep(5 * time.Second)
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	transaction.Amount = input.Amount
	transaction.Description = "Top Up from " + sourceOfFund.Name
	transaction.Category = "Top Up"

<<<<<<< HEAD
	transaction, err = s.transactionRepository.Save(transaction)
=======
	transaction, err = s.transactionRepository.Save(ctx, tx, transaction)
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	if err != nil {
		return transaction, err
	}

	wallet.Balance = wallet.Balance + input.Amount
	wallet, err = s.walletRepository.Update(wallet)
	if err != nil {
		return transaction, err
	}

	transaction.SourceOfFund = sourceOfFund
	transaction.User = *input.User
	transaction.Wallet = *wallet

<<<<<<< HEAD
=======
	tx.Commit()

>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	return transaction, nil
}

func (s *transactionService) CountTransaction(userID int) (int64, error) {
	totalTransactions, err := s.transactionRepository.Count(userID)
	if err != nil {
		return totalTransactions, err
	}

	return totalTransactions, nil
}

<<<<<<< HEAD
func (s *transactionService) Transfer(input *dto.TransferRequestBody) (*model.Transaction, error) {
=======
func (s *transactionService) Transfer(ctx context.Context, input *dto.TransferRequestBody) (*model.Transaction, error) {
	// Implement Atomicity for transaction
	tx := s.transactionRepository.Begin(ctx)

>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	myWallet, err := s.walletRepository.FindByUserId(int(input.User.ID))
	if err != nil {
		return &model.Transaction{}, err
	}
	if myWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}
	if myWallet.Balance < input.Amount {
		return &model.Transaction{}, &custom_error.InsufficientBallanceError{}
	}
	number := strconv.Itoa(input.WalletNumber)
	if myWallet.Number == number {
		return &model.Transaction{}, &custom_error.TransferToSameWalletError{}
	}

	destinationWallet, err := s.walletRepository.FindByNumber(number)
	if err != nil {
		return &model.Transaction{}, err
	}
	if destinationWallet.ID == 0 {
		return &model.Transaction{}, &custom_error.WalletNotFoundError{}
	}

	//create transaction for receiver
	transaction := &model.Transaction{}
	transaction.UserID = destinationWallet.User.ID
	transaction.DestinationID = myWallet.ID
	transaction.Amount = input.Amount
	transaction.Description = input.Description
	transaction.Category = "Receive Money"

<<<<<<< HEAD
	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

=======
	transaction, err = s.transactionRepository.Save(ctx, tx, transaction)

	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	time.Sleep(5 * time.Second)
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	// create transaction for sender
	transaction = &model.Transaction{}
	transaction.UserID = input.User.ID
	transaction.DestinationID = destinationWallet.ID
	transaction.Amount = input.Amount
	transaction.Description = input.Description
	transaction.Category = "Send Money"

<<<<<<< HEAD
	transaction, err = s.transactionRepository.Save(transaction)
	if err != nil {
=======
	transaction, err = s.transactionRepository.Save(ctx, tx, transaction)
	if err != nil {
		tx.Rollback()
>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
		return transaction, err
	}

	myWallet.Balance = myWallet.Balance - input.Amount
	myWallet, err = s.walletRepository.Update(myWallet)
	if err != nil {
		return transaction, err
	}

	destinationWallet.Balance = destinationWallet.Balance + input.Amount
	_, err = s.walletRepository.Update(destinationWallet)
	if err != nil {
		return transaction, err
	}

	balance := uint(myWallet.Balance)
	transaction.SourceOfFundID = &balance
	transaction.User = *input.User
	transaction.Wallet = *destinationWallet

<<<<<<< HEAD
=======
	tx.Commit()

>>>>>>> a0a5309 (fix(): add atomic transaction to the wallet)
	return transaction, nil
}
