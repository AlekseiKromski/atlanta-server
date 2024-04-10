import SearchBoxStyle from "./searchBox.module.css"
import {Button, Checkbox, Chip, Input, Select, SelectItem} from "@nextui-org/react";
import {useEffect, useState} from "react";
import {useSelector} from "react-redux"
import Wrapper from "../../common/wrapper/wrapper";
import moment from "moment"

export default function SearchBox({history, setHistory, labels, devices, callback, searchOptions}) {

    const application = useSelector((state) => state.application);

    const [startDate, setStartDate] = useState("")
    const [endDate, setEndDate] = useState("")
    const [selectedDatapointTypes, setSelectedDatapointTypes] = useState([])
    const [selectedDevice, setSelectedDevice] = useState("")
    const [showType, setShowType] = useState("")
    const [ignored, setIgnored] = useState(false)
    const [loader, setLoader] = useState(false)
    const [loaderQuery, setLoaderQuery] = useState(false)

    const find = () => {
        setLoader(true)

        let start = startDate !== "" ? moment(startDate).format().split("+")[0] + "Z" : null
        let end = endDate !== "" ? moment(endDate).format().split("+")[0] + "Z" : null
        application.axios.post(`/api/datapoints/find`, {
            start: start,
            end: end,
            select: selectedDatapointTypes,
            device: selectedDevice,
            ignored: ignored
        })
            .then(res => {
                callback({
                    type: showType,
                    data: res.data
                })
                setTimeout(() => setLoader(false), 1000)
            })
            .catch(e => {
                setTimeout(() => setLoader(false), 1000)
            })
    }

    const save = () => {
        setLoaderQuery(true)

        let start = startDate !== "" ? moment(startDate).format().split("+")[0] + "Z" : null
        let end = endDate !== "" ? moment(endDate).format().split("+")[0] + "Z" : null
        let payload = {
            key: "history",
            value: JSON.stringify([
                {
                    name: moment.utc( Date.now() ).format(),
                    value: {
                        start: start,
                        end: end,
                        select: selectedDatapointTypes,
                        device: selectedDevice,
                        ignored: ignored,
                        types: showType,
                    }
                },
                ...history.value
            ])
        }

        setHistory(
            {
                key: "history",
                value: [
                    {
                        name: moment.utc(Date.now()).format(),
                        value: {
                            start: start,
                            end: end,
                            select: selectedDatapointTypes,
                            device: selectedDevice,
                            ignored: ignored,
                            types: showType,
                        }
                    },
                    ...history.value
                ]
            }
    )

        application.axios.post(`/api/store/upsert`, payload)
            .then(res => {
                setTimeout(() => setLoaderQuery(false), 1000)
            })
            .catch(e => {
                setTimeout(() => setLoaderQuery(false), 1000)
            })
    }

    useEffect(() => {
        if (searchOptions == null) {
            return
        }
        setStartDate(searchOptions.start.slice(0, -1))
        setEndDate(searchOptions.end.slice(0, -1))
        setIgnored(searchOptions.ignored)
        setIgnored(searchOptions.ignored)
        setSelectedDatapointTypes(searchOptions.select)
        setShowType(searchOptions.types)

        if (devices.find(d => d.id === searchOptions.device)) {
            setSelectedDevice(searchOptions.device)
        }

    }, [searchOptions]);

    return (
        <Wrapper width="70%" title="Search" modal={{
            title: "Search help",
            body: (
                <p>
                    TODO: ...
                </p>
            )
        }}>
            <div>
                <div>
                    <div className="flex w-full gap-1.5">
                        <Input
                            labelPlacement="inside"
                            type="datetime-local"
                            placeholder="none"
                            label="Start"
                            className="max-w-full"
                            onChange={e => setStartDate(e.target.value)}
                            value={startDate}
                        />
                        <Input
                            type="datetime-local"
                            label="End"
                            className="max-w-full"
                            placeholder="none"
                            onChange={e => setEndDate(e.target.value)}
                            value={endDate}
                        />
                    </div>
                    <div className={SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5 flex-col"}>
                        <Select
                            label="Select datapoint type"
                            selectionMode="multiple"
                            className="max-w-max"
                            selectedKeys={selectedDatapointTypes}
                            onChange={e => {
                                if (e.target.value.search("all") !== -1) {

                                    // If all labels already set, we should remove it
                                    if (selectedDatapointTypes.length == labels.length) {
                                        setSelectedDatapointTypes([])
                                        return
                                    }

                                    setSelectedDatapointTypes(labels)
                                    return
                                }

                                let values = e.target.value.split(",")
                                setSelectedDatapointTypes(values)
                            }}
                        >
                            <SelectItem selected key="all" value="all">
                                All
                            </SelectItem>
                            {
                                labels && labels.map(label => (
                                    <SelectItem selected key={label} value={label}>
                                        {label}
                                    </SelectItem>
                                ))
                            }
                        </Select>
                    </div>
                    <div className={SearchBoxStyle.SelectDatapointType + " flex flex-col w-full gap-1.5"}>
                        {
                            <Select
                                label="Select device"
                                className="max-w-max"
                                onChange={e => setSelectedDevice(e.target.value)}
                                selectedKeys={[selectedDevice]}
                            >
                                {
                                    devices.map(device => (
                                        <SelectItem key={device.id} value={device.id}>
                                            {device.id}
                                        </SelectItem>
                                    ))
                                }
                            </Select>
                        }

                        {
                            selectedDevice &&
                            <div>
                                <Chip color="warning" variant="bordered">
                                    Selected: {devices.find(d => d.id === selectedDevice).description}
                                </Chip>
                            </div>
                        }

                    </div>
                    <div className={SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5"}>
                        <Select
                            label="Select show type"
                            className="max-w-xs"
                            onChange={e => setShowType(e.target.value)}
                            selectedKeys={[showType]}
                        >
                            <SelectItem key="map" value="map">
                                Map
                            </SelectItem>
                            <SelectItem key="chart" value="chart">
                                Chart
                            </SelectItem>
                        </Select>
                    </div>
                    <div className={SearchBoxStyle.SelectDatapointType + " flex w-full gap-1.5"}>
                        <Checkbox isSelected={ignored} onChange={e => setIgnored(e.target.checked)} color="warning">With
                            ignored</Checkbox>
                    </div>
                </div>
                <div className="flex justify-end gap-1">

                    <Button color="default" variant="flat" onClick={e => {
                        setStartDate("")
                        setEndDate("")
                        setSelectedDatapointTypes([])
                        setShowType("")
                        setIgnored(false)
                        callback({
                            type: "",
                            data: []
                        })
                        setSelectedDevice("")
                    }}>
                        Clear
                    </Button>

                    <Button isDisabled={
                        startDate === "" || endDate == "" || selectedDatapointTypes.length === 0 || showType === "" || selectedDevice == ""
                    } isLoading={loaderQuery} color="success" variant="flat" onClick={save}>
                        Save query
                    </Button>

                    <Button isDisabled={
                        startDate === "" || endDate == "" || selectedDatapointTypes.length === 0 || showType === "" || selectedDevice == ""
                    } isLoading={loader} color="secondary" onClick={find}>
                        Find
                    </Button>
                </div>
            </div>
        </Wrapper>
    )
}