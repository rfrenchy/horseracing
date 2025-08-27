'use client'

import "./layout.css"

import { Dispatch, SetStateAction, useState } from 'react';

type MainViewPages = "use" | "key-items";

export default function Items() {

    const [selectedItem, selectItemDescription] =
        useState(items[0].description)

    const [arrangeMenuHidden, arrangeMenuSetHidden] = useState(true)

    const [activePage, changeActivePage] = useState("use" as MainViewPages);

    const CurrentActivePage = getCurrentActivePage(activePage);


    return (
        <div className="item-page">
            <div className="top-bar">
                <MenuNav 
                    arrangeHidden={arrangeMenuHidden} 
                    activePage={activePage}
                    clickArrange={arrangeMenuSetHidden}
                    changeActivePage={changeActivePage}
                />
                <PageName />
            </div>
            <div className="description-bar">
                <span>{selectedItem}</span>
            </div>
            {CurrentActivePage}
            <MenuNavArrange hidden={arrangeMenuHidden}/>
        </div>)
}

type MenuNavProps = {
    arrangeHidden: boolean
    activePage: MainViewPages

    clickArrange: MenuHiddenSelect
    changeActivePage: ChangeActivePage
}

function getCurrentActivePage(activePage: MainViewPages) {
    switch (activePage) {
        case "use": return <UseView />;
        case "key-items": return <KeyItems keyItems={keyItemData}/>;
    }
}

function MenuNav(props: MenuNavProps) {
    return (<div className="menu-nav">
        <span onClick={() => props.changeActivePage("use")}>Use</span>
        <span onClick={() => props.clickArrange(!props.arrangeHidden)}>Arrange</span>
        <span onClick={() => props.changeActivePage("key-items")}>Key Items</span>
    </div>)
}

type MenuNavArrangeProps = {
    hidden: boolean
}

function MenuNavArrange(props: MenuNavArrangeProps) {
    return (<div className="arrange-menu" hidden={props.hidden}>
        <ul>
            <li>Customise</li>
            <li>Field</li>
            <li>Battle</li>
            <li>Throw</li>
            <li>Type</li>
            <li>Name</li>
            <li>Most</li>
            <li>Least</li>
        </ul>
    </div>)
}

function PageName() {
    return (
        <div className="page-name">
            <span>Item</span>
        </div>
    )
}

type ItemSelect = Dispatch<SetStateAction<string | undefined>>
type MenuHiddenSelect = Dispatch<SetStateAction<boolean>>
type ChangeActivePage = Dispatch<SetStateAction<MainViewPages>> 

type ItemListProps = {
    items: Item[];
    selectItemDescription: ItemSelect
}

type Item = {
    name: string,
    amount: number,
    selectable: boolean,
    description?: string
}

function ItemList(props: ItemListProps) {
    return (<div className="item-list">
        <ul>
            {props.items.map(x =>
                <ItemListItem
                    key={x.name}
                    selectItem={props.selectItemDescription}
                    item={x}
                />)}
        </ul>
    </div>);
}

type ItemListItemProps = {
    item: Item,
    selectItem: ItemSelect
}

function ItemListItem(props: ItemListItemProps) {
    const [amount, setAmount] = useState(props.item.amount);

    const selectable = props.item.selectable ? "selectable" : "unselectable";

    return (<li
        className={`item-list-item ${selectable}`}
        onClick={() => {
            if (props.item.selectable) {
                setAmount(amount - 1)
                props.selectItem(props.item.description)
            };
        }}
    >
        <span>{props.item.name}</span><span>:</span><span>{amount}</span>
    </li>)

}

const items: Item[] = [
    { name: "Soft", amount: 6, selectable: false, description: "Cures [Petrify]" },
    { name: "Turbo Ether", amount: 18, selectable: true, description: "Restores MP" },
    { name: "Ether", amount: 10, selectable: true },
    { name: "Phoenix Down", amount: 29, selectable: true },
    { name: "X-Potion", amount: 9, selectable: true },
    { name: "Remedy", amount: 1, selectable: false },
    { name: "Reagan Greens", amount: 3, selectable: false },
    { name: "M-Tentacles", amount: 5, selectable: false },
    { name: "Megalixir", amount: 6, selectable: true },
    { name: "Elixir", amount: 12, selectable: true },
]

type KeyItem = {
    name: string;
    description: string;
}

const keyItemData: KeyItem[] = [
    { name: "Silk Dress", description: "Dress made of silk"},
    { name: "Blonde Wig", description: ""},
    { name: "Glass Tiara", description: ""},
    { name: "Cologne", description: ""},
    { name: "Key to Ancients", description: ""},
    { name: "Lunar Harp", description: ""},
    { name: "Basement Key", description: ""},
    { name: "PHS", description: ""},
    { name: "Gold Ticket", description: ""},
    { name: "Keystone", description: ""},
    { name: "Leviathan Scales", description: ""},
    { name: "Glacier Map", description: ""},
    { name: "Snowboard", description: ""},
]

type KeyItemProps = {
    keyItems: KeyItem[]
}

function KeyItems(props: KeyItemProps) {
    return (<div className="key-item-list">
        {props.keyItems.map((k) => 
            <div key={k.name} className="key-item-data"><span>{k.name}</span></div>
        )}
    </div>)

}

function UseView(props :any) {
    return <div className="main-view">
        <div className="character-screen"></div>
         <ItemList
            items={items}
            selectItemDescription={props.selectItemDescription}
         />
    </div>
}