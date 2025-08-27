import "./miscmenu.css";

export type MiscMenuData = {
    CurrentTime: number;
    Gil: number;
}

export default function MiscMenu(miscMenuProps: MiscMenuData) {
    return (<div className="misc-menu">
        <div className="misc-menu-item"><span>Time</span><span>{miscMenuProps.CurrentTime}</span></div>
        <div className="misc-menu-item"><span>Gil</span><span>{miscMenuProps.Gil}</span></div>
    </div>)
}