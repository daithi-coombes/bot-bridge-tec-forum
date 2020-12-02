// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dao

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ConvictionBetaABI is the input ABI used to generate the binding from.
const ConvictionBetaABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"hasInitialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"D\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ABSTAIN_PROPOSAL_ID\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_script\",\"type\":\"bytes\"}],\"name\":\"getEVMScriptExecutor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRecoveryVault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_STAKED_PROPOSALS\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"onTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hookId\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"onRegisterAsHook\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"UPDATE_SETTINGS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxRatio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"allowRecoverability\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"appId\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CANCEL_PROPOSAL_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInitializationBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"transferToVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_sender\",\"type\":\"address\"},{\"name\":\"_role\",\"type\":\"bytes32\"},{\"name\":\"_params\",\"type\":\"uint256[]\"}],\"name\":\"canPerform\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weight\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEVMScriptRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_hookId\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"onRevokeAsHook\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kernel\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"onApprove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minThresholdStakePercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ONE_HUNDRED_PERCENT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPetrified\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CREATE_PROPOSALS_ROLE\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"decay\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"maxRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minThresholdStakePercentage\",\"type\":\"uint256\"}],\"name\":\"ConvictionSettingsChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"link\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"ProposalAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokensStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"conviction\",\"type\":\"uint256\"}],\"name\":\"StakeAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"entity\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokensStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalTokensStaked\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"conviction\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"conviction\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ProposalCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"script\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"input\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"ScriptResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RecoverToVault\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_stakeToken\",\"type\":\"address\"},{\"name\":\"_vault\",\"type\":\"address\"},{\"name\":\"_requestToken\",\"type\":\"address\"},{\"name\":\"_decay\",\"type\":\"uint256\"},{\"name\":\"_maxRatio\",\"type\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\"},{\"name\":\"_minThresholdStakePercentage\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_decay\",\"type\":\"uint256\"},{\"name\":\"_maxRatio\",\"type\":\"uint256\"},{\"name\":\"_weight\",\"type\":\"uint256\"},{\"name\":\"_minThresholdStakePercentage\",\"type\":\"uint256\"}],\"name\":\"setConvictionCalculationSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_title\",\"type\":\"string\"},{\"name\":\"_link\",\"type\":\"bytes\"},{\"name\":\"_requestedAmount\",\"type\":\"uint256\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"addProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stakeToProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"stakeAllToProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"withdrawAllFromProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"cancelProposal\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"name\":\"requestedAmount\",\"type\":\"uint256\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"stakedTokens\",\"type\":\"uint256\"},{\"name\":\"convictionLast\",\"type\":\"uint256\"},{\"name\":\"blockLast\",\"type\":\"uint64\"},{\"name\":\"proposalStatus\",\"type\":\"uint8\"},{\"name\":\"submitter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getProposalVoterStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getTotalVoterStake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterStakedProposals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_timePassed\",\"type\":\"uint64\"},{\"name\":\"_lastConv\",\"type\":\"uint256\"},{\"name\":\"_oldAmount\",\"type\":\"uint256\"}],\"name\":\"calculateConviction\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_requestedAmount\",\"type\":\"uint256\"}],\"name\":\"calculateThreshold\",\"outputs\":[{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ConvictionBeta is an auto generated Go binding around an Ethereum contract.
type ConvictionBeta struct {
	ConvictionBetaCaller     // Read-only binding to the contract
	ConvictionBetaTransactor // Write-only binding to the contract
	ConvictionBetaFilterer   // Log filterer for contract events
}

// ConvictionBetaCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConvictionBetaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvictionBetaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConvictionBetaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvictionBetaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConvictionBetaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConvictionBetaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConvictionBetaSession struct {
	Contract     *ConvictionBeta   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConvictionBetaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConvictionBetaCallerSession struct {
	Contract *ConvictionBetaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ConvictionBetaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConvictionBetaTransactorSession struct {
	Contract     *ConvictionBetaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ConvictionBetaRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConvictionBetaRaw struct {
	Contract *ConvictionBeta // Generic contract binding to access the raw methods on
}

// ConvictionBetaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConvictionBetaCallerRaw struct {
	Contract *ConvictionBetaCaller // Generic read-only contract binding to access the raw methods on
}

// ConvictionBetaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConvictionBetaTransactorRaw struct {
	Contract *ConvictionBetaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConvictionBeta creates a new instance of ConvictionBeta, bound to a specific deployed contract.
func NewConvictionBeta(address common.Address, backend bind.ContractBackend) (*ConvictionBeta, error) {
	contract, err := bindConvictionBeta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConvictionBeta{ConvictionBetaCaller: ConvictionBetaCaller{contract: contract}, ConvictionBetaTransactor: ConvictionBetaTransactor{contract: contract}, ConvictionBetaFilterer: ConvictionBetaFilterer{contract: contract}}, nil
}

// NewConvictionBetaCaller creates a new read-only instance of ConvictionBeta, bound to a specific deployed contract.
func NewConvictionBetaCaller(address common.Address, caller bind.ContractCaller) (*ConvictionBetaCaller, error) {
	contract, err := bindConvictionBeta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaCaller{contract: contract}, nil
}

// NewConvictionBetaTransactor creates a new write-only instance of ConvictionBeta, bound to a specific deployed contract.
func NewConvictionBetaTransactor(address common.Address, transactor bind.ContractTransactor) (*ConvictionBetaTransactor, error) {
	contract, err := bindConvictionBeta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaTransactor{contract: contract}, nil
}

// NewConvictionBetaFilterer creates a new log filterer instance of ConvictionBeta, bound to a specific deployed contract.
func NewConvictionBetaFilterer(address common.Address, filterer bind.ContractFilterer) (*ConvictionBetaFilterer, error) {
	contract, err := bindConvictionBeta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaFilterer{contract: contract}, nil
}

// bindConvictionBeta binds a generic wrapper to an already deployed contract.
func bindConvictionBeta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConvictionBetaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvictionBeta *ConvictionBetaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvictionBeta.Contract.ConvictionBetaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvictionBeta *ConvictionBetaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.ConvictionBetaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvictionBeta *ConvictionBetaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.ConvictionBetaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConvictionBeta *ConvictionBetaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConvictionBeta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConvictionBeta *ConvictionBetaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConvictionBeta *ConvictionBetaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.contract.Transact(opts, method, params...)
}

// ABSTAINPROPOSALID is a free data retrieval call binding the contract method 0x1522eb43.
//
// Solidity: function ABSTAIN_PROPOSAL_ID() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) ABSTAINPROPOSALID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "ABSTAIN_PROPOSAL_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ABSTAINPROPOSALID is a free data retrieval call binding the contract method 0x1522eb43.
//
// Solidity: function ABSTAIN_PROPOSAL_ID() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) ABSTAINPROPOSALID() (*big.Int, error) {
	return _ConvictionBeta.Contract.ABSTAINPROPOSALID(&_ConvictionBeta.CallOpts)
}

