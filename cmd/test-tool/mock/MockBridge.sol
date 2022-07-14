// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.11;

import "./openzeppelin-contracts/contracts/utils/cryptography/SignatureChecker.sol";
import "./openzeppelin-contracts/contracts/utils/cryptography/ECDSA.sol";

contract MockBridge {
    address public storedMpc;

    constructor(address _mpc) {
        storedMpc = _mpc;
    }

    modifier onlySignedByMPC(bytes32 hash, bytes memory signature) {
        require(SignatureChecker.isValidSignatureNow(mpc(), hash, signature), "MockBridge: invalid signature");
        _;
    }

    function mpc() public view returns (address) {
        return storedMpc;
    }

    function receiveRequestV2Signed(bytes memory _callData, address _receiveSide, bytes memory signature)
        external
        onlySignedByMPC(keccak256(bytes.concat("receiveRequestV2", _callData, bytes20(_receiveSide))), signature)
    {
    }
}
