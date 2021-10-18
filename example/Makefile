.PHONY: dump
dump:
	mkdir -p out
	# Generate Config
	nftool traits dump --layers ./layers --out ./out/config.yaml

.PHONY: collection
collection:
	# Generate List of Collection Attributes
	nftool traits make --amount 10 --config ./out/config.yaml --out ./out/collection.json

.PHONY: images
images:
	# Generate Images from Collection Attributes
	mkdir -p ./out/images
	nftool img gen --width 300 --height 300 --collection ./out/collection.json --config ./out/config.yaml --out ./out/images

.PHONY: rarity
rarity:
	# Generate traits rarity report
	nftool rarity traits --collection ./out/collection.json --out ./out/traits_rarity.json

	# Generate collection rarity report
	nftool rarity collection --collection ./out/collection.json --out ./out/collection_rarity.json

.PHONY: provenance
provenance:
	# Generate provenace report
	nftool provenance --images ./out/images --out ./out/provenance.json --startingIndex 2