// ABSTAINPROPOSALID is a free data retrieval call binding the contract method 0x1522eb43.
//
// Solidity: function ABSTAIN_PROPOSAL_ID() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) ABSTAINPROPOSALID() (*big.Int, error) {
	return _ConvictionBeta.Contract.ABSTAINPROPOSALID(&_ConvictionBeta.CallOpts)
}

// CANCELPROPOSALROLE is a free data retrieval call binding the contract method 0x8af07b10.
//
// Solidity: function CANCEL_PROPOSAL_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCaller) CANCELPROPOSALROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "CANCEL_PROPOSAL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELPROPOSALROLE is a free data retrieval call binding the contract method 0x8af07b10.
//
// Solidity: function CANCEL_PROPOSAL_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaSession) CANCELPROPOSALROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.CANCELPROPOSALROLE(&_ConvictionBeta.CallOpts)
}

// CANCELPROPOSALROLE is a free data retrieval call binding the contract method 0x8af07b10.
//
// Solidity: function CANCEL_PROPOSAL_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCallerSession) CANCELPROPOSALROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.CANCELPROPOSALROLE(&_ConvictionBeta.CallOpts)
}

// CREATEPROPOSALSROLE is a free data retrieval call binding the contract method 0xeeabf67e.
//
// Solidity: function CREATE_PROPOSALS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCaller) CREATEPROPOSALSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "CREATE_PROPOSALS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATEPROPOSALSROLE is a free data retrieval call binding the contract method 0xeeabf67e.
//
// Solidity: function CREATE_PROPOSALS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaSession) CREATEPROPOSALSROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.CREATEPROPOSALSROLE(&_ConvictionBeta.CallOpts)
}

// CREATEPROPOSALSROLE is a free data retrieval call binding the contract method 0xeeabf67e.
//
// Solidity: function CREATE_PROPOSALS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCallerSession) CREATEPROPOSALSROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.CREATEPROPOSALSROLE(&_ConvictionBeta.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) D(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "D")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) D() (*big.Int, error) {
	return _ConvictionBeta.Contract.D(&_ConvictionBeta.CallOpts)
}

// D is a free data retrieval call binding the contract method 0x0f529ba2.
//
// Solidity: function D() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) D() (*big.Int, error) {
	return _ConvictionBeta.Contract.D(&_ConvictionBeta.CallOpts)
}

// MAXSTAKEDPROPOSALS is a free data retrieval call binding the contract method 0x406244d8.
//
// Solidity: function MAX_STAKED_PROPOSALS() view returns(uint64)
func (_ConvictionBeta *ConvictionBetaCaller) MAXSTAKEDPROPOSALS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "MAX_STAKED_PROPOSALS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXSTAKEDPROPOSALS is a free data retrieval call binding the contract method 0x406244d8.
//
// Solidity: function MAX_STAKED_PROPOSALS() view returns(uint64)
func (_ConvictionBeta *ConvictionBetaSession) MAXSTAKEDPROPOSALS() (uint64, error) {
	return _ConvictionBeta.Contract.MAXSTAKEDPROPOSALS(&_ConvictionBeta.CallOpts)
}

// MAXSTAKEDPROPOSALS is a free data retrieval call binding the contract method 0x406244d8.
//
// Solidity: function MAX_STAKED_PROPOSALS() view returns(uint64)
func (_ConvictionBeta *ConvictionBetaCallerSession) MAXSTAKEDPROPOSALS() (uint64, error) {
	return _ConvictionBeta.Contract.MAXSTAKEDPROPOSALS(&_ConvictionBeta.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) ONEHUNDREDPERCENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "ONE_HUNDRED_PERCENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _ConvictionBeta.Contract.ONEHUNDREDPERCENT(&_ConvictionBeta.CallOpts)
}

// ONEHUNDREDPERCENT is a free data retrieval call binding the contract method 0xdd0081c7.
//
// Solidity: function ONE_HUNDRED_PERCENT() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) ONEHUNDREDPERCENT() (*big.Int, error) {
	return _ConvictionBeta.Contract.ONEHUNDREDPERCENT(&_ConvictionBeta.CallOpts)
}

// UPDATESETTINGSROLE is a free data retrieval call binding the contract method 0x4f629fbb.
//
// Solidity: function UPDATE_SETTINGS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCaller) UPDATESETTINGSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "UPDATE_SETTINGS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPDATESETTINGSROLE is a free data retrieval call binding the contract method 0x4f629fbb.
//
// Solidity: function UPDATE_SETTINGS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaSession) UPDATESETTINGSROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.UPDATESETTINGSROLE(&_ConvictionBeta.CallOpts)
}

// UPDATESETTINGSROLE is a free data retrieval call binding the contract method 0x4f629fbb.
//
// Solidity: function UPDATE_SETTINGS_ROLE() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCallerSession) UPDATESETTINGSROLE() ([32]byte, error) {
	return _ConvictionBeta.Contract.UPDATESETTINGSROLE(&_ConvictionBeta.CallOpts)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ConvictionBeta *ConvictionBetaCaller) AllowRecoverability(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "allowRecoverability", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) AllowRecoverability(token common.Address) (bool, error) {
	return _ConvictionBeta.Contract.AllowRecoverability(&_ConvictionBeta.CallOpts, token)
}

// AllowRecoverability is a free data retrieval call binding the contract method 0x7e7db6e1.
//
// Solidity: function allowRecoverability(address token) view returns(bool)
func (_ConvictionBeta *ConvictionBetaCallerSession) AllowRecoverability(token common.Address) (bool, error) {
	return _ConvictionBeta.Contract.AllowRecoverability(&_ConvictionBeta.CallOpts, token)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCaller) AppId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "appId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaSession) AppId() ([32]byte, error) {
	return _ConvictionBeta.Contract.AppId(&_ConvictionBeta.CallOpts)
}

// AppId is a free data retrieval call binding the contract method 0x80afdea8.
//
// Solidity: function appId() view returns(bytes32)
func (_ConvictionBeta *ConvictionBetaCallerSession) AppId() ([32]byte, error) {
	return _ConvictionBeta.Contract.AppId(&_ConvictionBeta.CallOpts)
}

// CalculateConviction is a free data retrieval call binding the contract method 0xdac61260.
//
// Solidity: function calculateConviction(uint64 _timePassed, uint256 _lastConv, uint256 _oldAmount) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) CalculateConviction(opts *bind.CallOpts, _timePassed uint64, _lastConv *big.Int, _oldAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "calculateConviction", _timePassed, _lastConv, _oldAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateConviction is a free data retrieval call binding the contract method 0xdac61260.
//
// Solidity: function calculateConviction(uint64 _timePassed, uint256 _lastConv, uint256 _oldAmount) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) CalculateConviction(_timePassed uint64, _lastConv *big.Int, _oldAmount *big.Int) (*big.Int, error) {
	return _ConvictionBeta.Contract.CalculateConviction(&_ConvictionBeta.CallOpts, _timePassed, _lastConv, _oldAmount)
}

