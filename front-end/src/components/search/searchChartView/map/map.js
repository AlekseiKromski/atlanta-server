import 'leaflet/dist/leaflet.css';
import L from 'leaflet';
import {MapContainer, Marker, Popup, TileLayer} from "react-leaflet";
import MapStyle from "./map.module.css"
import icon from 'leaflet/dist/images/marker-icon.png';
import iconShadow from 'leaflet/dist/images/marker-shadow.png';
import {
    Button,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Popover, PopoverContent, PopoverTrigger,
    useDisclosure
} from "@nextui-org/react";
import {useState} from "react";
let DefaultIcon = L.icon({
    iconUrl: icon,
    shadowUrl: iconShadow
});
L.Marker.prototype.options.icon = DefaultIcon;

export default function Map({wrapper, labels, datapoints}) {
    datapoints = datapoints.data.datapoints

    const position = [59.3573116, 27.4136646]
    const {isOpen, onOpen, onOpenChange} = useDisclosure();

    let lastLength = null
    let result = true
    labels.forEach(label => {
        if (!datapoints[label]) {
            return;
        }
        let currentLength = datapoints[label].length
        if (lastLength == null) {
            lastLength = currentLength
            return
        }

        if (currentLength !== lastLength ){
            result = false
        }
    })

    return (
        <div className={wrapper ? MapStyle.Wrapper : ""}>
            <div className={MapStyle.Header + " flex justify-between items-center rounded-md"}>
                <h1 className="font-bold">Map view</h1>
                <Button size="sm" color="primary" variant="light" onPress={onOpen}>Help</Button>
                <Modal backdrop="blur" isOpen={isOpen} onOpenChange={onOpenChange}>
                    <ModalContent>
                        {(onClose) => (
                            <>
                                <ModalHeader className="flex flex-col gap-1">Search instruction</ModalHeader>
                                <ModalBody>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                                        Nullam pulvinar risus non risus hendrerit venenatis.
                                        Pellentesque sit amet hendrerit risus, sed porttitor quam.
                                    </p>
                                    <p>
                                        Lorem ipsum dolor sit amet, consectetur adipiscing elit.
                                        Nullam pulvinar risus non risus hendrerit venenatis.
                                        Pellentesque sit amet hendrerit risus, sed porttitor quam.
                                    </p>
                                    <p>
                                        Magna exercitation reprehenderit magna aute tempor cupidatat consequat elit
                                        dolor adipisicing. Mollit dolor eiusmod sunt ex incididunt cillum quis.
                                        Velit duis sit officia eiusmod Lorem aliqua enim laboris do dolor eiusmod.
                                        Et mollit incididunt nisi consectetur esse laborum eiusmod pariatur
                                        proident Lorem eiusmod et. Culpa deserunt nostrud ad veniam.
                                    </p>
                                </ModalBody>
                                <ModalFooter>
                                    <Button color="danger" variant="light" onPress={onClose}>
                                        Close
                                    </Button>
                                </ModalFooter>
                            </>
                        )}
                    </ModalContent>
                </Modal>
            </div>
            <div className={MapStyle.Body + " flex flex-col"}>
                <div>
                    <Popover placement="right">
                        <PopoverTrigger>
                            {
                                <Button color={result ? "success" : "danger"}>Show state {result ? "✅" : "⛔"}</Button>
                            }
                        </PopoverTrigger>
                        <PopoverContent>
                            <b>Count of records</b>
                            <ul>
                                {
                                    labels && labels.length !== 0 &&
                                    labels.map(label => {
                                        if (datapoints[label]) {
                                            return <li>Count {label} records: <b>{datapoints[label].length}</b></li>
                                        }
                                    })
                                }
                            </ul>
                        </PopoverContent>
                    </Popover>
                </div>
                <MapContainer className={MapStyle.Map} center={position} zoom={10} scrollWheelZoom={true}>
                    <TileLayer
                        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>'
                        url="https://{s}.basemaps.cartocdn.com/light_all/{z}/{x}/{y}{r}.png"
                    />
                    {
                        datapoints["Geo-position"] && datapoints["Geo-position"].map((dp, index) => {
                            let coords = dp.value.split(",")
                            return (
                                <Marker position={[coords[0], coords[1]]}>
                                    <Popup>
                                        <ul>
                                            <li>Value: <b>{dp.value}</b></li>
                                            <li>Unit: <b>{dp.unit}</b></li>
                                            <li>Measurement time: <b>{dp.measurement_time}</b></li>
                                            <li>Created: <b>{dp.created_at}</b></li>
                                            <li>Updated: <b>{dp.updated_at}</b></li>

                                            {
                                                labels && labels.length !== 0 &&
                                                labels.map(label => {
                                                    if (datapoints[label] && datapoints[label][index]) {
                                                        return <li>{label}: <b>{datapoints[label][index].value}</b></li>
                                                    }
                                                })
                                            }

                                        </ul>
                                    </Popup>
                                </Marker>
                            )
                        })
                    }
                </MapContainer>
            </div>
        </div>
    );
}