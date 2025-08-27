import "./CharacterItem.css"

export type CharacterProps = {
    name: string,
    level: number,
    currentHP: number,
    maxHP: number,
    currentMP: number,
    maxMP: number
}

export default function CharacterItem(characterProps: CharacterProps) {
    return (<div className="character-item">
        <div>
            <img />
        </div>
        <div>
            <span>{characterProps.name}</span>
            <CharacterStats {...characterProps} />
        </div>
        <div>
            <LevelStats />
        </div>
    </div>)
}



function CharacterStats(props: CharacterProps) {
    return (<div className="character-stats">
        <div><span className="menu-item-title">LV</span><span>{props.level}</span></div>
        <div><span className="menu-item-title">HP</span><span>{props.currentHP}/{props.maxHP}</span></div>
        <div><span className="menu-item-title">MP</span><span>{props.currentMP}/{props.maxMP}</span></div>
    </div>)
}


function LevelStats(props: any) {
    return (<div className="level-stats">
        <div>
            <span>next level</span>
            <svg className="">
                <LevelBar />
            </svg>
        </div>
        <div>
            <span>Limit level {props.limitLevelValue || 2}</span>
            <svg>
                <LevelBar />
            </svg>
        </div>
    </div>)
}

function LevelBar() {
    return <rect x="0" y="0" width="100%" 
        height="12" 
        rx="6" 
        fill="pink" 
        role="progressbar" 
        stroke="white"
        radius={0} 
        strokeWidth={2}></rect>
}