// CalculateConviction is a free data retrieval call binding the contract method 0xdac61260.
//
// Solidity: function calculateConviction(uint64 _timePassed, uint256 _lastConv, uint256 _oldAmount) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) CalculateConviction(_timePassed uint64, _lastConv *big.Int, _oldAmount *big.Int) (*big.Int, error) {
	return _ConvictionBeta.Contract.CalculateConviction(&_ConvictionBeta.CallOpts, _timePassed, _lastConv, _oldAmount)
}

// CalculateThreshold is a free data retrieval call binding the contract method 0x59a5db8b.
//
// Solidity: function calculateThreshold(uint256 _requestedAmount) view returns(uint256 _threshold)
func (_ConvictionBeta *ConvictionBetaCaller) CalculateThreshold(opts *bind.CallOpts, _requestedAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "calculateThreshold", _requestedAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateThreshold is a free data retrieval call binding the contract method 0x59a5db8b.
//
// Solidity: function calculateThreshold(uint256 _requestedAmount) view returns(uint256 _threshold)
func (_ConvictionBeta *ConvictionBetaSession) CalculateThreshold(_requestedAmount *big.Int) (*big.Int, error) {
	return _ConvictionBeta.Contract.CalculateThreshold(&_ConvictionBeta.CallOpts, _requestedAmount)
}

// CalculateThreshold is a free data retrieval call binding the contract method 0x59a5db8b.
//
// Solidity: function calculateThreshold(uint256 _requestedAmount) view returns(uint256 _threshold)
func (_ConvictionBeta *ConvictionBetaCallerSession) CalculateThreshold(_requestedAmount *big.Int) (*big.Int, error) {
	return _ConvictionBeta.Contract.CalculateThreshold(&_ConvictionBeta.CallOpts, _requestedAmount)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ConvictionBeta *ConvictionBetaCaller) CanPerform(opts *bind.CallOpts, _sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "canPerform", _sender, _role, _params)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ConvictionBeta.Contract.CanPerform(&_ConvictionBeta.CallOpts, _sender, _role, _params)
}

// CanPerform is a free data retrieval call binding the contract method 0xa1658fad.
//
// Solidity: function canPerform(address _sender, bytes32 _role, uint256[] _params) view returns(bool)
func (_ConvictionBeta *ConvictionBetaCallerSession) CanPerform(_sender common.Address, _role [32]byte, _params []*big.Int) (bool, error) {
	return _ConvictionBeta.Contract.CanPerform(&_ConvictionBeta.CallOpts, _sender, _role, _params)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) Decay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "decay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) Decay() (*big.Int, error) {
	return _ConvictionBeta.Contract.Decay(&_ConvictionBeta.CallOpts)
}

// Decay is a free data retrieval call binding the contract method 0x54fac919.
//
// Solidity: function decay() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) Decay() (*big.Int, error) {
	return _ConvictionBeta.Contract.Decay(&_ConvictionBeta.CallOpts)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) GetEVMScriptExecutor(opts *bind.CallOpts, _script []byte) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getEVMScriptExecutor", _script)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _ConvictionBeta.Contract.GetEVMScriptExecutor(&_ConvictionBeta.CallOpts, _script)
}

// GetEVMScriptExecutor is a free data retrieval call binding the contract method 0x2914b9bd.
//
// Solidity: function getEVMScriptExecutor(bytes _script) view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetEVMScriptExecutor(_script []byte) (common.Address, error) {
	return _ConvictionBeta.Contract.GetEVMScriptExecutor(&_ConvictionBeta.CallOpts, _script)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) GetEVMScriptRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getEVMScriptRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) GetEVMScriptRegistry() (common.Address, error) {
	return _ConvictionBeta.Contract.GetEVMScriptRegistry(&_ConvictionBeta.CallOpts)
}

// GetEVMScriptRegistry is a free data retrieval call binding the contract method 0xa479e508.
//
// Solidity: function getEVMScriptRegistry() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetEVMScriptRegistry() (common.Address, error) {
	return _ConvictionBeta.Contract.GetEVMScriptRegistry(&_ConvictionBeta.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) GetInitializationBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getInitializationBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) GetInitializationBlock() (*big.Int, error) {
	return _ConvictionBeta.Contract.GetInitializationBlock(&_ConvictionBeta.CallOpts)
}

// GetInitializationBlock is a free data retrieval call binding the contract method 0x8b3dd749.
//
// Solidity: function getInitializationBlock() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetInitializationBlock() (*big.Int, error) {
	return _ConvictionBeta.Contract.GetInitializationBlock(&_ConvictionBeta.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(uint256 requestedAmount, address beneficiary, uint256 stakedTokens, uint256 convictionLast, uint64 blockLast, uint8 proposalStatus, address submitter)
func (_ConvictionBeta *ConvictionBetaCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (struct {
	RequestedAmount *big.Int
	Beneficiary     common.Address
	StakedTokens    *big.Int
	ConvictionLast  *big.Int
	BlockLast       uint64
	ProposalStatus  uint8
	Submitter       common.Address
}, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getProposal", _proposalId)

	outstruct := new(struct {
		RequestedAmount *big.Int
		Beneficiary     common.Address
		StakedTokens    *big.Int
		ConvictionLast  *big.Int
		BlockLast       uint64
		ProposalStatus  uint8
		Submitter       common.Address
	})

	outstruct.RequestedAmount = out[0].(*big.Int)
	outstruct.Beneficiary = out[1].(common.Address)
	outstruct.StakedTokens = out[2].(*big.Int)
	outstruct.ConvictionLast = out[3].(*big.Int)
	outstruct.BlockLast = out[4].(uint64)
	outstruct.ProposalStatus = out[5].(uint8)
	outstruct.Submitter = out[6].(common.Address)

	return *outstruct, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(uint256 requestedAmount, address beneficiary, uint256 stakedTokens, uint256 convictionLast, uint64 blockLast, uint8 proposalStatus, address submitter)
func (_ConvictionBeta *ConvictionBetaSession) GetProposal(_proposalId *big.Int) (struct {
	RequestedAmount *big.Int
	Beneficiary     common.Address
	StakedTokens    *big.Int
	ConvictionLast  *big.Int
	BlockLast       uint64
	ProposalStatus  uint8
	Submitter       common.Address
}, error) {
	return _ConvictionBeta.Contract.GetProposal(&_ConvictionBeta.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(uint256 requestedAmount, address beneficiary, uint256 stakedTokens, uint256 convictionLast, uint64 blockLast, uint8 proposalStatus, address submitter)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetProposal(_proposalId *big.Int) (struct {
	RequestedAmount *big.Int
	Beneficiary     common.Address
	StakedTokens    *big.Int
	ConvictionLast  *big.Int
	BlockLast       uint64
	ProposalStatus  uint8
	Submitter       common.Address
}, error) {
	return _ConvictionBeta.Contract.GetProposal(&_ConvictionBeta.CallOpts, _proposalId)
}

// GetProposalVoterStake is a free data retrieval call binding the contract method 0xe0dd2c38.
//
// Solidity: function getProposalVoterStake(uint256 _proposalId, address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) GetProposalVoterStake(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getProposalVoterStake", _proposalId, _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalVoterStake is a free data retrieval call binding the contract method 0xe0dd2c38.
//
// Solidity: function getProposalVoterStake(uint256 _proposalId, address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) GetProposalVoterStake(_proposalId *big.Int, _voter common.Address) (*big.Int, error) {
	return _ConvictionBeta.Contract.GetProposalVoterStake(&_ConvictionBeta.CallOpts, _proposalId, _voter)
}

// GetProposalVoterStake is a free data retrieval call binding the contract method 0xe0dd2c38.
//
// Solidity: function getProposalVoterStake(uint256 _proposalId, address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetProposalVoterStake(_proposalId *big.Int, _voter common.Address) (*big.Int, error) {
	return _ConvictionBeta.Contract.GetProposalVoterStake(&_ConvictionBeta.CallOpts, _proposalId, _voter)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) GetRecoveryVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getRecoveryVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) GetRecoveryVault() (common.Address, error) {
	return _ConvictionBeta.Contract.GetRecoveryVault(&_ConvictionBeta.CallOpts)
}

// GetRecoveryVault is a free data retrieval call binding the contract method 0x32f0a3b5.
//
// Solidity: function getRecoveryVault() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetRecoveryVault() (common.Address, error) {
	return _ConvictionBeta.Contract.GetRecoveryVault(&_ConvictionBeta.CallOpts)
}

// GetTotalVoterStake is a free data retrieval call binding the contract method 0xddc90e7e.
//
// Solidity: function getTotalVoterStake(address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) GetTotalVoterStake(opts *bind.CallOpts, _voter common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getTotalVoterStake", _voter)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVoterStake is a free data retrieval call binding the contract method 0xddc90e7e.
//
// Solidity: function getTotalVoterStake(address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) GetTotalVoterStake(_voter common.Address) (*big.Int, error) {
	return _ConvictionBeta.Contract.GetTotalVoterStake(&_ConvictionBeta.CallOpts, _voter)
}

