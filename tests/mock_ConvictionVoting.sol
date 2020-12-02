pragma solidity ^0.4.24;

contract ConvictionVoting {
    uint256 constant public ABSTAIN_PROPOSAL_ID = 1;

    event ProposalAdded(address indexed entity, uint256 indexed id, string title, bytes link, uint256 amount, address beneficiary);

    function AddProposal() public {
        emit ProposalAdded(0x0, ABSTAIN_PROPOSAL_ID, "Abstain proposal", "", 0, 0x0);        
    }
}