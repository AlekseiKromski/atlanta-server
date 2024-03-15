import {
    Select, SelectItem,
} from "@nextui-org/react";
import Wrapper from "../../common/wrapper/wrapper";

export default function SearchHistory() {

    return (
        <Wrapper width="30%" title="History" modal={{
            title: "History help",
            body: (
                <p>
                    TODO: ...
                </p>
            )
        }}>
            <div className="flex flex-col justify-between h-full">
                <div>
                    <Select
                        label="Select record"
                        className="max-w-xs"
                    >
                        <SelectItem key="map" value="map">
                            Map
                        </SelectItem>
                    </Select>
                </div>
            </div>
        </Wrapper>
    )
}