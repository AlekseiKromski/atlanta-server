import {useEffect, useState} from "react";
import LiveStyle from "./live.module.css";
import {Chip} from "@nextui-org/react";
import Device from "../../components/live/device/device";
import LiveView from "../../components/live/liveView/liveView";
import {useSelector} from "react-redux"

export default function Live() {
    const application = useSelector((state) => state.application);

    const [device, setDevice] = useState("")
    const [devices, setDevices] = useState([])
    const [labels, setLabels] = useState([])
    const [datapoints, setDatapoints] = useState([
        {
            "id": "0bc3731e-526e-460b-9b32-ab63e0e3ac16",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "101738.000000",
            "type": "float",
            "unit": "Pa",
            "label": "Pressure",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.235184Z",
            "updated_at": "2024-03-13T23:29:19.235184Z"
        },
        {
            "id": "c188273b-6e18-42fa-bcce-1cf051dd064b",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "34.910000",
            "type": "float",
            "unit": "M",
            "label": "BMP180 Altitude",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.279575Z",
            "updated_at": "2024-03-13T23:29:19.279575Z"
        },
        {
            "id": "711e49ca-a3b7-4f1b-98c0-fd5bf23ce7fc",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "16.500000",
            "type": "float",
            "unit": "C",
            "label": "BMP180 TEMP",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.307835Z",
            "updated_at": "2024-03-13T23:29:19.307835Z"
        },
        {
            "id": "ec0f6b5c-8044-4e36-b42e-cd5ea5560cd9",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "43.000000",
            "type": "float",
            "unit": "percentage",
            "label": "Humidity",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.336873Z",
            "updated_at": "2024-03-13T23:29:19.336873Z"
        },
        {
            "id": "02c42d19-0732-4281-ad14-9edd45c6c7f1",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "18.000000",
            "type": "float",
            "unit": "C",
            "label": "Temperature",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.445471Z",
            "updated_at": "2024-03-13T23:29:19.445471Z"
        },
        {
            "id": "a8430398-1a35-4989-a95b-e6c300ddf048",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "59.337040,27.420391",
            "type": "float",
            "unit": "coordinate",
            "label": "Geo-position",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.473826Z",
            "updated_at": "2024-03-13T23:29:19.473826Z"
        },
        {
            "id": "3713b91c-3423-43d7-a7f9-b21c410d8f39",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "5.000000",
            "type": "float",
            "unit": "M",
            "label": "Altitude",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:19Z",
            "created_at": "2024-03-13T23:29:19.502447Z",
            "updated_at": "2024-03-13T23:29:19.502447Z"
        },
        {
            "id": "d3d6aefe-bd43-41c0-bec5-0f2d6ddf9000",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "34.910000",
            "type": "float",
            "unit": "M",
            "label": "BMP180 Altitude",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.292125Z",
            "updated_at": "2024-03-13T23:29:14.292125Z"
        },
        {
            "id": "a38a7cb0-385b-4a34-b133-63b63aa97648",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "59.337040,27.420391",
            "type": "float",
            "unit": "coordinate",
            "label": "Geo-position",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.338811Z",
            "updated_at": "2024-03-13T23:29:14.338811Z"
        },
        {
            "id": "914643d9-e04d-4f80-95c8-d01d43aa0bc2",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "5.000000",
            "type": "float",
            "unit": "M",
            "label": "Altitude",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.366273Z",
            "updated_at": "2024-03-13T23:29:14.366273Z"
        },
        {
            "id": "d80b31c6-40f7-4c04-a6c4-903e0b1fe062",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "18.000000",
            "type": "float",
            "unit": "C",
            "label": "Temperature",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.4224Z",
            "updated_at": "2024-03-13T23:29:14.4224Z"
        },
        {
            "id": "0ccadf41-8116-469e-a608-0ea666ff8271",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "16.500000",
            "type": "float",
            "unit": "C",
            "label": "BMP180 TEMP",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.447349Z",
            "updated_at": "2024-03-13T23:29:14.447349Z"
        },
        {
            "id": "c8c497d6-de40-466b-88a1-2e0261fa3d11",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "43.000000",
            "type": "float",
            "unit": "percentage",
            "label": "Humidity",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.474702Z",
            "updated_at": "2024-03-13T23:29:14.474702Z"
        },
        {
            "id": "3f5fd9ba-6568-4875-96ec-52b750759757",
            "device_id": "3cc76ff4-cbaa-436c-b727-45d526facfc7",
            "value": "101738.000000",
            "type": "float",
            "unit": "Pa",
            "label": "Pressure",
            "flags": "",
            "measurement_time": "2024-03-13T23:29:14Z",
            "created_at": "2024-03-13T23:29:14.729369Z",
            "updated_at": "2024-03-13T23:29:14.729369Z"
        }
    ])
    const [lastMessage, setLastMessage] = useState([])
    const [connected, setConnected] = useState(false)
    const [connection, setConnection] = useState(null)

    useEffect(() => {
        if (connection == null){
            let connection = new WebSocket(
                 process.env.REACT_APP_WEBSOCKET_URL + "/ws/connect?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTEyMDIzMzQsImlkIjoidGVzdC11c2VyLWlkIn0.BRiX9d94lyt-AKgu4Oul-Oje44v98tVCIUJsmFRoDDQ",
            )
            connection.onopen = e => {
                setConnected(true)
            }
            connection.onmessage = (event) => {
                let received = JSON.parse(event.data)
                if (received.action !== "INFO") {
                    return
                }
                let payload = JSON.parse(received.payload)
                setLastMessage([...payload])
            }

            connection.onerror = () => {
                setConnected(false)
            }

            connection.onclose = () => {
                setConnected(false)
            }

            setConnection(connection)
        }
        if (labels.length === 0) {
            application.axios.get("/api/datapoints/info/labels")
                .then(res => setLabels(res.data))
                .catch(e => console.log(e))
        }

        setDatapoints([...lastMessage, ...datapoints])
    }, [lastMessage])

    useEffect(() => {
        return () => {
            if (connection != null) {
                connection.close()
            }
        }
    }, [connection]);

    return (
        <div className={LiveStyle.LiveBody + " w-full flex flex-col"}>
            {connected ? <Chip color="success">Connected</Chip> : <Chip color="danger">Disconnected</Chip>}

            <Device device={devices} setDevice={setDevice} devices={
                datapoints
                    .map(dp => dp.device_id)
                    .filter((value, index, array) => array.indexOf(value) === index)
            } />
            {
                device && <LiveView labels={labels} device={device} datapoints={datapoints}/>


            }
        </div>
    )
}