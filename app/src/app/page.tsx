import "./layout.css"

import CharacterMenu from "./character-menu/layout";
import SelectMenu from "./select-menu/SelectMenu";
import Location from "./location-menu/Location";
import MiscMenu from "./misc-menu/MiscMenu";


export default function Home() {
  return (
    <div 
      className="font-sans character-page"
    >
      <CharacterMenu/>
      <SelectMenu />
      <Location name={currentLocation} />
      <MiscMenu {...miscMenuData}/>
    </div>
  );
}

const currentLocation = "Underwater Reactor";

const miscMenuData = {
  CurrentTime: 123,
  Gil: 196893
}