// GetTotalVoterStake is a free data retrieval call binding the contract method 0xddc90e7e.
//
// Solidity: function getTotalVoterStake(address _voter) view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) GetTotalVoterStake(_voter common.Address) (*big.Int, error) {
	return _ConvictionBeta.Contract.GetTotalVoterStake(&_ConvictionBeta.CallOpts, _voter)
}

// GetVoterStakedProposals is a free data retrieval call binding the contract method 0x6ae1e541.
//
// Solidity: function getVoterStakedProposals(address _voter) view returns(uint256[])
func (_ConvictionBeta *ConvictionBetaCaller) GetVoterStakedProposals(opts *bind.CallOpts, _voter common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "getVoterStakedProposals", _voter)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetVoterStakedProposals is a free data retrieval call binding the contract method 0x6ae1e541.
//
// Solidity: function getVoterStakedProposals(address _voter) view returns(uint256[])
func (_ConvictionBeta *ConvictionBetaSession) GetVoterStakedProposals(_voter common.Address) ([]*big.Int, error) {
	return _ConvictionBeta.Contract.GetVoterStakedProposals(&_ConvictionBeta.CallOpts, _voter)
}

// GetVoterStakedProposals is a free data retrieval call binding the contract method 0x6ae1e541.
//
// Solidity: function getVoterStakedProposals(address _voter) view returns(uint256[])
func (_ConvictionBeta *ConvictionBetaCallerSession) GetVoterStakedProposals(_voter common.Address) ([]*big.Int, error) {
	return _ConvictionBeta.Contract.GetVoterStakedProposals(&_ConvictionBeta.CallOpts, _voter)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ConvictionBeta *ConvictionBetaCaller) HasInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "hasInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) HasInitialized() (bool, error) {
	return _ConvictionBeta.Contract.HasInitialized(&_ConvictionBeta.CallOpts)
}

// HasInitialized is a free data retrieval call binding the contract method 0x0803fac0.
//
// Solidity: function hasInitialized() view returns(bool)
func (_ConvictionBeta *ConvictionBetaCallerSession) HasInitialized() (bool, error) {
	return _ConvictionBeta.Contract.HasInitialized(&_ConvictionBeta.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ConvictionBeta *ConvictionBetaCaller) IsPetrified(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "isPetrified")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) IsPetrified() (bool, error) {
	return _ConvictionBeta.Contract.IsPetrified(&_ConvictionBeta.CallOpts)
}

// IsPetrified is a free data retrieval call binding the contract method 0xde4796ed.
//
// Solidity: function isPetrified() view returns(bool)
func (_ConvictionBeta *ConvictionBetaCallerSession) IsPetrified() (bool, error) {
	return _ConvictionBeta.Contract.IsPetrified(&_ConvictionBeta.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) Kernel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "kernel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) Kernel() (common.Address, error) {
	return _ConvictionBeta.Contract.Kernel(&_ConvictionBeta.CallOpts)
}

// Kernel is a free data retrieval call binding the contract method 0xd4aae0c4.
//
// Solidity: function kernel() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) Kernel() (common.Address, error) {
	return _ConvictionBeta.Contract.Kernel(&_ConvictionBeta.CallOpts)
}

// MaxRatio is a free data retrieval call binding the contract method 0x76365af7.
//
// Solidity: function maxRatio() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) MaxRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "maxRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRatio is a free data retrieval call binding the contract method 0x76365af7.
//
// Solidity: function maxRatio() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) MaxRatio() (*big.Int, error) {
	return _ConvictionBeta.Contract.MaxRatio(&_ConvictionBeta.CallOpts)
}

// MaxRatio is a free data retrieval call binding the contract method 0x76365af7.
//
// Solidity: function maxRatio() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) MaxRatio() (*big.Int, error) {
	return _ConvictionBeta.Contract.MaxRatio(&_ConvictionBeta.CallOpts)
}

// MinThresholdStakePercentage is a free data retrieval call binding the contract method 0xdb2d878c.
//
// Solidity: function minThresholdStakePercentage() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) MinThresholdStakePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "minThresholdStakePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinThresholdStakePercentage is a free data retrieval call binding the contract method 0xdb2d878c.
//
// Solidity: function minThresholdStakePercentage() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) MinThresholdStakePercentage() (*big.Int, error) {
	return _ConvictionBeta.Contract.MinThresholdStakePercentage(&_ConvictionBeta.CallOpts)
}

// MinThresholdStakePercentage is a free data retrieval call binding the contract method 0xdb2d878c.
//
// Solidity: function minThresholdStakePercentage() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) MinThresholdStakePercentage() (*big.Int, error) {
	return _ConvictionBeta.Contract.MinThresholdStakePercentage(&_ConvictionBeta.CallOpts)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) ProposalCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "proposalCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) ProposalCounter() (*big.Int, error) {
	return _ConvictionBeta.Contract.ProposalCounter(&_ConvictionBeta.CallOpts)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) ProposalCounter() (*big.Int, error) {
	return _ConvictionBeta.Contract.ProposalCounter(&_ConvictionBeta.CallOpts)
}

