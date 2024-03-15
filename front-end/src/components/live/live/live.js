import Wrapper from "../../common/wrapper/wrapper";
import {Button, Chip} from "@nextui-org/react";
import LiveStyle from "./live.module.css";
import {useState} from "react";

export default function Live({reconnect, connected}) {

    const [loader, setLoader] = useState(false)

    return (
        <Wrapper
            title="Server status"
            width="30%"
            modal={
                {
                    title: "Server status help",
                    body: (
                        <p>
                            TODO: ...
                        </p>
                    )
                }
            }
        >
            <div className={LiveStyle.LiveBody + " flex flex-col"}>
                {connected ? <Chip color="success">Connected</Chip> : <Chip color="danger">Disconnected</Chip>}

                {!connected && <Button onClick={() => {
                    setLoader(true)
                    reconnect()
                    setTimeout(() => {
                        setLoader(false)
                    }, 1000)
                }} color="primary" isLoading={loader} variant="bordered">
                    Reconnect
                </Button>}
            </div>
        </Wrapper>
    )
}