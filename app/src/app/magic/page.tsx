'use client';

import "./magic.css"

import { Dispatch, SetStateAction, useState } from 'react'

type ActivePage = "" | "magic" | "summon" | "enemy-skill"


export default function Magic() {

    const [activePage, setActivePage] = useState("" as ActivePage)
    const [description, descriptionChange] = useState()

    return <div className="magic-page">
        <div className="top-view">
            <div className="character-screen" ></div>
            <div className="magic-info">
                <div className="screen-name"><span>Magic</span></div>
                <MagicTopSelect setActivePage={setActivePage} />
            </div>
        </div>
        <div className="bottom-view">
            <div className="description-bar"></div>
            <BottomViewSwitch activePage={activePage} />
        </div>
    </div>
}

type MagicTopSelectProps = {
    setActivePage: Dispatch<SetStateAction<ActivePage>>
}

function MagicTopSelect(props: MagicTopSelectProps) {
    return (
        <div className="magic-top-select">
            <ul>
                <li onClick={() => props.setActivePage("magic")}><span>Magic</span></li>
                <li onClick={() => props.setActivePage("summon")}><span>Summon</span></li>
                <li onClick={() => props.setActivePage("enemy-skill")}><span>Enemy-Skill</span></li>
            </ul>
        </div>
    )
}

type BottomViewProps = {
    activePage: ActivePage
}

function BottomViewSwitch(props: BottomViewProps) {
    switch (props.activePage) {
        case "magic": return <MagicView magic={magicData} />
        case "summon": return <SummonView />
        case "enemy-skill": return <EnemySkillView magic={enemeySkillData} />
        default: <></>
    }
}

type MagicData = {
    name: string;
    description: string;
    mp: number;
    selectable: boolean;
}

const magicData: MagicData[] = [
    { name: "Cure", description: "Restores HP", mp: 5, selectable: true },
    { name: "Cure2", description: "Restores HP", mp: 15, selectable: true },
    { name: "Cure3", description: "Restores HP", mp: 32, selectable: true },
    { name: "Regen", description: "Restores HP", mp: 32, selectable: false },
    { name: "Ultima", description: "Extreme Magic Attack", mp: 87, selectable: false },
]

const enemeySkillData: MagicData[] = [
    { name: "Frog Song", description: "Causes [Sleepel/Frog] on one opponent", mp: 5, selectable: true },
    { name: "L4", description: "", mp: 0, selectable: true },
    { name: "Magic Hammer", description: "", mp: 0, selectable: true },
    { name: "White Wind", description: "", mp: 0, selectable: true },
    { name: "Big Guard", description: "", mp: 0, selectable: true },
    { name: "Death Force", description: "", mp: 0, selectable: true },
    { name: "Flame Thrower", description: "", mp: 0, selectable: true },
    { name: "Laser", description: "", mp: 0, selectable: true },
    { name: "Matra Magic", description: "", mp: 0, selectable: true },
    { name: "Bad Breath", description: "", mp: 0, selectable: true },
    { name: "Beta", description: "", mp: 0, selectable: true },
    { name: "Aqualung", description: "", mp: 0, selectable: true },
]

type MagicViewProps = {
    magic: MagicData[]
}

function MagicView(props: MagicViewProps) {
    return (<ul className="magic-select-list">
        {props.magic.map(m =>
            <li
                className={`${m.selectable ? "" : "unselectable"}`}
                key={m.name}
            >{m.name}
            </li>)}
    </ul>)
}

function SummonView() {
    return <div>Enemy-Skill View</div>
}

function EnemySkillView(props: MagicViewProps) {
    return (<ul className="summon-select-list">
        {props.magic.map(m =>
            <li
                className={`${m.selectable ? "" : "unselectable"}`}
                key={m.name}
            >{m.name}
            </li>)}
    </ul>)
}


