import {useEffect, useState} from "react";
import LiveStyle from "./live.module.css";
import Device from "../../components/live/device/device";
import LiveView from "../../components/live/liveView/liveView";
import {useSelector} from "react-redux"
import LiveServer from "../../components/live/live/live"

export default function Live() {
    const application = useSelector((state) => state.application);

    const [device, setDevice] = useState("")
    const [devices, setDevices] = useState([])
    const [labels, setLabels] = useState([])
    const [datapoints, setDatapoints] = useState([])
    const [lastMessage, setLastMessage] = useState([])
    const [connected, setConnected] = useState(false)
    const [connection, setConnection] = useState(null)

    function reconnect() {
        if (connection != null) {
            connection.close()
        }

        connect()
    }

    function connect() {
        let connect = new WebSocket(
            process.env.REACT_APP_WEBSOCKET_URL + "/ws/connect?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MTEyMDIzMzQsImlkIjoidGVzdC11c2VyLWlkIn0.BRiX9d94lyt-AKgu4Oul-Oje44v98tVCIUJsmFRoDDQ",
        )
        connect.onopen = e => {
            setConnected(true)
        }
        connect.onmessage = (event) => {
            let received = JSON.parse(event.data)
            if (received.action !== "INFO") {
                return
            }
            let payload = JSON.parse(received.payload)
            setLastMessage([...payload])
        }

        connect.onerror = () => {
            setConnected(false)
        }

        connect.onclose = () => {
            setConnected(false)
        }

        setConnection(connect)
    }

    useEffect(() => {
        if (connection == null) {
            connect()
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
            <div className={LiveStyle.LiveStatus + " flex"}>
                <Device
                    device={devices}
                    setDevice={setDevice}
                    devices={
                        datapoints
                            .map(dp => dp.device_id)
                            .filter((value, index, array) => array.indexOf(value) === index)
                    }
                />
                <LiveServer reconnect={reconnect} connected={connected}/>
            </div>
            {
                device && <LiveView labels={labels} device={device} datapoints={datapoints}/>
            }
        </div>
    )
}