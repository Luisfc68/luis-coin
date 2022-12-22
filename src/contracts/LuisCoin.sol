//SPDX-License-Identifier: UNLICENSED

pragma solidity >=0.7.0 <0.9.0;

contract LuisCoin {

    address public minter;
    mapping (address => uint) balances;

    event Sent(address from, address to, uint amount);

    constructor() {
        minter = msg.sender;
        balances[minter] = 50;
    }

    function mint(address receiver, uint amount) public {
        if (msg.sender != minter) {
            return;
        }
        balances[receiver] += amount;
    }

    function send(address receiver, uint amount) external {
        if (balances[msg.sender] < amount) {
            return;
        }
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        emit Sent(msg.sender, receiver, amount);
    }

    function checkBalance(address account) external view returns (uint) {
        return balances[account];
    }

}