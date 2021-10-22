# nftool

A suite of tools for NFT generative art.

## Features

* Metadata
   * Generate ERC-721 and ERC-1155 compatible metadata
* Traits/Attributes/Properties Generation
   * Configure custom rarity
   * Generate collection attributes configuration file
   * Merge collections
   * Shuffle collection
   * Find collisions between collections
* Image Generation
   * Generate images from collection description
   * Generate images in parallel
   * Generate only missing images (if you delete a few images from the output folder)
* Rarity
   * Generate traits rarity
   * Generate collection rarity
* Provenance
   * Generate provenance
* OpenSea
   * Update metadata of collection

## Install

### Homebrew on Macos

`brew install alephao/formulae/nftool`

### Using Go

`go install github.com/alephao/nftool@latest`

Or clone the repository, change to the root folder and run `go install`.

### Manually

Download the binary from the [releases page](https://github.com/alephao/nftool/releases) and move it to your bin path like `/usr/local/bin`.

## Getting Started

To get started, install `nftool` following the instructions above and cd to the examples folder in this repo.

You can always use `nftool help [command]` to see all the documentatio and all the options of a command.

First generate the configuration file from the layers folders.

```sh
mkdir -p out
nftool traits dump --layers ./layers --out ./out/config.yaml
```

You can open `config.yaml` and play around with the rarity weights, and optional configuration.

Then generate a collection from the `config.yaml`

```sh
nftool traits make --amount 10 --config ./out/config.yaml --out ./out/collection.json
```

Now we can use the collection.json to generate the images. Run the following commands:

```sh
mkdir -p ./out/images
nftool img gen --width 300 --height 300 --collection ./out/collection.json --config ./out/config.yaml --out ./out/images
```

We can also generate rarity reports for traits (which trait is more rare and how many times it shows up) and the collection rarity rank.

```sh
# Generate traits rarity report
nftool rarity traits --collection ./out/collection.json --out ./out/traits_rarity.json

# Generate collection rarity rank report
nftool rarity collection --collection ./out/collection.json --out ./out/collection_rarity.json
```

To generate the provenance for this collection it's easy:

```sh
nftool provenance --images ./out/images --out ./out/provenance.json --startingIndex 2
```

Note: the startingIndex should be a number smaller than the total amount of items in the collection and it's usually generated on-chain.

Now we need to generate the actual metadata that we'll upload to IPFS (or another storage service).

```
	mkdir -p ./out/metadata
	# Generate ERC-721 metadata
	nftool metadata --collection ./out/collection.json --config ./out/config.yaml --out ./out/metadata
```

Note: if you want to generate the metadata following the erc-1155 convention for the id/file-name, add the flag --erc1155

Check out the output in `./out/metadata/1`. If you want to change anything, you can do so by editing the values in `./out/config.yaml`.

## Documentation

[Link to the generated docs](./docs/docs.md)

## Contributing

* For requests and questions, please open an issue.
* PRs accepted.

## License

[MIT © Aleph Retamal](LICENSE)