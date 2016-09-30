#!/usr/bin/env python
# ------------------------------------------------------------------
# Script per fregare i dati da PokèAPI e renderli compatibili a PBG.
# Si ringrazia il grande capo Gianni Modica per lo script!
# ------------------------------------------------------------------
import requests
import json

typesTable = {
    "normal": 0, "fire": 1, "fightning": 2, "water": 3, "flying": 4, "grass": 5,
    "poison": 6, "electric": 7, "ground": 8, "psychic": 9, "rock": 10, "ice": 11,
    "bug": 12, "dragon": 13, "ghost": 14, "dark": 15, "steel": 16, "fairy": 17, "???": 18
}


def type_to_gotype(typ: str):
    typ1 = typesTable[typ]
    return typ1 if typ1 is not None else -1


def adding_type(pkorigin, pkdest):
    pkdest["type"].insert(0, type_to_gotype(((pkorigin["types"])[0])["type"]["name"]))
    if len(pkorigin["types"]) > 1:
        pkdest["type"].insert(1, type_to_gotype(((pkorigin["types"])[1])["type"]["name"]))
    else:
        pkdest["type"].insert(1, -1)


def adding_base_stats(pkorigin, pkdest):
    j = 0
    for st in pkorigin["stats"]:
        pkdest["baseStats"].insert(j, st["base_stat"])
        j = j + 1


def adding_sprites(i: int, pkdest):
    pkdest["sprites"] = {
        # Works with hotlinking
        "front": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/" + str(i) + ".png",
        "back": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/" + str(i) + ".png"
    }


def download_pokèmon(num: int):
    pokes = {
        "generation": 1,
        "count": num,
        "pokemons": []
    }

    for i in range(1, num + 1):
        r = requests.get('http://pokeapi.co/api/v2/pokemon/' + str(i) + '/')
        parsed_json = r.json()
        pkm = {
            "name": parsed_json["name"].title(),
            "pokedex": parsed_json["id"],
            "baseStats": [],
            "type": []
        }

        # Type creation
        adding_type(parsed_json, pkm)

        # Base statistics creation
        adding_base_stats(parsed_json, pkm)

        # Sprites
        adding_sprites(i, pkm)

        # Add pokemon into the dictionary
        pokes["pokemons"].insert(i, pkm)
        print("Donwloaded pokèmon " + str(i) + ": " + str(pkm))

    # Write dictionary to the disk
    with open("pokedb.json", "w") as file:
        val = json.dumps(pokes, indent=2, sort_keys=True)
        file.write(val + "\n")


def main():
    download_pokèmon(10)


if __name__ == '__main__':
    main()
