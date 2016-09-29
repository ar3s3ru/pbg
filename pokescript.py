#!/usr/bin/env python
#
# Script per fregare i dati da PokèAPI e renderli compatibili a PBG.
# Si ringrazia il grande capo Gianni Modica per lo script!
#
import requests
import json

typesTable = {
            "normal": 0, "fire": 1, "fightning": 2, "water": 3, "flying": 4, "grass": 5,
            "poison": 6, "electric": 7, "ground": 8, "psychic": 9, "rock": 10, "ice": 11,
            "bug": 12, "dragon": 13, "ghost": 14, "dark": 15, "steel": 16, "fairy": 17, "???": 18
        }

def typeToGType(typ: str):
    typ1 = typesTable[typ]
    return typ1 if typ1 is not None else -1

def downloadPokèmon(num: int):
	pokes = {}
	pokes["generation"] = 1
	pokes["count"]      = num
	pokes["pokemons"]   = []

	for i in range (1, num + 1):
		r = requests.get('http://pokeapi.co/api/v2/pokemon/'+ str(i) + '/')
		parsedJson = r.json()

		pkm = {}

		pkm["name"]      = parsedJson["name"]
		pkm["name"]      = pkm["name"].title()

		pkm["pokedex"]   = parsedJson["id"]
		pkm["baseStats"] = []

		pkm["type"]      = []
		pkm["type"].insert(0, typeToGType(((parsedJson["types"])[0])["type"]["name"]))

		if len(parsedJson["types"]) > 1:
			pkm["type"].insert(1, typeToGType(((parsedJson["types"])[1])["type"]["name"]))
		else:
			pkm["type"].insert(1, -1)

		j = 0
		for st in parsedJson["stats"]:
			pkm["baseStats"].insert(j, st["base_stat"])
			j=j+1

		pokes["pokemons"].insert(i, pkm)
		print("Donwloaded pokèmon " + str(i) + ": " + str(pkm))

	val = json.dumps(pokes, indent=4, sort_keys=True)
	file = open("pokedb.json","w")
	file.write(val + "\n")
	file.close()


def main():
	downloadPokèmon(20)

if __name__ == '__main__':
	main()
