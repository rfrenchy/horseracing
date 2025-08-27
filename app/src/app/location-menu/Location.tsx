import "./location.css";


export type LocationProps = {
    name: string
}

export default function Location(locationProps: LocationProps) {
    return (<div className="location-item"><span>{locationProps.name}</span></div>)
}