// RequestToken is a free data retrieval call binding the contract method 0x4d7ec628.
//
// Solidity: function requestToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) RequestToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "requestToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RequestToken is a free data retrieval call binding the contract method 0x4d7ec628.
//
// Solidity: function requestToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) RequestToken() (common.Address, error) {
	return _ConvictionBeta.Contract.RequestToken(&_ConvictionBeta.CallOpts)
}

// RequestToken is a free data retrieval call binding the contract method 0x4d7ec628.
//
// Solidity: function requestToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) RequestToken() (common.Address, error) {
	return _ConvictionBeta.Contract.RequestToken(&_ConvictionBeta.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) StakeToken() (common.Address, error) {
	return _ConvictionBeta.Contract.StakeToken(&_ConvictionBeta.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) StakeToken() (common.Address, error) {
	return _ConvictionBeta.Contract.StakeToken(&_ConvictionBeta.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) TotalStaked() (*big.Int, error) {
	return _ConvictionBeta.Contract.TotalStaked(&_ConvictionBeta.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) TotalStaked() (*big.Int, error) {
	return _ConvictionBeta.Contract.TotalStaked(&_ConvictionBeta.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvictionBeta *ConvictionBetaCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvictionBeta *ConvictionBetaSession) Vault() (common.Address, error) {
	return _ConvictionBeta.Contract.Vault(&_ConvictionBeta.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_ConvictionBeta *ConvictionBetaCallerSession) Vault() (common.Address, error) {
	return _ConvictionBeta.Contract.Vault(&_ConvictionBeta.CallOpts)
}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCaller) Weight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConvictionBeta.contract.Call(opts, &out, "weight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaSession) Weight() (*big.Int, error) {
	return _ConvictionBeta.Contract.Weight(&_ConvictionBeta.CallOpts)
}

// Weight is a free data retrieval call binding the contract method 0xa1aab33f.
//
// Solidity: function weight() view returns(uint256)
func (_ConvictionBeta *ConvictionBetaCallerSession) Weight() (*big.Int, error) {
	return _ConvictionBeta.Contract.Weight(&_ConvictionBeta.CallOpts)
}

// AddProposal is a paid mutator transaction binding the contract method 0x03341082.
//
// Solidity: function addProposal(string _title, bytes _link, uint256 _requestedAmount, address _beneficiary) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) AddProposal(opts *bind.TransactOpts, _title string, _link []byte, _requestedAmount *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "addProposal", _title, _link, _requestedAmount, _beneficiary)
}

// AddProposal is a paid mutator transaction binding the contract method 0x03341082.
//
// Solidity: function addProposal(string _title, bytes _link, uint256 _requestedAmount, address _beneficiary) returns()
func (_ConvictionBeta *ConvictionBetaSession) AddProposal(_title string, _link []byte, _requestedAmount *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.AddProposal(&_ConvictionBeta.TransactOpts, _title, _link, _requestedAmount, _beneficiary)
}

// AddProposal is a paid mutator transaction binding the contract method 0x03341082.
//
// Solidity: function addProposal(string _title, bytes _link, uint256 _requestedAmount, address _beneficiary) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) AddProposal(_title string, _link []byte, _requestedAmount *big.Int, _beneficiary common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.AddProposal(&_ConvictionBeta.TransactOpts, _title, _link, _requestedAmount, _beneficiary)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) CancelProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "cancelProposal", _proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaSession) CancelProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.CancelProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// CancelProposal is a paid mutator transaction binding the contract method 0xe0a8f6f5.
//
// Solidity: function cancelProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) CancelProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.CancelProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) ExecuteProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "executeProposal", _proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaSession) ExecuteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.ExecuteProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) ExecuteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.ExecuteProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address _stakeToken, address _vault, address _requestToken, uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) Initialize(opts *bind.TransactOpts, _stakeToken common.Address, _vault common.Address, _requestToken common.Address, _decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "initialize", _stakeToken, _vault, _requestToken, _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address _stakeToken, address _vault, address _requestToken, uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaSession) Initialize(_stakeToken common.Address, _vault common.Address, _requestToken common.Address, _decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.Initialize(&_ConvictionBeta.TransactOpts, _stakeToken, _vault, _requestToken, _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x2b4656c8.
//
// Solidity: function initialize(address _stakeToken, address _vault, address _requestToken, uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) Initialize(_stakeToken common.Address, _vault common.Address, _requestToken common.Address, _decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.Initialize(&_ConvictionBeta.TransactOpts, _stakeToken, _vault, _requestToken, _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// OnApprove is a paid mutator transaction binding the contract method 0xda682aeb.
//
// Solidity: function onApprove(address _holder, address _spender, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaTransactor) OnApprove(opts *bind.TransactOpts, _holder common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "onApprove", _holder, _spender, _amount)
}

// OnApprove is a paid mutator transaction binding the contract method 0xda682aeb.
//
// Solidity: function onApprove(address _holder, address _spender, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) OnApprove(_holder common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnApprove(&_ConvictionBeta.TransactOpts, _holder, _spender, _amount)
}

// OnApprove is a paid mutator transaction binding the contract method 0xda682aeb.
//
// Solidity: function onApprove(address _holder, address _spender, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaTransactorSession) OnApprove(_holder common.Address, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnApprove(&_ConvictionBeta.TransactOpts, _holder, _spender, _amount)
}

// OnRegisterAsHook is a paid mutator transaction binding the contract method 0x4d4eb6ce.
//
// Solidity: function onRegisterAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) OnRegisterAsHook(opts *bind.TransactOpts, _hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "onRegisterAsHook", _hookId, _token)
}

// OnRegisterAsHook is a paid mutator transaction binding the contract method 0x4d4eb6ce.
//
// Solidity: function onRegisterAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaSession) OnRegisterAsHook(_hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnRegisterAsHook(&_ConvictionBeta.TransactOpts, _hookId, _token)
}

// OnRegisterAsHook is a paid mutator transaction binding the contract method 0x4d4eb6ce.
//
// Solidity: function onRegisterAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) OnRegisterAsHook(_hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnRegisterAsHook(&_ConvictionBeta.TransactOpts, _hookId, _token)
}

// OnRevokeAsHook is a paid mutator transaction binding the contract method 0xc2239b24.
//
// Solidity: function onRevokeAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) OnRevokeAsHook(opts *bind.TransactOpts, _hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "onRevokeAsHook", _hookId, _token)
}

// OnRevokeAsHook is a paid mutator transaction binding the contract method 0xc2239b24.
//
// Solidity: function onRevokeAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaSession) OnRevokeAsHook(_hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnRevokeAsHook(&_ConvictionBeta.TransactOpts, _hookId, _token)
}

