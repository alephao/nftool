const { ethers } = require("hardhat")
const { MerkleTree } = require("merkletreejs")
const keccak256 = require("keccak256")
const tokens = require("./tokens.json")
const fs = require("fs-extra")

let merkleTree

function hashToken(account, tokenId) {
    return Buffer.from(
        ethers.utils.solidityKeccak256(["address", "uint256"], [account, tokenId]).slice(2),
        "hex"
    )
}

async function main() {
    fs.removeSync("whitelist.json") //cleanup old one

    merkleTree = new MerkleTree(
        Object.entries(tokens).map((token) => hashToken(...token)),
        keccak256,
        { sortPairs: true, hashLeaves: false }
    )

    var data = []
    for (const [account, tokenId] of Object.entries(tokens)) {
        const proof = merkleTree.getHexProof(hashToken(account, tokenId))
        console.log(`${account} tokenId: ${tokenId} proof: ${proof}`)
        data.push({ Address: account, TokenID: tokenId, Proof: proof })
    }

    var dictstring = JSON.stringify(data)
    fs.writeFileSync("whitelist.json", dictstring)

    console.log("---------------------------------------------")
    console.log("Total accounts processed = ", merkleTree.getLeaves().length)
    console.log("---------------------------------------------")
    console.log("Merkle ROOT :>> ", merkleTree.getHexRoot())
    console.log("---------------------------------------------")
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error)
        process.exit(1)
    })
