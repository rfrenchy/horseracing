import "./layout.css"

import CharacterItem from "./components/CharacterItem";
import { CharacterProps } from "../types";


export default function CharacterMenu() {
    return <div className="character-menu">
        <CharacterItem {...Cloud}/>
        <CharacterItem {...Tifa}/>
        <CharacterItem {...Barret}/>
    </div>
}

const Cloud: CharacterProps = {
    name: "Cloud",
    level: 48,
    currentHP: 3000,
    maxHP: 3000,
    currentMP: 264,
    maxMP: 473
}

const Tifa: CharacterProps = {
    name: "Tifa",
    level: 44,
    currentHP: 3200,
    maxHP: 3200,
    currentMP: 118,
    maxMP: 393
}

const Barret: CharacterProps = {
    name: "Barret",
    level: 45,
    currentHP: 3334,
    maxHP: 3334,
    currentMP: 212,
    maxMP: 361
};