// OnRevokeAsHook is a paid mutator transaction binding the contract method 0xc2239b24.
//
// Solidity: function onRevokeAsHook(uint256 _hookId, address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) OnRevokeAsHook(_hookId *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnRevokeAsHook(&_ConvictionBeta.TransactOpts, _hookId, _token)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaTransactor) OnTransfer(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "onTransfer", _from, _to, _amount)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaSession) OnTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnTransfer(&_ConvictionBeta.TransactOpts, _from, _to, _amount)
}

// OnTransfer is a paid mutator transaction binding the contract method 0x4a393149.
//
// Solidity: function onTransfer(address _from, address _to, uint256 _amount) returns(bool)
func (_ConvictionBeta *ConvictionBetaTransactorSession) OnTransfer(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.OnTransfer(&_ConvictionBeta.TransactOpts, _from, _to, _amount)
}

// SetConvictionCalculationSettings is a paid mutator transaction binding the contract method 0xc35ac76d.
//
// Solidity: function setConvictionCalculationSettings(uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) SetConvictionCalculationSettings(opts *bind.TransactOpts, _decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "setConvictionCalculationSettings", _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// SetConvictionCalculationSettings is a paid mutator transaction binding the contract method 0xc35ac76d.
//
// Solidity: function setConvictionCalculationSettings(uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaSession) SetConvictionCalculationSettings(_decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.SetConvictionCalculationSettings(&_ConvictionBeta.TransactOpts, _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// SetConvictionCalculationSettings is a paid mutator transaction binding the contract method 0xc35ac76d.
//
// Solidity: function setConvictionCalculationSettings(uint256 _decay, uint256 _maxRatio, uint256 _weight, uint256 _minThresholdStakePercentage) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) SetConvictionCalculationSettings(_decay *big.Int, _maxRatio *big.Int, _weight *big.Int, _minThresholdStakePercentage *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.SetConvictionCalculationSettings(&_ConvictionBeta.TransactOpts, _decay, _maxRatio, _weight, _minThresholdStakePercentage)
}

// StakeAllToProposal is a paid mutator transaction binding the contract method 0x1df67d2f.
//
// Solidity: function stakeAllToProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) StakeAllToProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "stakeAllToProposal", _proposalId)
}

// StakeAllToProposal is a paid mutator transaction binding the contract method 0x1df67d2f.
//
// Solidity: function stakeAllToProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaSession) StakeAllToProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.StakeAllToProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// StakeAllToProposal is a paid mutator transaction binding the contract method 0x1df67d2f.
//
// Solidity: function stakeAllToProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) StakeAllToProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.StakeAllToProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// StakeToProposal is a paid mutator transaction binding the contract method 0xfc370051.
//
// Solidity: function stakeToProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) StakeToProposal(opts *bind.TransactOpts, _proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "stakeToProposal", _proposalId, _amount)
}

// StakeToProposal is a paid mutator transaction binding the contract method 0xfc370051.
//
// Solidity: function stakeToProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaSession) StakeToProposal(_proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.StakeToProposal(&_ConvictionBeta.TransactOpts, _proposalId, _amount)
}

// StakeToProposal is a paid mutator transaction binding the contract method 0xfc370051.
//
// Solidity: function stakeToProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) StakeToProposal(_proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.StakeToProposal(&_ConvictionBeta.TransactOpts, _proposalId, _amount)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) TransferToVault(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "transferToVault", _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ConvictionBeta *ConvictionBetaSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.TransferToVault(&_ConvictionBeta.TransactOpts, _token)
}

// TransferToVault is a paid mutator transaction binding the contract method 0x9d4941d8.
//
// Solidity: function transferToVault(address _token) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) TransferToVault(_token common.Address) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.TransferToVault(&_ConvictionBeta.TransactOpts, _token)
}

// WithdrawAllFromProposal is a paid mutator transaction binding the contract method 0x4fb3cbbc.
//
// Solidity: function withdrawAllFromProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) WithdrawAllFromProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "withdrawAllFromProposal", _proposalId)
}

// WithdrawAllFromProposal is a paid mutator transaction binding the contract method 0x4fb3cbbc.
//
// Solidity: function withdrawAllFromProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaSession) WithdrawAllFromProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.WithdrawAllFromProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// WithdrawAllFromProposal is a paid mutator transaction binding the contract method 0x4fb3cbbc.
//
// Solidity: function withdrawAllFromProposal(uint256 _proposalId) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) WithdrawAllFromProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.WithdrawAllFromProposal(&_ConvictionBeta.TransactOpts, _proposalId)
}

// WithdrawFromProposal is a paid mutator transaction binding the contract method 0x649c649c.
//
// Solidity: function withdrawFromProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaTransactor) WithdrawFromProposal(opts *bind.TransactOpts, _proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.contract.Transact(opts, "withdrawFromProposal", _proposalId, _amount)
}

// WithdrawFromProposal is a paid mutator transaction binding the contract method 0x649c649c.
//
// Solidity: function withdrawFromProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaSession) WithdrawFromProposal(_proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.WithdrawFromProposal(&_ConvictionBeta.TransactOpts, _proposalId, _amount)
}

// WithdrawFromProposal is a paid mutator transaction binding the contract method 0x649c649c.
//
// Solidity: function withdrawFromProposal(uint256 _proposalId, uint256 _amount) returns()
func (_ConvictionBeta *ConvictionBetaTransactorSession) WithdrawFromProposal(_proposalId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ConvictionBeta.Contract.WithdrawFromProposal(&_ConvictionBeta.TransactOpts, _proposalId, _amount)
}

// ConvictionBetaConvictionSettingsChangedIterator is returned from FilterConvictionSettingsChanged and is used to iterate over the raw logs and unpacked data for ConvictionSettingsChanged events raised by the ConvictionBeta contract.
type ConvictionBetaConvictionSettingsChangedIterator struct {
	Event *ConvictionBetaConvictionSettingsChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaConvictionSettingsChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaConvictionSettingsChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaConvictionSettingsChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaConvictionSettingsChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaConvictionSettingsChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaConvictionSettingsChanged represents a ConvictionSettingsChanged event raised by the ConvictionBeta contract.
type ConvictionBetaConvictionSettingsChanged struct {
	Decay                       *big.Int
	MaxRatio                    *big.Int
	Weight                      *big.Int
	MinThresholdStakePercentage *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterConvictionSettingsChanged is a free log retrieval operation binding the contract event 0xab0e1d16e026a71973fa1c1862074f152ee6c430f36fce58005928c3bb158836.
//
// Solidity: event ConvictionSettingsChanged(uint256 decay, uint256 maxRatio, uint256 weight, uint256 minThresholdStakePercentage)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterConvictionSettingsChanged(opts *bind.FilterOpts) (*ConvictionBetaConvictionSettingsChangedIterator, error) {

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "ConvictionSettingsChanged")
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaConvictionSettingsChangedIterator{contract: _ConvictionBeta.contract, event: "ConvictionSettingsChanged", logs: logs, sub: sub}, nil
}

