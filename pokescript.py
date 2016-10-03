#!/usr/bin/env python
# ------------------------------------------------------------------
# Script per fregare i dati da PokèAPI e renderli compatibili a PBG.
# Si ringrazia il grande capo Gianni Modica per lo script!
# ------------------------------------------------------------------
from urllib.request import Request, urlopen
import json

categoryTable = {
    "physical": 0, "special": 1, "status": 2
}

typesTable = {
    "normal": 0, "fire": 1, "fighting": 2, "water": 3, "flying": 4, "grass": 5,
    "poison": 6, "electric": 7, "ground": 8, "psychic": 9, "rock": 10, "ice": 11,
    "bug": 12, "dragon": 13, "ghost": 14, "dark": 15, "steel": 16, "fairy": 17, "???": 18
}

baseStatTable = {
    "hp": 0, "attack": 1, "defense": 2, "special-attack": 3, "special-defense": 4, "speed": 5
}


def type_to_gotype(typ: str):
    typ1 = typesTable[typ]
    return typ1 if typ1 is not None else -1


def category_to_gocategory(cat: str):
    cat1 = categoryTable[cat]
    return cat1 if cat1 is not None else 0


def adding_type(pkorigin: dict, pkdest: dict):
    pkdest["type"][0] = type_to_gotype(((pkorigin["types"])[0])["type"]["name"])
    if len(pkorigin["types"]) > 1:
        pkdest["type"][1] = type_to_gotype(((pkorigin["types"])[1])["type"]["name"])


def adding_base_stats(pkorigin: dict, pkdest: dict):
    for st in pkorigin["stats"]:
        pkdest["baseStats"][baseStatTable[st["stat"]["name"]]] = st["base_stat"]


def adding_sprites(i: int, pkdest: dict):
    pkdest["sprites"] = {
        # Works with hotlinking
        "front": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/" + str(i) + ".png",
        "back": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/back/" + str(i) + ".png"
    }


def download_pokèmon(num: int):
    req = Request('http://pokeapi.co/api/v2/pokemon/'  + str(num), headers={'User-Agent': 'Mozilla/5.0'})
    r = urlopen(req)
    parsed_json = json.loads(r.read().decode('utf-8'))
    pkm = {
        "name": parsed_json["name"].title(),
        "pokedex": num,
        "baseStats": [ -1, -1, -1, -1, -1, -1 ],
        "type": [ -1, -1 ]
    }

    # Type creation
    adding_type(parsed_json, pkm)
    # Base statistics creation
    adding_base_stats(parsed_json, pkm)
    # Sprites
    adding_sprites(num, pkm)

    # Done
    print("Donwloaded pokèmon " + str(num) + ": " + str(pkm))
    return pkm


def download_move(num: int):
    req = Request('http://pokeapi.co/api/v2/move/' + str(num), headers={'User-Agent': 'Mozilla/5.0'})
    r = urlopen(req)
    parsed_json = json.loads(r.read().decode('utf-8'))
    move = {
        "name": parsed_json["names"][0]["name"],
        "accuracy": parsed_json["accuracy"],
        "pps": parsed_json["pp"],
        "priority": parsed_json["priority"],
        "power": parsed_json["power"],
        "type": type_to_gotype(parsed_json["type"]["name"]),
        "category": category_to_gocategory(parsed_json["damage_class"]["name"])
    }

    print("Donwloaded move " + str(num) + ": " + str(move))
    return move


def main():
    pokemon_num = 5  # 151
    move_num = 5     # 165
    d = {
        "generation": 1,
        "pokemon_count": pokemon_num,
        "move_count": move_num,
        "pokemons": [download_pokèmon(i) for i in range(1, pokemon_num + 1)],
        "moves": [download_move(i) for i in range(1, move_num + 1)]
    }

    # Write dictionary to the disk
    with open("pokedb.json", "w") as file:
       val = json.dumps(d, indent=4, sort_keys=True)
       file.write(val + "\n")


if __name__ == '__main__':
    main()
