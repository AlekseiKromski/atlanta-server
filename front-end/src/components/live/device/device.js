import {
    Select, SelectItem,
} from "@nextui-org/react";
import DeviceStyle from "./device.module.css";
import Wrapper from "../../common/wrapper/wrapper";

export default function Device({device, setDevice, devices}) {

    return (
        <Wrapper
            title="Devices"
            modal={
                {
                    title: "Devices",
                    body: (
                        <p>
                            TODO: ...
                        </p>
                    )
                }
            }
        >
            <div className={DeviceStyle.Body}>
                {
                    devices.length === 0
                        ? <span>no devices</span>
                        : <Select
                            label="Select device"
                            className="max-w-full"
                            onChange={e => setDevice(e.target.value)}
                            value={device}
                        >
                            {
                                devices
                                    .map(device => (
                                        <SelectItem key={device} value={device}>
                                            {device}
                                        </SelectItem>
                                    ))
                            }
                        </Select>
                }
            </div>
        </Wrapper>
    )
}