// WatchConvictionSettingsChanged is a free log subscription operation binding the contract event 0xab0e1d16e026a71973fa1c1862074f152ee6c430f36fce58005928c3bb158836.
//
// Solidity: event ConvictionSettingsChanged(uint256 decay, uint256 maxRatio, uint256 weight, uint256 minThresholdStakePercentage)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchConvictionSettingsChanged(opts *bind.WatchOpts, sink chan<- *ConvictionBetaConvictionSettingsChanged) (event.Subscription, error) {

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "ConvictionSettingsChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaConvictionSettingsChanged)
				if err := _ConvictionBeta.contract.UnpackLog(event, "ConvictionSettingsChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConvictionSettingsChanged is a log parse operation binding the contract event 0xab0e1d16e026a71973fa1c1862074f152ee6c430f36fce58005928c3bb158836.
//
// Solidity: event ConvictionSettingsChanged(uint256 decay, uint256 maxRatio, uint256 weight, uint256 minThresholdStakePercentage)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseConvictionSettingsChanged(log types.Log) (*ConvictionBetaConvictionSettingsChanged, error) {
	event := new(ConvictionBetaConvictionSettingsChanged)
	if err := _ConvictionBeta.contract.UnpackLog(event, "ConvictionSettingsChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaProposalAddedIterator is returned from FilterProposalAdded and is used to iterate over the raw logs and unpacked data for ProposalAdded events raised by the ConvictionBeta contract.
type ConvictionBetaProposalAddedIterator struct {
	Event *ConvictionBetaProposalAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaProposalAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaProposalAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaProposalAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaProposalAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaProposalAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaProposalAdded represents a ProposalAdded event raised by the ConvictionBeta contract.
type ConvictionBetaProposalAdded struct {
	Entity      common.Address
	Id          *big.Int
	Title       string
	Link        []byte
	Amount      *big.Int
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalAdded is a free log retrieval operation binding the contract event 0x9fc9ab43300923a3b8d47f890603798c652956421ae650ed96c526f311d2c2ac.
//
// Solidity: event ProposalAdded(address indexed entity, uint256 indexed id, string title, bytes link, uint256 amount, address beneficiary)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterProposalAdded(opts *bind.FilterOpts, entity []common.Address, id []*big.Int) (*ConvictionBetaProposalAddedIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "ProposalAdded", entityRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaProposalAddedIterator{contract: _ConvictionBeta.contract, event: "ProposalAdded", logs: logs, sub: sub}, nil
}

// WatchProposalAdded is a free log subscription operation binding the contract event 0x9fc9ab43300923a3b8d47f890603798c652956421ae650ed96c526f311d2c2ac.
//
// Solidity: event ProposalAdded(address indexed entity, uint256 indexed id, string title, bytes link, uint256 amount, address beneficiary)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchProposalAdded(opts *bind.WatchOpts, sink chan<- *ConvictionBetaProposalAdded, entity []common.Address, id []*big.Int) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "ProposalAdded", entityRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaProposalAdded)
				if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalAdded is a log parse operation binding the contract event 0x9fc9ab43300923a3b8d47f890603798c652956421ae650ed96c526f311d2c2ac.
//
// Solidity: event ProposalAdded(address indexed entity, uint256 indexed id, string title, bytes link, uint256 amount, address beneficiary)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseProposalAdded(log types.Log) (*ConvictionBetaProposalAdded, error) {
	event := new(ConvictionBetaProposalAdded)
	if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaProposalCancelledIterator is returned from FilterProposalCancelled and is used to iterate over the raw logs and unpacked data for ProposalCancelled events raised by the ConvictionBeta contract.
type ConvictionBetaProposalCancelledIterator struct {
	Event *ConvictionBetaProposalCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaProposalCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaProposalCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaProposalCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaProposalCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaProposalCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaProposalCancelled represents a ProposalCancelled event raised by the ConvictionBeta contract.
type ConvictionBetaProposalCancelled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterProposalCancelled is a free log retrieval operation binding the contract event 0x416e669c63d9a3a5e36ee7cc7e2104b8db28ccd286aa18966e98fa230c73b08c.
//
// Solidity: event ProposalCancelled(uint256 indexed id)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterProposalCancelled(opts *bind.FilterOpts, id []*big.Int) (*ConvictionBetaProposalCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "ProposalCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaProposalCancelledIterator{contract: _ConvictionBeta.contract, event: "ProposalCancelled", logs: logs, sub: sub}, nil
}

// WatchProposalCancelled is a free log subscription operation binding the contract event 0x416e669c63d9a3a5e36ee7cc7e2104b8db28ccd286aa18966e98fa230c73b08c.
//
// Solidity: event ProposalCancelled(uint256 indexed id)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchProposalCancelled(opts *bind.WatchOpts, sink chan<- *ConvictionBetaProposalCancelled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "ProposalCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaProposalCancelled)
				if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCancelled is a log parse operation binding the contract event 0x416e669c63d9a3a5e36ee7cc7e2104b8db28ccd286aa18966e98fa230c73b08c.
//
// Solidity: event ProposalCancelled(uint256 indexed id)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseProposalCancelled(log types.Log) (*ConvictionBetaProposalCancelled, error) {
	event := new(ConvictionBetaProposalCancelled)
	if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the ConvictionBeta contract.
type ConvictionBetaProposalExecutedIterator struct {
	Event *ConvictionBetaProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaProposalExecuted represents a ProposalExecuted event raised by the ConvictionBeta contract.
type ConvictionBetaProposalExecuted struct {
	Id         *big.Int
	Conviction *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed id, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterProposalExecuted(opts *bind.FilterOpts, id []*big.Int) (*ConvictionBetaProposalExecutedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "ProposalExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaProposalExecutedIterator{contract: _ConvictionBeta.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed id, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *ConvictionBetaProposalExecuted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "ProposalExecuted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaProposalExecuted)
				if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0xf758fc91e01b00ea6b4a6138756f7f28e021f9bf21db6dbf8c36c88eb737257a.
//
// Solidity: event ProposalExecuted(uint256 indexed id, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseProposalExecuted(log types.Log) (*ConvictionBetaProposalExecuted, error) {
	event := new(ConvictionBetaProposalExecuted)
	if err := _ConvictionBeta.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaRecoverToVaultIterator is returned from FilterRecoverToVault and is used to iterate over the raw logs and unpacked data for RecoverToVault events raised by the ConvictionBeta contract.
type ConvictionBetaRecoverToVaultIterator struct {
	Event *ConvictionBetaRecoverToVault // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaRecoverToVaultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaRecoverToVault)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaRecoverToVault)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaRecoverToVaultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaRecoverToVaultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaRecoverToVault represents a RecoverToVault event raised by the ConvictionBeta contract.
type ConvictionBetaRecoverToVault struct {
	Vault  common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecoverToVault is a free log retrieval operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterRecoverToVault(opts *bind.FilterOpts, vault []common.Address, token []common.Address) (*ConvictionBetaRecoverToVaultIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaRecoverToVaultIterator{contract: _ConvictionBeta.contract, event: "RecoverToVault", logs: logs, sub: sub}, nil
}

// WatchRecoverToVault is a free log subscription operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchRecoverToVault(opts *bind.WatchOpts, sink chan<- *ConvictionBetaRecoverToVault, vault []common.Address, token []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "RecoverToVault", vaultRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaRecoverToVault)
				if err := _ConvictionBeta.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecoverToVault is a log parse operation binding the contract event 0x596caf56044b55fb8c4ca640089bbc2b63cae3e978b851f5745cbb7c5b288e02.
//
// Solidity: event RecoverToVault(address indexed vault, address indexed token, uint256 amount)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseRecoverToVault(log types.Log) (*ConvictionBetaRecoverToVault, error) {
	event := new(ConvictionBetaRecoverToVault)
	if err := _ConvictionBeta.contract.UnpackLog(event, "RecoverToVault", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaScriptResultIterator is returned from FilterScriptResult and is used to iterate over the raw logs and unpacked data for ScriptResult events raised by the ConvictionBeta contract.
type ConvictionBetaScriptResultIterator struct {
	Event *ConvictionBetaScriptResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaScriptResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaScriptResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaScriptResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaScriptResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaScriptResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaScriptResult represents a ScriptResult event raised by the ConvictionBeta contract.
type ConvictionBetaScriptResult struct {
	Executor   common.Address
	Script     []byte
	Input      []byte
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterScriptResult is a free log retrieval operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterScriptResult(opts *bind.FilterOpts, executor []common.Address) (*ConvictionBetaScriptResultIterator, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaScriptResultIterator{contract: _ConvictionBeta.contract, event: "ScriptResult", logs: logs, sub: sub}, nil
}

// WatchScriptResult is a free log subscription operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchScriptResult(opts *bind.WatchOpts, sink chan<- *ConvictionBetaScriptResult, executor []common.Address) (event.Subscription, error) {

	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "ScriptResult", executorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaScriptResult)
				if err := _ConvictionBeta.contract.UnpackLog(event, "ScriptResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseScriptResult is a log parse operation binding the contract event 0x5229a5dba83a54ae8cb5b51bdd6de9474cacbe9dd332f5185f3a4f4f2e3f4ad9.
//
// Solidity: event ScriptResult(address indexed executor, bytes script, bytes input, bytes returnData)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseScriptResult(log types.Log) (*ConvictionBetaScriptResult, error) {
	event := new(ConvictionBetaScriptResult)
	if err := _ConvictionBeta.contract.UnpackLog(event, "ScriptResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaStakeAddedIterator is returned from FilterStakeAdded and is used to iterate over the raw logs and unpacked data for StakeAdded events raised by the ConvictionBeta contract.
type ConvictionBetaStakeAddedIterator struct {
	Event *ConvictionBetaStakeAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaStakeAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaStakeAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaStakeAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaStakeAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaStakeAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaStakeAdded represents a StakeAdded event raised by the ConvictionBeta contract.
type ConvictionBetaStakeAdded struct {
	Entity            common.Address
	Id                *big.Int
	Amount            *big.Int
	TokensStaked      *big.Int
	TotalTokensStaked *big.Int
	Conviction        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStakeAdded is a free log retrieval operation binding the contract event 0x28d9b583e0c477691a08f6c1e00fedc0895ed4221487c627fa96a7024119f499.
//
// Solidity: event StakeAdded(address indexed entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterStakeAdded(opts *bind.FilterOpts, entity []common.Address, id []*big.Int) (*ConvictionBetaStakeAddedIterator, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "StakeAdded", entityRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaStakeAddedIterator{contract: _ConvictionBeta.contract, event: "StakeAdded", logs: logs, sub: sub}, nil
}

// WatchStakeAdded is a free log subscription operation binding the contract event 0x28d9b583e0c477691a08f6c1e00fedc0895ed4221487c627fa96a7024119f499.
//
// Solidity: event StakeAdded(address indexed entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchStakeAdded(opts *bind.WatchOpts, sink chan<- *ConvictionBetaStakeAdded, entity []common.Address, id []*big.Int) (event.Subscription, error) {

	var entityRule []interface{}
	for _, entityItem := range entity {
		entityRule = append(entityRule, entityItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "StakeAdded", entityRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaStakeAdded)
				if err := _ConvictionBeta.contract.UnpackLog(event, "StakeAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeAdded is a log parse operation binding the contract event 0x28d9b583e0c477691a08f6c1e00fedc0895ed4221487c627fa96a7024119f499.
//
// Solidity: event StakeAdded(address indexed entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseStakeAdded(log types.Log) (*ConvictionBetaStakeAdded, error) {
	event := new(ConvictionBetaStakeAdded)
	if err := _ConvictionBeta.contract.UnpackLog(event, "StakeAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ConvictionBetaStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the ConvictionBeta contract.
type ConvictionBetaStakeWithdrawnIterator struct {
	Event *ConvictionBetaStakeWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ConvictionBetaStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConvictionBetaStakeWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ConvictionBetaStakeWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ConvictionBetaStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConvictionBetaStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConvictionBetaStakeWithdrawn represents a StakeWithdrawn event raised by the ConvictionBeta contract.
type ConvictionBetaStakeWithdrawn struct {
	Entity            common.Address
	Id                *big.Int
	Amount            *big.Int
	TokensStaked      *big.Int
	TotalTokensStaked *big.Int
	Conviction        *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x16f23283da3097bc9027dcdf31f24863b1520556f04818d406f0e6ecd08580f5.
//
// Solidity: event StakeWithdrawn(address entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, id []*big.Int) (*ConvictionBetaStakeWithdrawnIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.FilterLogs(opts, "StakeWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return &ConvictionBetaStakeWithdrawnIterator{contract: _ConvictionBeta.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x16f23283da3097bc9027dcdf31f24863b1520556f04818d406f0e6ecd08580f5.
//
// Solidity: event StakeWithdrawn(address entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *ConvictionBetaStakeWithdrawn, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ConvictionBeta.contract.WatchLogs(opts, "StakeWithdrawn", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConvictionBetaStakeWithdrawn)
				if err := _ConvictionBeta.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x16f23283da3097bc9027dcdf31f24863b1520556f04818d406f0e6ecd08580f5.
//
// Solidity: event StakeWithdrawn(address entity, uint256 indexed id, uint256 amount, uint256 tokensStaked, uint256 totalTokensStaked, uint256 conviction)
func (_ConvictionBeta *ConvictionBetaFilterer) ParseStakeWithdrawn(log types.Log) (*ConvictionBetaStakeWithdrawn, error) {
	event := new(ConvictionBetaStakeWithdrawn)
	if err := _ConvictionBeta.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}
