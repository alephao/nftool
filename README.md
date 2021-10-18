# nftool

A suite of tools for NFT generative art.

## Features

* Traits/Attributes/Properties Generation
   * Configure custom rarity
   * Generate collection attributes configuration file
   * Merge collections
   * Shuffle collection
   * Find collisions between collections
* Image Generation
   * Generate images from collection description
   * Generate images in parallel
* Rarity
   * Generate traits rarity
   * Generate collection rarity
* Provenance
   * Generate provenance

## Install

### Using Go

Clone the repo, change to the root directory and run `go install`

Currently you can only install via golang. Other ways to install will come soon.

## Getting Started

To get started, install `nftool` following the instructions above and cd to the examples folder in this repo.

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

## Contributing

* For requests and questions, please open an issue.
* PRs accepted.

## License

[MIT © Aleph Retamal](LICENSE)