import {useEffect, useState} from "react";
import LiveStyle from "./live.module.css";
import {Chip} from "@nextui-org/react";
import Device from "../../components/live/device/device";
import LiveView from "../../components/live/liveView/liveView";
import {useSelector} from "react-redux"
import Wrapper from "../../components/common/wrapper/wrapper";

export default function Live() {
    const application = useSelector((state) => state.application);

    const [device, setDevice] = useState("")
    const [devices, setDevices] = useState([])
    const [labels, setLabels] = useState([])
    const [datapoints, setDatapoints] = useState([])
    const [lastMessage, setLastMessage] = useState([])
    const [connected, setConnected] = useState(false)
    const [connection, setConnection] = useState(null)

    useEffect(() => {
        if (connection == null) {
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

            <Device
                device={devices}
                setDevice={setDevice}
                devices={
                    datapoints
                        .map(dp => dp.device_id)
                        .filter((value, index, array) => array.indexOf(value) === index)
                }
            />
            {
                device && <LiveView labels={labels} device={device} datapoints={datapoints}/>
            }
        </div>
    )
}