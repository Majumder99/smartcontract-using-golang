// SPDX-License-Identifier: MIT
pragma solidity 0.8.23;

contract Todo {
    address public owner;
    Task[] tasks;

    struct Task {
        string content;
        bool status;
    }

    constructor() {
        owner = msg.sender;
    }

    modifier isOwner() {
        require(msg.sender == owner, "You are not the owner");
        _;
    }

    function add(string memory _content) public isOwner {
        tasks.push(Task(_content, false));
    }

    function get(uint _id) public view isOwner returns (Task memory) {
        return tasks[_id];
    }

    function list() public view isOwner returns (Task[] memory) {
        return tasks;
    }

    function update(uint _id, string memory _newcontent) public isOwner {
        tasks[_id].content = _newcontent;
    }

    function remove(uint _id) public isOwner {
        delete tasks[_id];
    }
}
