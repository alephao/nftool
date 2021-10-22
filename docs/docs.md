### `nftool img gen`

Generate images from a collection.json and the layers folder

```
nftool img gen [flags]
```

**Example**:

```
nftool img gen \
	--width 800 \
	--height 800 \
	--collection ./collection.json \
	--config ./config.yaml \
	--out ./generated-imgs \
	--parallel
```

**Options:**

```
      --collection nftool traits make   path to the collection json file generated via nftool traits make
      --config nftool traits dump       path to the configuration file generated via nftool traits dump
      --height int                      the height of the generated image
      --onlyMissing                     generate only the images that are missing in the output dir
      --out string                      path to the folder where all imgs will be generated
      --parallel                        generate images in parallel using all your cores
      --png                             generate png images instead of jpg images
      --startingIndex int               the initial index of the image names, e.g.: you have a collection of 100 and want the names to be from 1000 to 1100, use --starrtingIndex 1000
      --width int                       the width of the generated image
```

### `nftool metadata`

Generate ERC-721 and ERC-1155 compliant metadata from a collection json generated with `nftool traits make`.

This command will generate one file for each token, so if you have 1000 items in the collection json, it will generate 1000 files.

```
nftool metadata [flags]
```

**Example**:

```
nftool metadata \
	--collection ./collection.json \
	--config ./config.yaml \
	--out ./metadata
		
```

**Options:**

```
      --collection nftool traits make   path to the collection json generated via nftool traits make
      --config nftool traits dump       path to the config json generated via nftool traits dump
      --erc1155                         if the file names generated should follow the ERC-1155 pattern: hex value padded with zeros to size 64
      --out string                      path to the folder where the generated metadata files will be created
```

### `nftool opensea update`

Asks opensea to update the metadata of your collection.

```
nftool opensea update [flags]
```

**Example**:

```
nftool opensea update \
	--contract 0x0000000000000000000000000000000000000000 \
	--from 0 \
	--to 1000 \
	--parallel 4
```

**Options:**

```
      --contract string   address of the contract in mainnet
      --from int          the starting id for a the range of ids you want to update
      --parallel int      divide requests between your cpus (default 1)
      --to int            the last id for a the range of ids you want to update
```

### `nftool provenance`

Generate the provenace report of a collection

How provenance is generated:

First we generate the hash for each image using a SHA256 algorithm

Then we append each hash following a specific order. The order is the original order offseted by the "startingIndex". We calculate each new index of images with the following formula:

(tokenId + startingIndex) \% collectionSize

Usually the startIndex is a number that is randomly generated on-chain.

After appending all hashes, we hash the result using SHA-256 again and that's the "final proof". This proof is usually stored on-chain.

The provenance report contains
- final proof hash
- concatenated hashes
- hashes
- starting index

```
nftool provenance [flags]
```

**Example**:

```
nftool provenance \
	--imgs ./imgs \
	--startingIndex 123 \
	--out ./provenance.json
```

**Options:**

```
      --images string       path to the directory containing the generated images
      --out string          where to save the provenance report
      --startingIndex int   startingIndex
```

### `nftool rarity collection`

Generate collection rarity report

The report contains all the items in the collection, ordered by most rare from least rare.

```
nftool rarity collection [flags]
```

**Example**:

```
nftool rarity collection \
	--collection ./out/collection.json \
	--out ./out/collection_rarity.json
```

**Options:**

```
      --collection string   path to collection json
      --out string          output path for the collection report, should have a .json extension
```

### `nftool rarity traits`

Generate traits rarity report

The report contains the number of appearances of each trait in the collection.

```
nftool rarity traits [flags]
```

**Example**:

```
nftool rarity traits \
	--collection ./collection.json \
	--out ./traits_rarity.json
```

**Options:**

```
      --collection string   path to collection json
      --out string          output path for the trait report, should have a .json extension
```

### `nftool traits collisions`

Find collisions from multiple collection files and generate a report

```
nftool traits collisions [flags]
```

**Example**:

```
nftool traits collisions \
	--file ./collection_1.json \
	--file ./collection_2.json \
	--out ./collision_report.json
```

**Options:**

```
      --file stringArray   list paths to collections
      --out string         path to save the collision report
```

### `nftool traits dump`

Generate a yaml configuration file from a directory containing all the layers and following the layer naming guidelines.

Layer Naming Guidelines:

1. Layer Names

Inside the folder, each directory is a Layer, and you should name them following the pattern:

00-Background
01-Body
02-Face
03-Facial Hair

"Body" will be on top of "Background"
"Face" on top of "Body"
"Facial Hair" on top of "Face"

2. Layer Variant Names and Rarity

Each variant should follow the pattern <rarity weight>-<variant name>.png

00-Background
├── 1-Gray.png
└── 5-Haunted Mansion.png
01-Body
├── 1-Cool Hoodie.png
└── 1-Shinning Armour.png

The rarity of the traits will be calculated as a weighted average. For the example above, here is how we calculate the chance of a variant showing up:

First sum all rarity weights: 1 (Gray) + 5 (Haunted Mansion) = 6
Divide the individal weight by the sum. 
	For Gray: 1/6 = 16.7%
	For Haunted Mansion: 5/6 = 83.3%

In other words, for every 6 images generated, 1 will have the Gray background and 5 will have the Haunted Mansion background.

3. Optional Layers

If you want a layer to be optional, add a "_" suffix to the folder like so:
	
	00-Background_

Then after running 'nftool traits dump', you can change the chance of a trait not beign selected in the generated yaml file, under the "optional_weight" property.

```
nftool traits dump [flags]
```

**Example**:

```
nftool traits dump \
	--layers ./layers \
	--out ./out/config.yaml
```

**Options:**

```
      --layers string   path to where all the layer folders are
      --out string      output path to the config, should have a .yaml extension
```

### `nftool traits make`

Generate a collection from a config file or a folder with the structured layers.

```
nftool traits make [flags]
```

**Example**:

```
nftool traits make \
	--amount 10 \
	--config ./config.yaml \
	--out ./collection.json
```

**Options:**

```
      --amount int      the amount of items you want to generate
      --config string   path the configuration yaml file
      --out string      output path to the collection, should have a .json extension
```

### `nftool traits merge`

Merge multiple attributes files in one, ignoring collisions

```
nftool traits merge [flags]
```

**Example**:

```
nftool traits merge \
	--file ./collection_1.json \
	--file ./collection_2.json \
	--out ./collection_merged.json
```

**Options:**

```
      --file stringArray   list of collection json files to be merged
      --out string         path to save the new collection json
```

### `nftool traits shuffle`

shuffle a collection json

```
nftool traits shuffle [flags]
```

**Example**:

```
nftool traits shuffle --path ./collection.json
```

**Options:**

```
      --path string   path to the collection json
```

### `nftool version`

Print the version number of nftool

```
nftool version
```

